package response

import (
	"laundro-api-ca/business/orders"
	"time"
)

type Orders struct {
	ID                  uint		`json:"id"`			
	CreatedAt           time.Time	`json:"created_at"`		
	UserID              uint		`json:"customer_id"`	
	LaundromatID        uint		`json:"laundromat_id"`	
	LaundromatName      string		`json:"laundromat_name"`	
	PaymentID           int			`json:"payment_id"`
	PaymentGateway      string		`json:"payment_gateway"`	
	ProductID           uint		`json:"product_id"`	
	ProductName         string		`json:"product_name"`	
	EstimatedFinishTime time.Time	`json:"estimated_finish_time"`		
	Weight              int			`json:"weight"`
	TotalPrice          int			`json:"total_price"`
}

func FromDomain(domain orders.Domain) Orders {
	return Orders{
		ID                  : domain.ID,
		CreatedAt           : domain.CreatedAt,
		UserID              : domain.UserID,
		LaundromatID        : domain.LaundromatID,
		LaundromatName      : domain.LaundromatName,
		PaymentID           : domain.PaymentID,
		PaymentGateway      : domain.PaymentGateway,
		ProductID           : domain.ProductID,
		ProductName         : domain.ProductName,
		EstimatedFinishTime : domain.EstimatedFinishTime,
		Weight              : domain.Weight,
		TotalPrice          : domain.TotalPrice,
	}
}

func FromDomainArray(domain []orders.Domain) []Orders{
	res := []Orders{}
	for _, val := range domain{
		res = append(res, FromDomain(val))
	}
	return res
}