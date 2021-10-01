package orders

import (
	"laundro-api-ca/business/products"
	"time"
)

type Domain struct {
	ID                  uint
	CreatedAt           time.Time
	UserID              uint
	LaundromatID        uint
	LaundromatName		string
	PaymentID           int
	PaymentGateway		string
	ProductID           uint
	ProductName			string
	EstimatedFinishTime time.Time
	Weight              int
	TotalPrice          int
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