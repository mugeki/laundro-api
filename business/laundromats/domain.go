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
	Insert(laundroData *Domain, addressData *addresses.Domain) (Domain, error)
	GetLaundromatsByIP() ([]Domain, error)
	GetLaundromatsByName(laundroData *Domain) ([]Domain, error)
	UpdateLaundromat(id uint, laundroData *Domain, addressData *addresses.Domain) (Domain, error)
	DeleteLaundromat(id uint) (string, error)
}

type Repository interface{
	Insert(laundroData *Domain, addressData *addresses.Domain) (Domain, error)
	GetLaundromatsByIP() ([]Domain, error)
	GetLaundromatsByName(name string) ([]Domain, error)
	UpdateLaundromat(id uint, laundroData *Domain, addressData *addresses.Domain) (Domain, error)
	DeleteLaundromat(id uint) (string, error)
}