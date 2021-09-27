package users

import (
	"laundro-api-ca/business/users"
	"laundro-api-ca/drivers/databases/addresses"
	"time"

	"gorm.io/gorm"
)

type Roles struct{
	ID 	 int 	`gorm:"primaryKey"`
	Name string
}

type Users struct {
	gorm.Model
	Username    string    			`json:"username" gorm:"unique"`
	Password    string    			`json:"password"`
	Email       string    			`json:"email"`
	Fullname    string    			`json:"fullname"`
	DateOfBirth time.Time 			`json:"dob"`
	PhoneNumber string   			`json:"phone_number"`
	RoleID      uint      			`json:"role_id"`
	Role		Roles				`gorm:"constraint:OnUpdate:NO ACTION,OnDelete:RESTRICT;"`
	AddressID   uint      			`json:"address_id"`
	Address		addresses.Addresses	`gorm:"constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION"`
}

func (rec *Users) toDomain() users.Domain{
	return users.Domain{
		Id          : rec.ID,
		Username    : rec.Username,
		Password    : rec.Password,
		Email       : rec.Email,
		Fullname    : rec.Fullname,
		DateOfBirth : rec.DateOfBirth,
		PhoneNumber : rec.PhoneNumber,
		RoleID      : rec.RoleID,
		AddressID   : rec.AddressID,
		CreatedAt   : rec.CreatedAt,
		UpdatedAt	: rec.UpdatedAt,
	}
}

func fromDomain(domain users.Domain) *Users{
	return &Users{
		Model: gorm.Model{
			ID: domain.Id,
			CreatedAt: domain.CreatedAt,
			UpdatedAt: domain.UpdatedAt,
		},
		Username    : domain.Username,
		Password    : domain.Password,
		Email       : domain.Email,
		Fullname    : domain.Fullname,
		DateOfBirth : domain.DateOfBirth,
		PhoneNumber : domain.PhoneNumber,
		RoleID      : domain.RoleID,
		AddressID   : domain.AddressID,
	}
}