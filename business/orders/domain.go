package orders

import (
	"laundro-api-ca/business/products"
	"time"
)

type Domain struct {
	ID                  uint		`json:"id"`
	CreatedAt           time.Time	`json:"created_at"`
	UserID              uint		`json:"user_id"`
	LaundromatID        uint		`json:"laundromat_id"`
	LaundromatName		string		`json:"laundromat_name"`
	PaymentID           int			`json:"payment_id"`
	PaymentGateway		string		`json:"payment_gateway"`
	ProductID           uint		`json:"product_id"`
	ProductName			string		`json:"product_name"`
	EstimatedFinishTime time.Time	`json:"estimated_finish_time"`
	Weight              int			`json:"weight"`
	TotalPrice          int			`json:"total_price"`
}

type Service interface {
	Create(userId uint, orderData *Domain) (Domain, error)
	GetByUserID(userId uint) ([]Domain, error)
}

type Repository interface {
	Create(orderData *Domain, productData *products.Domain) (Domain, error)
	GetByUserID(userId uint) ([]Domain, error)
	GetPaymentGateway(paymentId int) (string, error)
}