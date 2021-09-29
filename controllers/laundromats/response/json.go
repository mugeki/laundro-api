package response

import (
	"laundro-api-ca/business/laundromats"
)

type Laundromats struct {
	Name      string  `json:"name"`
	Status    bool     `json:"status"`
	OwnerID   uint    `json:"owner_id"`
	AddressID uint    `json:"address"`
}

func FromDomain(domain laundromats.Domain) Laundromats {
	return Laundromats{
		Name      : domain.Name,
		Status    : domain.Status,
		OwnerID   : domain.OwnerID,
		AddressID : domain.AddressID,
	}
}