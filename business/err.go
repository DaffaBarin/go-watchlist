package business

import "errors"

var (
	ErrInternalServer = errors.New("something gone wrong, contact administrator")

	ErrNotFound = errors.New("data not found")

	ErrIDNotFound = errors.New("id not found")

	ErrDuplicateData = errors.New("duplicate data")

	ErrUsernameNotFound = errors.New("username empty")

	ErrEmailNotFound = errors.New("email empty")

	ErrPasswordNotFound = errors.New("password empty")

	ErrWrongPassword = errors.New("password wrong")
)
