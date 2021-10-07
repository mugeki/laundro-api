package request

import (
	"laundro-api-ca/business/products"
)

type Products struct {
	Category    	string      `json:"category" valid:"required"`
	KgLimit  		int         `json:"kg_limit" valid:"required"`
	KgPrice 		int         `json:"kg_price" valid:"required"`
	EstimatedHour 	int		 	`json:"estimated_hour" valid:"required"`
}

func (req *Products) ToDomain() (*products.Domain) {
	return &products.Domain{
		CategoryName  : req.Category,
		KgLimit  	  : req.KgLimit,
		KgPrice 	  : req.KgPrice,
		EstimatedHour : req.EstimatedHour,
	}
}