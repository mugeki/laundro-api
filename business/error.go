package business

import "errors"

var (
	ErrDuplicateData = errors.New("Data already exist")
	ErrInvalidLoginInfo = errors.New("Username or password is invalid")
)