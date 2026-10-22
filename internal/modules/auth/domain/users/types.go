package users

type CreateUserInput struct {
	Name          string         `json:"name"`
	Email         string         `json:"email"`
	Phone         string         `json:"phone"`
	Username      string         `json:"username"`
	PasswordHash  string         `json:"-"` // Never serialize the password hash.
	EmailVerified bool           `json:"email_verified"`
	PhoneVerified bool           `json:"phone_verified"`
	Image         string         `json:"image,omitempty"`
	Status        string         `json:"status"`
	Preference    map[string]any `json:"preference"`
	MFAEnabled    bool           `json:"mfa_enabled"`
	RoleID        int            `json:"role"`
}
