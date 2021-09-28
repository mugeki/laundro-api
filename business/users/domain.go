package users

import (
	"laundro-api-ca/business/addresses"
	"time"
)

type Domain struct {
	Id          uint
	Username    string    
	Password    string    
	Email       string    
	Fullname    string    
	DateOfBirth time.Time 
	PhoneNumber string    
	RoleID      uint      
	AddressID   uint      
	CreatedAt   time.Time
	UpdatedAt	time.Time
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

