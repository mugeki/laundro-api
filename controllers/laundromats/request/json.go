package request

import (
	"laundro-api-ca/business/addresses"
	"laundro-api-ca/business/laundromats"
	"laundro-api-ca/controllers/addresses/request"
)

type Laundromats struct {
	Name    string            `json:"name" valid:"required,minstringlength(3)"`
	Status  bool              `json:"status" valid:"-"`
	Address request.Addresses `json:"address" valid:"required"`
}

func (req *Laundromats) ToDomain() (*laundromats.Domain, *addresses.Domain) {
	return &laundromats.Domain{
		Name      : req.Name,
		Status    : req.Status,
	}, &addresses.Domain{
		Street		: req.Address.Street,
		PostalCode	: req.Address.PostalCode,
		City		: req.Address.City,
		Province	: req.Address.Province,
	}
}