package request

import (
	"laundro-api-ca/business/addresses"
	"laundro-api-ca/business/users"
	"laundro-api-ca/controllers/addresses/request"
	"time"
)

type Users struct {			
	Username    string    			`json:"username" valid:"required,minstringlength(6)"`
	Password    string    			`json:"password" valid:"required,minstringlength(6)"`
	Email       string    			`json:"email" valid:"required,email"`
	Fullname    string    			`json:"fullname" valid:"required,minstringlength(3)"`
	DateOfBirth time.Time 			`json:"date_of_birth" valid:"required"`
	PhoneNumber string    			`json:"phone_number" valid:"required,stringlength(8|15)"`
	RoleID      uint    			`json:"role_id" valid:"required"`
	Address		request.Addresses	`json:"address" valid:"required"`
}

type UsersLogin struct{
	Username    string    			`json:"username" valid:"required,minstringlength(6)"`
	Password    string    			`json:"password" valid:"required,minstringlength(6)"`
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