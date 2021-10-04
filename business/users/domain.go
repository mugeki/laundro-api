package users

import (
	"laundro-api-ca/business/addresses"
	"time"
)

type Domain struct {
	Id          uint		`json:"id"`	
	Username    string    	`json:"username"`
	Password    string    	`json:"password"`
	Email       string    	`json:"email"`
	Fullname    string    	`json:"fullname"`
	DateOfBirth time.Time 	`json:"date_of_birth"`
	PhoneNumber string    	`json:"phone_number"`
	RoleID      uint      	`json:"role_id"`
	AddressID   uint      	`json:"address_id"`
	CreatedAt   time.Time	`json:"created_at"`
	UpdatedAt	time.Time	`json:"updated_at"`
}

type Service interface {
	Register(userData *Domain, addressData *addresses.Domain) (Domain, error)
	Login(username, password string) (string, error)
	GetByID(id uint) (Domain, error)
}

type Repository interface {
	Register(userData *Domain) (Domain, error)
	GetByUsername(username string) (Domain, error)
	GetByID(id uint) (Domain, error)
}

