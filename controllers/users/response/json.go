package response

import (
	"laundro-api-ca/business/users"
	"time"
)

type Users struct {
	ID          uint      `json:"id"`
	Username    string    `json:"username"`
	Password    string    `json:"password"`
	Email       string    `json:"email"`
	Fullname    string    `json:"fullname"`
	DateOfBirth time.Time `json:"date_of_birth"`
	PhoneNumber string    `json:"phone_number"`
	RoleID      uint      `json:"role_id"`
	AddressID   uint      `json:"address_id"`
}

func FromDomain(domain users.Domain) Users {
	return Users{
		ID:          domain.Id,
		Username:    domain.Username,
		Password:    domain.Password,
		Email:       domain.Email,
		Fullname:    domain.Fullname,
		DateOfBirth: domain.DateOfBirth,
		PhoneNumber: domain.PhoneNumber,
		RoleID:      domain.RoleID,
		AddressID:   domain.AddressID,
	}
}