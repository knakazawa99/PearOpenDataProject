package types

type AuthType string

const (
	TypeAdmin = AuthType("admin")
	TypeUser  = AuthType("user")
	TypeNone  = AuthType("")
)
