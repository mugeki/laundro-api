package response

import (
	"laundro-api-ca/business/products"
)

type Products struct {
	ID				uint		`json:"id"`
	Category    	string      `json:"category"`
	KgLimit  		int         `json:"kg_limit"`
	KgPrice 		int         `json:"kg_price"`
	EstimatedHour 	int		 	`json:"estimated_time"`
}

func FromDomain(domain products.Domain) Products {
	return Products{
		ID			  : domain.Id,
		Category 	  : domain.CategoryName,
		KgLimit  	  : domain.KgLimit,
		KgPrice 	  : domain.KgPrice,
		EstimatedHour : domain.EstimatedHour,
	}
}

func FromDomainArray(domain []products.Domain) []Products{
	res := []Products{}
	for _, val := range domain{
		res = append(res, FromDomain(val))
	}
	return res
}