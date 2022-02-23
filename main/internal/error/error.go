package ecoerror

import "errors"

var (
	// ErrNoRecord no record found in database user
	ErrNoRecord = errors.New("no matching record found")
	// ErrInvalidPassword invalid password error
	ErrInvalidPassword = errors.New("invalid password")
	// ErrDuplicateEmail duplicate email error
	ErrDuplicateEmail = errors.New("invalid email: duplicated")
	// ErrDuplicateNickname duplicate nickname error
	ErrDuplicateNickname = errors.New("invalid nickname: duplicated")
	// ErrInactiveAccount inactive account error
	ErrInactiveAccount = errors.New("inactive account")
	// ErrTypeConvertFailed failure while converting type
	ErrTypeConvertFailed = errors.New("type convert failed")
	// ErrInvalidHeader invalid header
	ErrInvalidHeader = errors.New("invalid header")
	// ErrInvalidQuery no query found
	ErrInvalidQuery = errors.New("invalid query")
	// ErrInvalidJson no query found
	ErrInvalidJson = errors.New("invalid json")
	// ErrNoAuthorization no authorization error
	ErrNoAuthorization = errors.New("invalid token")
	// ErrDuplicateData duplicate data
	ErrDuplicateData = errors.New("duplicated data")
	// ErrDuplicateModelName duplicate model name
	ErrDuplicateModelName = errors.New("duplicated model name")
	// ErrPermissionDenied permission denied
	ErrPermissionDenied = errors.New("permission denied")
)
