package request

import "laundro-api-ca/business/orders"

type Orders struct {
	LaundromatID uint		`json:"laundromat_id"`		
	PaymentID    int		`json:"payment_id"`
	ProductName  string		`json:"product"`
	Weight       int		`json:"weight"`
}

func (req *Orders) ToDomain() *orders.Domain {
	return &orders.Domain{
		LaundromatID: req.LaundromatID,
		PaymentID:    req.PaymentID,
		ProductName:  req.ProductName,
		Weight:       req.Weight,
	}
}