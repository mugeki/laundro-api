package laundromats

import (
	"laundro-api-ca/business/laundromats"
	"laundro-api-ca/drivers/databases/addresses"
	"laundro-api-ca/drivers/databases/users"

	"gorm.io/gorm"
)

type Laundromats struct {
	gorm.Model
	Name      string  			  `json:"name" gorm:"unique"`
	OwnerID   uint    			  `json:"owner_id" gorm:"->;<-:create"`
	Owner     users.Users    	  `gorm:"constraint:OnUpdate:NO ACTION,OnDelete:CASCADE;"`
	AddressID uint    			  `json:"address_id"`
	Address   addresses.Addresses `gorm:"constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION;"`
	Status    bool     			  `json:"status" gorm:"default:false"`
}

func (rec *Laundromats) toDomain() laundromats.Domain{
	return laundromats.Domain{
		Id        : rec.ID,
		Name      : rec.Name,
		OwnerID   : rec.OwnerID,
		AddressID : rec.AddressID,
		Status    : rec.Status,
		CreatedAt : rec.CreatedAt,
		UpdatedAt : rec.UpdatedAt,
	}
}

func ToDomainArray(rec []Laundromats) []laundromats.Domain{
	domain := []laundromats.Domain{}

	for _, val := range rec{
		domain = append(domain, val.toDomain())
	}
	return domain
}

func fromDomain(domain laundromats.Domain) *Laundromats{
	return &Laundromats{
		Model: gorm.Model{
			ID: domain.Id,
			CreatedAt: domain.CreatedAt,
			UpdatedAt: domain.UpdatedAt,
		},
		Name      : domain.Name,
		OwnerID   : domain.OwnerID,
		AddressID : domain.AddressID,
		Status    : domain.Status,
	}
}