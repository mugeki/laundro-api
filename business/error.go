package business

import "errors"

var (
	ErrDuplicateData = errors.New("Data already exist")
	ErrInvalidLoginInfo = errors.New("Username or password is invalid")
	ErrUserNotFound = errors.New("User not found")
	ErrInternalServer = errors.New("Something went wrong")
	ErrNearestLaundromatNotFound = errors.New("No laundromat found within your area")
	ErrLaundromatNotFound = errors.New("Laundromat not found")
	ErrLaundromatNotAvailable = errors.New("Laundromat not available")
	ErrUnauthorized = errors.New("User Unauthorized")
	ErrProductNotFound = errors.New("Product not found")
	ErrInvalidProductCategory = errors.New("Invalid category")
	ErrOrdersNotFound = errors.New("No order has made")
	ErrWeightExceed = errors.New("Weight limit exceeded")
	ErrInvalidPayment = errors.New("Invalid payment method")
)