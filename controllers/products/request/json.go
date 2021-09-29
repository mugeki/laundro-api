package request

import (
	"laundro-api-ca/business/products"
	"time"
)

type Products struct {
	Category    	string      `json:"category"`
	KgLimit  		int         `json:"kg_limit"`
	KgPrice 		int         `json:"kg_price"`
	EstimatedTime 	time.Time 	`json:"estimated_time"`
}

func (req *Products) ToDomain() (*products.Domain) {
	return &products.Domain{
		CategoryName  : req.Category,
		KgLimit  	  : req.KgLimit,
		KgPrice 	  : req.KgPrice,
		EstimatedTime : req.EstimatedTime,
	}
}