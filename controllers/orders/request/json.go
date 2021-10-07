package request

import "laundro-api-ca/business/orders"

type Orders struct {
	LaundromatID uint		`json:"laundromat_id" valid:"required"`		
	PaymentID    int		`json:"payment_id" valid:"required"`
	ProductName  string		`json:"product_name" valid:"required"`
	Weight       int		`json:"weight" valid:"required"`
}

func (req *Orders) ToDomain() *orders.Domain {
	return &orders.Domain{
		LaundromatID: req.LaundromatID,
		PaymentID:    req.PaymentID,
		ProductName:  req.ProductName,
		Weight:       req.Weight,
	}
}