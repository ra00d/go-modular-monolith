package auth

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"slices"
	"time"

	"modular-monolith/lib/config"
	"modular-monolith/lib/logger"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

// JWTManager handles all JWT operations
type JWTManager struct {
	redisClient *redis.Client
	secretKey   []byte
	issuer      string
	tokenTTL    time.Duration
}

// CustomClaims defines your JWT claims structure
type CustomClaims struct {
	UserID      string   `json:"user_id"`
	Status      string   `json:"status"`
	Identifier  string   `json:"identifier"`
	Role        string   `json:"role"`
	Roles       []string `json:"roles"`
	Permissions []string `json:"permissions"`
	TokenID     string   `json:"token_id"` // Unique token identifier
	jwt.RegisteredClaims
}

// CreatTokenParams CustomClaims defines your JWT claims structure
type CreatTokenParams struct {
	UserID      string   `json:"user_id"`
	Identifier  string   `json:"identifier"`
	Status      string   `json:"status"`
	Role        string   `json:"role"`
	Roles       []string `json:"roles"`
	Permissions []string `json:"permissions"`
}

// TokenResponse contains the generated token and metadata
type TokenResponse struct {
	AccessToken string    `json:"access_token"`
	TokenType   string    `json:"token_type"`
	ExpiresAt   time.Time `json:"expires_at"`
	TokenID     string    `json:"token_id"`
}

// NewJWTManager creates a new JWT manager
func NewJWTManager(redisClient *redis.Client, cfg config.JwtConfig) *JWTManager {
	return &JWTManager{
		redisClient: redisClient,
		secretKey:   []byte(cfg.SecretKey),
		issuer:      cfg.Issure,
		tokenTTL:    time.Duration(cfg.Ttl) * time.Minute,
	}
}

// CreateAndStoreToken Create and store JWT token
func (m *JWTManager) CreateAndStoreToken(
	ctx context.Context,
	args CreatTokenParams,
) (*TokenResponse, error) {
	tokenID := uuid.New().String()
	now := time.Now()
	expiresAt := now.Add(m.tokenTTL)
	userID, identifier, role, roles, permissions, status := args.UserID, args.Identifier, args.Role, args.Roles, args.Permissions, args.Status
	// Create claims
	claims := &CustomClaims{
		UserID:      userID,
		Identifier:  identifier,
		Status:      status,
		Role:        role,
		Roles:       roles,
		Permissions: permissions,
		TokenID:     tokenID,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    m.issuer,
			Subject:   userID,
			Audience:  []string{"api"},
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			ID:        tokenID,
		},
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(m.secretKey)
	if err != nil {
		return nil, fmt.Errorf("failed to sign token: %w", err)
	}

	// Store in Redis with user session data
	sessionKey := fmt.Sprintf("session:%s", tokenID)
	sessionData := SessionData{
		ID:          tokenID,
		UserID:      userID,
		Identifier:  identifier,
		Status:      status,
		Role:        role,
		Roles:       roles,
		Permissions: permissions,
		Token:       tokenString,
		CreatedAt:   now,
		ExpiresAt:   expiresAt,
	}

	sessionBytes, err := json.Marshal(sessionData)
	if err != nil {
		return nil, fmt.Errorf("failed to store session: %w", err)
	}
	// Store hash with TTL
	err = m.redisClient.HSet(ctx, sessionKey, "sessionData", string(sessionBytes)).Err()
	if err != nil {
		return nil, fmt.Errorf("failed to store session: %w", err)
	}
	// Set expiration on the Redis key
	err = m.redisClient.Expire(ctx, sessionKey, m.tokenTTL).Err()
	if err != nil {
		return nil, fmt.Errorf("failed to set session expiration: %w", err)
	}

	// Also store token to user mapping for easy revocation
	userTokenKey := fmt.Sprintf("user_tokens:%s", userID)
	err = m.redisClient.SAdd(ctx, userTokenKey, tokenID).Err()
	if err != nil {
		return nil, fmt.Errorf("failed to map user to token: %w", err)
	}
	err = m.redisClient.Expire(ctx, userTokenKey, m.tokenTTL).Err()
	if err != nil {
		return nil, fmt.Errorf("failed to set user token mapping expiration: %w", err)
	}

	return &TokenResponse{
		AccessToken: tokenString,
		TokenType:   "Bearer",
		ExpiresAt:   expiresAt,
		TokenID:     tokenID,
	}, nil
}

// ValidateToken Validate token and check if it exists in Redis
func (m *JWTManager) ValidateToken(ctx context.Context, tokenString string) (*CustomClaims, error) {
	// Parse and validate token
	token, err := jwt.ParseWithClaims(
		tokenString,
		&CustomClaims{},
		func(token *jwt.Token) (any, error) {
			// Validate signing method
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return m.secretKey, nil
		},
	)
	if err != nil {
		logger.Error(err)
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token claims")
	}

	// Check if token exists and is valid in Redis
	sessionKey := fmt.Sprintf("session:%s", claims.TokenID)
	exists, err := m.redisClient.Exists(ctx, sessionKey).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to check session in Redis: %w", err)
	}

	if exists == 0 {
		return nil, errors.New("token has been revoked or expired")
	}

	// Optional: Verify token hasn't been tampered with by comparing stored data
	storedUserData, err := m.GetSessionData(ctx, claims.TokenID)
	if err != nil {
		return nil, fmt.Errorf("failed to get stored session data: %w", err)
	}
	storedUserID := storedUserData.UserID
	if storedUserID != claims.UserID {
		return nil, errors.New("token data mismatch")
	}

	return claims, nil
}

// GetClaimsFromToken Get user claims from token (without validation)
func (m *JWTManager) GetClaimsFromToken(tokenString string) (*CustomClaims, error) {
	parser := jwt.Parser{}
	claims := &CustomClaims{}

	token, _, err := parser.ParseUnverified(tokenString, claims)
	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	if claims, ok := token.Claims.(*CustomClaims); ok {
		return claims, nil
	}

	return nil, errors.New("failed to extract claims")
}

// Additional helper methods

// GetToken get the bearerToken from the request
func (m JWTManager) GetToken(ctx context.Context) string {
	return bearerToken(g.RequestFromCtx(ctx))
}

func (m JWTManager) RevokeCurrentToken(ctx context.Context) error {
	token := m.GetToken(ctx)
	err := m.RevokeToken(ctx, token)
	if err != nil {
		return err
	}
	return nil
}

// RevokeToken  Revoke token (logout)
func (m *JWTManager) RevokeToken(ctx context.Context, tokenString string) error {
	claims, err := m.ValidateToken(ctx, tokenString)
	if err != nil {
		return fmt.Errorf("token validation failed: %w", err)
	}

	// Delete session from Redis
	sessionKey := fmt.Sprintf("session:%s", claims.TokenID)
	err = m.redisClient.Del(ctx, sessionKey).Err()
	if err != nil {
		return fmt.Errorf("failed to revoke session: %w", err)
	}

	// Remove from user's token mapping
	userTokenKey := fmt.Sprintf("user_tokens:%s", claims.UserID)
	err = m.redisClient.SRem(ctx, userTokenKey, claims.TokenID).Err()
	if err != nil {
		return fmt.Errorf("failed to remove from user token mapping: %w", err)
	}

	return nil
}

// RevokeAllUserTokens Revoke all tokens for a user (force logout from all devices)
func (m *JWTManager) RevokeAllUserTokens(ctx context.Context, userID string) error {
	userTokenKey := fmt.Sprintf("user_tokens:%s", userID)

	// Get all token IDs for the user
	tokenIDs, err := m.redisClient.SMembers(ctx, userTokenKey).Result()
	if err != nil {
		return fmt.Errorf("failed to get user tokens: %w", err)
	}

	// Delete each session
	for _, tokenID := range tokenIDs {
		sessionKey := fmt.Sprintf("session:%s", tokenID)
		m.redisClient.Del(ctx, sessionKey) // Ignore errors
	}

	// Delete the user token mapping
	err = m.redisClient.Del(ctx, userTokenKey).Err()
	if err != nil {
		return fmt.Errorf("failed to delete user token mapping: %w", err)
	}

	return nil
}

// RefreshToken Refresh token (create new token, invalidate old one)
func (m *JWTManager) RefreshToken(
	ctx context.Context,
) (*TokenResponse, error) {
	oldTokenString := m.GetToken(ctx)
	claims, err := m.ValidateToken(ctx, oldTokenString)
	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	// Revoke old token
	err = m.RevokeToken(ctx, oldTokenString)
	if err != nil {
		return nil, fmt.Errorf("failed to revoke old token: %w", err)
	}

	// Create new token
	newToken, err := m.CreateAndStoreToken(ctx, CreatTokenParams{
		UserID:      claims.UserID,
		Role:        claims.Role,
		Roles:       claims.Roles,
		Permissions: claims.Permissions,
		Identifier:  claims.Identifier,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create new token: %w", err)
	}

	return newToken, nil
}

// IsTokenRevoked Check if token is revoked (for middleware)
func (m *JWTManager) IsTokenRevoked(ctx context.Context, tokenID string) (bool, error) {
	sessionKey := fmt.Sprintf("session:%s", tokenID)
	exists, err := m.redisClient.Exists(ctx, sessionKey).Result()
	if err != nil {
		return false, err
	}
	return exists == 0, nil
}

// GetSessionData Get session data from token
func (m *JWTManager) GetSessionData(
	ctx context.Context,
	tokenID string,
) (SessionData, error) {
	sessionKey := fmt.Sprintf("session:%s", tokenID)
	var sessionData SessionData
	res, err := m.redisClient.HGetAll(ctx, sessionKey).Result()
	if err != nil {
		return SessionData{}, err
	}
	err = json.Unmarshal([]byte(res["sessionData"]), &sessionData)
	if err != nil {
		return SessionData{}, err
	}
	return sessionData, nil
}

// userTokenKey is the Redis set holding every token id issued to a user.
func userTokenKey(userID string) string {
	return fmt.Sprintf("user_tokens:%s", userID)
}

// ListSessions returns the sessions currently active for a user, newest first.
// The raw JWT is stripped from every entry: listing sessions must never be a
// way to harvest live tokens.
func (m *JWTManager) ListSessions(ctx context.Context, userID string) ([]SessionData, error) {
	tokenIDs, err := m.redisClient.SMembers(ctx, userTokenKey(userID)).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to get user tokens: %w", err)
	}

	sessions := make([]SessionData, 0, len(tokenIDs))
	for _, tokenID := range tokenIDs {
		session, err := m.GetSessionData(ctx, tokenID)
		if err != nil {
			// The set entry outlives its hash once the TTL lapses; a session we
			// cannot read is a session that is no longer valid, so drop it.
			m.redisClient.SRem(ctx, userTokenKey(userID), tokenID)
			continue
		}
		session.Token = ""
		sessions = append(sessions, session)
	}

	slices.SortFunc(sessions, func(a, b SessionData) int {
		return b.CreatedAt.Compare(a.CreatedAt)
	})
	return sessions, nil
}

// RevokeSession revokes one of a user's sessions by its token id.
func (m *JWTManager) RevokeSession(ctx context.Context, userID, tokenID string) error {
	// Confirm the session belongs to the caller before touching it, otherwise
	// any authenticated user could revoke somebody else's session by guessing
	// an id.
	session, err := m.GetSessionData(ctx, tokenID)
	if err != nil {
		return fmt.Errorf("session not found: %w", err)
	}
	if session.UserID != userID {
		return errors.New("session does not belong to this user")
	}

	if err := m.redisClient.Del(ctx, fmt.Sprintf("session:%s", tokenID)).Err(); err != nil {
		return fmt.Errorf("failed to revoke session: %w", err)
	}
	if err := m.redisClient.SRem(ctx, userTokenKey(userID), tokenID).Err(); err != nil {
		return fmt.Errorf("failed to remove from user token mapping: %w", err)
	}
	return nil
}
