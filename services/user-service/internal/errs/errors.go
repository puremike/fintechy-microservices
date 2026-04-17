package errs

import (
	"errors"
)

var (
	// User errors
	ErrUserNotFound       = errors.New("user not found")
	ErrInvalidUserDetails = errors.New("invalid user details")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrFailedToCreateUser = errors.New("failed to create user")

	// Auth errors
	ErrPermissionDenied = errors.New("permission denied")

	// Internal errors
	ErrInternal = errors.New("internal server error")
)
