package api

import "errors"

var (
	ErrPayloadNotFound = errors.New("payload not found")
	ErrPayloadInvalid  = errors.New("payload invalid")
)

var (
	ErrSessionNotFound = errors.New("session not found")
	ErrSessionExpired  = errors.New("session expired")
	ErrSessionInvalid  = errors.New("session invalid")
)

var (
	ErrUserInfoConflict          = errors.New("user info conflict")
	ErrUserWrongPassword         = errors.New("user wrong password")
	ErrUserNotFound              = errors.New("user not found")
	ErrUserExceedGroupPermission = errors.New("user exceed the group permission")
)
