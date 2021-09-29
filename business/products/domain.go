package products

import (
	"time"
)

type Domain struct {
	Id             uint
	KgLimit        int
	KgPrice        int
	EstimatedTime  time.Time
	CategoryID     int
	CategoryName   string
	LaundromatID   uint
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type Service interface{
	Insert(laundroID uint, productData *Domain) (Domain, error)
	GetAllByLaundromat(laundroID uint) ([]Domain, error)
	Update(id uint, productData *Domain) (Domain, error)
	Delete(id uint) (string, error)
	GetLaundromatID(id uint) uint
}

type Repository interface{
	Insert(productData *Domain) (Domain, error)
	GetAllByLaundromat(laundroID uint) ([]Domain, error)
	Update(id uint, productData *Domain) (Domain, error)
	Delete(id uint) (string, error)
	GetCategoryID(name string) (int, error)
	GetLaundromatID(id uint) uint
}