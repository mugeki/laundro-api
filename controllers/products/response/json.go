package response

import (
	"laundro-api-ca/business/products"
	"time"
)

type Products struct {
	ID				uint		`json:"id"`
	Category    	string      `json:"category"`
	KgLimit  		int         `json:"kg_limit"`
	KgPrice 		int         `json:"kg_price"`
	EstimatedTime 	time.Time 	`json:"estimated_time"`
}

func FromDomain(domain products.Domain) Products {
	return Products{
		ID			  : domain.Id,
		Category 	  : domain.CategoryName,
		KgLimit  	  : domain.KgLimit,
		KgPrice 	  : domain.KgPrice,
		EstimatedTime : domain.EstimatedTime,
	}
}

func FromDomainArray(domain []products.Domain) []Products{
	res := []Products{}
	for _, val := range domain{
		res = append(res, FromDomain(val))
	}
	return res
}