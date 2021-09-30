package products

import (
	"laundro-api-ca/business/products"
	"laundro-api-ca/drivers/databases/laundromats"

	"gorm.io/gorm"
)

type Category struct {
	ID		int		`gorm:"primaryKey"`
	Name	string
}

type Products struct {
	gorm.Model
	KgLimit        int          			`json:"kg_limit"`
	KgPrice        int          			`json:"kg_price"`
	EstimatedHour  int 			  			`json:"estimated_hour"`
	CategoryID     int          			`json:"category_id"`
	Category       Category 				`gorm:"constraint:OnUpdate:NO ACTION,OnDelete:RESTRICT;"`
	LaundromatID   uint						`json:"laundromat_id"`
	Laundromat	   laundromats.Laundromats  `gorm:"constraint:OnUpdate:NO ACTION,OnDelete:CASCADE;"`
}

func (rec *Products) toDomain() products.Domain{
	return products.Domain{
		Id            : rec.ID,
		KgLimit       : rec.KgLimit,
		KgPrice       : rec.KgPrice,
		EstimatedHour : rec.EstimatedHour,
		CategoryID    : rec.CategoryID,
		CategoryName  : rec.Category.Name,
		CreatedAt     : rec.CreatedAt,
		UpdatedAt     : rec.UpdatedAt,
	}
}

func toDomainArray(rec []Products) []products.Domain{
	domain := []products.Domain{}

	for _, val := range rec{
		domain = append(domain, val.toDomain())
	}
	return domain
}

func FromDomain(domain products.Domain) *Products{
	return &Products{
		Model: gorm.Model{
			ID: domain.Id,
			CreatedAt: domain.CreatedAt,
			UpdatedAt: domain.UpdatedAt,
		},
		KgLimit       : domain.KgLimit,
		KgPrice       : domain.KgPrice,
		EstimatedHour : domain.EstimatedHour,
		Category	  : Category{domain.CategoryID,domain.CategoryName},
		LaundromatID  : domain.LaundromatID,
	}
}