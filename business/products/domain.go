package products

import (
	"laundro-api-ca/business/laundromats"
	"time"
)

type Domain struct {
	Id             uint			`json:"id"`
	KgLimit        int			`json:"kg_limit"`
	KgPrice        int			`json:"kg_price"`
	EstimatedHour  int			`json:"estimated_hour"`
	CategoryID     int			`json:"category_id"`
	CategoryName   string		`json:"category_name"`
	LaundromatID   uint			`json:"laundromat_id"`
	LaundromatName string		`json:"laundromat_name"`
	CreatedAt      time.Time	`json:"created_at"`
	UpdatedAt      time.Time	`json:"updated_at"`
}

type Service interface{
	Insert(laundroID uint, productData *Domain) (Domain, error)
	GetAllByLaundromat(laundroID uint) ([]Domain, error)
	Update(id uint, productData *Domain) (Domain, error)
	Delete(id uint) (string, error)
	GetLaundromatID(id uint) uint
	GetLaundromatByCategory(categoryId int) ([]laundromats.Domain, error)
}

type Repository interface{
	Insert(productData *Domain) (Domain, error)
	GetAllByLaundromat(laundroID uint) ([]Domain, error)
	Update(id uint, productData *Domain) (Domain, error)
	Delete(id uint) (string, error)
	GetCategoryID(name string) (int, error)
	GetCategoryNameByProductID(id uint) string
	GetByCategoryName(categoryName string) (Domain, error)
	GetLaundromatID(id uint) uint
	GetLaundromatByCategory(categoryId int) ([]laundromats.Domain, error)
}