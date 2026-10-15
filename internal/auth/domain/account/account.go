// Package account
package account

type AccountIdentifier struct {
	Email    *string
	Phone    *string
	Username *string
}
type Identifier string

const (
	EmailIdentfier    Identifier = "email"
	PhoneIdentfier    Identifier = "phone"
	UsernameIdentfier Identifier = "username"
)

func (a AccountIdentifier) GetIdentifierColumn() Identifier {
	switch {
	case a.Email != nil:
		return EmailIdentfier
	case a.Phone != nil:
		return PhoneIdentfier
	case a.Username != nil:
		return UsernameIdentfier
	default:
		return ""
	}
}

func (a AccountIdentifier) Column() string {
	switch {
	case a.Email != nil && *a.Email != "":
		return string(EmailIdentfier)
	case a.Phone != nil && *a.Phone != "":
		return string(PhoneIdentfier)
	case a.Username != nil && *a.Username != "":
		return string(UsernameIdentfier)
	default:
		return ""
	}
}

func (a AccountIdentifier) Value() string {
	switch {
	case a.Email != nil && *a.Email != "":
		return *a.Email
	case a.Phone != nil && *a.Phone != "":
		return *a.Phone
	case a.Username != nil && *a.Username != "":
		return *a.Username
	default:
		return ""
	}
}
