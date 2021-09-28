package addresses

import (
	"laundro-api-ca/business/addresses"

	"gorm.io/gorm"
)

type mysqlAddressesRepository struct {
	Conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) addresses.Repository{
	return &mysqlAddressesRepository{
		Conn: conn,
	}
}

func (mysqlRepo *mysqlAddressesRepository) Insert(address *addresses.Domain) (addresses.Domain, error){
	rec := FromDomain(*address)
	queryString := "street = ? AND postal_code = ? AND city = ? AND province = ?"
	err := mysqlRepo.Conn.First(&rec, queryString ,address.Street, address.PostalCode, address.City, address.Province).Error
	if err != nil {
		if err := mysqlRepo.Conn.Create(&rec).Error; err != nil {
			return addresses.Domain{}, err
		}
	}
	return rec.toDomain(), nil
}

func (mysqlRepo *mysqlAddressesRepository) FindByID(id uint) (addresses.Domain, error){
	rec := Addresses{}
	err := mysqlRepo.Conn.First(&rec, id).Error
	if err != nil {
		return addresses.Domain{}, err
	}
	return rec.toDomain(), nil
}

func (mysqlRepo *mysqlAddressesRepository) FindByCity(city string) ([]addresses.Domain, error){
	rec := []Addresses{}
	err := mysqlRepo.Conn.Find(&rec, "city = ?", city).Error
	if err != nil {
		return []addresses.Domain{}, err
	}
	domainAddresses := []addresses.Domain{}
	for _, val := range rec{
		domainAddresses = append(domainAddresses, val.toDomain())
	}
	return domainAddresses, nil
}