package request

import (
	"laundro-api-ca/business/addresses"
	"laundro-api-ca/business/laundromats"
	"laundro-api-ca/controllers/addresses/request"
)

type Laundromats struct {
	Name    string            `json:"name"`
	Status  int               `json:"status"`
	OwnerID uint              `json:"owner_id"`
	Address request.Addresses `json:"address"`
}

func (req *Laundromats) ToDomain() (*laundromats.Domain, *addresses.Domain) {
	return &laundromats.Domain{
		Name      : req.Name,
		OwnerID   : req.OwnerID,
		Status    : req.Status,
	}, &addresses.Domain{
		Street		: req.Address.Street,
		PostalCode	: req.Address.PostalCode,
		City		: req.Address.City,
		Province	: req.Address.Province,
	}
}