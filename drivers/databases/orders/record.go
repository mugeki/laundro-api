package orders

import (
	"laundro-api-ca/business/orders"
	"laundro-api-ca/drivers/databases/laundromats"
	"laundro-api-ca/drivers/databases/products"
	"laundro-api-ca/drivers/databases/users"
	"time"

	"gorm.io/gorm"
)

type Payment struct {
	ID		int `gorm:"primaryKey"`
	Gateway	string
}

type Orders struct {
	ID					uint					`gorm:"primaryKey"`
	CreatedAt			time.Time				`json:"created_at"`
	DeletedAt			gorm.DeletedAt			`gorm:"index"`
	UserID              uint					`json:"customer_id"`
	User				users.Users				`gorm:"constraint:OnUpdate:NO ACTION,OnDelete:CASCADE;"`
	LaundromatID        uint					`json:"laundromat_id"`
	Laundromat			laundromats.Laundromats	`gorm:"constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION;"`
	PaymentID           int						`json:"payment_id"`
	Payment				Payment					`gorm:"constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION;"`
	ProductID           uint					`json:"product_id"`
	Product				products.Products		`gorm:"constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION;"`
	EstimatedFinishTime time.Time				`json:"estimated_finish_time"`
	Weight              int						`json:"weight"`
	TotalPrice          int						`jsonn:"total_price"`
}

func (rec *Orders) toDomain() orders.Domain{
	return orders.Domain{
		ID                  : rec.ID,
		CreatedAt           : rec.CreatedAt,
		UserID              : rec.UserID,
		LaundromatID        : rec.LaundromatID,
		LaundromatName		: rec.Laundromat.Name,
		PaymentID           : rec.PaymentID,
		PaymentGateway		: rec.Payment.Gateway,
		ProductID           : rec.ProductID,
		ProductName			: rec.Product.Category.Name,
		EstimatedFinishTime : rec.EstimatedFinishTime,
		Weight              : rec.Weight,
		TotalPrice          : rec.TotalPrice,
	}
}

func toDomainArray(rec []Orders) []orders.Domain{
	domain := []orders.Domain{}

	for _, val := range rec{
		domain = append(domain, val.toDomain())
	}
	return domain
}

func fromDomain(domain orders.Domain) *Orders{
	return &Orders{
		ID					: domain.ID,
		CreatedAt			: domain.CreatedAt,
		UserID              : domain.UserID,
		Laundromat			: laundromats.Laundromats{
								Model: gorm.Model{
									ID: domain.LaundromatID,
								},
								Name: domain.LaundromatName,
							},
		Payment				: Payment{domain.PaymentID, domain.PaymentGateway},
		ProductID			: domain.ProductID,
		Product				: products.Products{
								Model: gorm.Model{
									ID: domain.ProductID,
								},
								Category: products.Category{
									Name: domain.ProductName,
								},
							},
		EstimatedFinishTime : domain.EstimatedFinishTime,
		Weight              : domain.Weight,
		TotalPrice          : domain.TotalPrice,
	}
}