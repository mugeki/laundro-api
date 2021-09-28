package laundromats

import (
	"laundro-api-ca/business/addresses"
	"time"
)

type Domain struct {
	Id        uint
	Name      string
	OwnerID   uint
	AddressID uint
	Status    int
	CreatedAt time.Time
	UpdatedAt time.Time
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
	Update(id uint, laundroData *Domain) (Domain, error)
	Delete(id uint) (string, error)
}