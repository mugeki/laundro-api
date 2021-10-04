package laundromats

import (
	"laundro-api-ca/business/addresses"
	"time"
)

type Domain struct {
	Id        uint		`json:"id"`
	Name      string	`json:"name"`
	OwnerID   uint		`json:"owner_id"`
	AddressID uint		`json:"address_id"`
	Status    bool		`json:"status"`
	CreatedAt time.Time	`json:"created_at"`
	UpdatedAt time.Time	`json:"updated_at"`
}

type Service interface{
	Insert(userID uint, laundroData *Domain, addressData *addresses.Domain) (Domain, error)
	GetByIP() ([]Domain, error)
	GetByName(name string) ([]Domain, error)
	GetByID(id uint) (Domain, error)
	Update(id uint, laundroData *Domain, addressData *addresses.Domain) (Domain, error)
	Delete(id uint) (string, error)
}

type Repository interface{
	Insert(laundroData *Domain) (Domain, error)
	GetByAddress(addressID []uint) ([]Domain, error)
	GetByName(name string) ([]Domain, error)
	GetByID(id uint) (Domain, error)
	GetStatusByID(id uint) bool
	Update(id uint, laundroData *Domain) (Domain, error)
	Delete(id uint) (string, error)
}