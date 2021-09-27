package addresses

import (
	"laundro-api-ca/business/addresses"
)

type Addresses struct {
	ID         uint   `json:"id" gorm:"primaryKey"`
	Street     string `json:"street"`
	PostalCode int    `json:"postal_code"`
	City       string `json:"city"`
	Province   string `json:"province"`
}

func (rec *Addresses) toDomain() addresses.Domain {
	return addresses.Domain{
		ID         : rec.ID,
		Street     : rec.Street,
		PostalCode : rec.PostalCode,
		City       : rec.City,
		Province   : rec.Province,
	}
}

func FromDomain(domain addresses.Domain) *Addresses {
	return &Addresses{
		ID         : domain.ID,
		Street     : domain.Street,
		PostalCode : domain.PostalCode,
		City       : domain.City,
		Province   : domain.Province,
	}
}