package request

import (
	"laundro-api-ca/business/addresses"
	"laundro-api-ca/business/users"
	"laundro-api-ca/controllers/addresses/request"
	"time"
)

type Users struct {			
	Username    string    			`json:"username"`
	Password    string    			`json:"password"`
	Email       string    			`json:"email"`
	Fullname    string    			`json:"fullname"`
	DateOfBirth time.Time 			`json:"date_of_birth"`
	PhoneNumber string    			`json:"phone_number"`
	RoleID      uint    			`json:"role_id"`
	Address		request.Addresses	`json:"address"`
}

type UsersLogin struct{
	Username    string   `json:"username"`
	Password    string   `json:"password"`
}

func (req *Users) ToDomain() (*users.Domain, *addresses.Domain) {
	return &users.Domain{
		Username	: req.Username,
		Password	: req.Password,
		Email       : req.Email,
		Fullname    : req.Fullname,
		DateOfBirth : req.DateOfBirth,
		PhoneNumber : req.PhoneNumber,
		RoleID      : req.RoleID,
	}, &addresses.Domain{
		Street		: req.Address.Street,
		PostalCode	: req.Address.PostalCode,
		City		: req.Address.City,
		Province	: req.Address.Province,
	}
}