package auth

import "time"

type SessionData struct {
	// ID is the token id (jti) the session is keyed by, so a session can be
	// listed and revoked individually.
	ID          string    `json:"id,omitempty"`
	UserID      string    `json:"user_id,omitempty"`
	Identifier  string    `json:"identifier,omitempty"`
	Status      string    `json:"status,omitempty"`
	Role        string    `json:"role,omitempty"`
	Roles       []string  `json:"roles,omitempty"`
	Permissions []string  `json:"permissions,omitempty"`
	Token       string    `json:"token,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	ExpiresAt   time.Time `json:"expires_at"`
}
