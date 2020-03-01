package model

import (
	"fmt"
	"io"
	"strconv"
)

// Role - User role
type Role string

// possible role types
const (
	RoleAdmin           Role = "ADMIN"
	RoleUser            Role = "USER"
	RoleUnauthenticated Role = "UNAUTHENTICATED"
)

// AllRole - all possible roles
var AllRole = []Role{
	RoleAdmin,
	RoleUser,
	RoleUnauthenticated,
}

// GetRoleFromValue - returns the role type for the given value
// If not a valid role, anonymous will be sent as default
func GetRoleFromValue(role string) Role {
	r := Role(role)
	switch Role(role) {
	case RoleUser, RoleAdmin:
		return r
	default:
		return RoleUnauthenticated
	}
}

// IsValid - verify if role exists
func (e Role) IsValid() bool {
	switch e {
	case RoleAdmin, RoleUser, RoleUnauthenticated:
		return true
	}
	return false
}

func (e Role) String() string {
	return string(e)
}

// UnmarshalGQL - unmarshal incoming value
func (e *Role) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Role(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Role", str)
	}
	return nil
}

// MarshalGQL - marshal role
func (e Role) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
