package api

import "errors"

//goland:noinspection GoUnusedGlobalVariable
var (
	ErrInvalidParameter = errors.New("invalid parameter")
	ErrInvalidUserName  = errors.New("invalid user name")
	ErrInvalidUserAge   = errors.New("invalid user age")
	ErrInvalidUserID    = errors.New("invalid user ID")
)
