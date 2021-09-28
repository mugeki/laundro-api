package business

import "errors"

var (
	ErrDuplicateData = errors.New("Data already exist")
	ErrInvalidLoginInfo = errors.New("Username or password is invalid")
	ErrInternalServer = errors.New("Something went wrong")
	ErrNearestLaundromatNotFound = errors.New("No laundromat found within your area")
	ErrLaundromatNotFound = errors.New("No laundromat with such name")
	ErrUnauthorized = errors.New("User Unauthorized")
)