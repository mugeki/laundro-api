package laundromats

import (
	"laundro-api-ca/business/addresses"
	"laundro-api-ca/business/laundromats"
	repoAddr "laundro-api-ca/drivers/databases/addresses"

	"gorm.io/gorm"
)

type mysqlLaundromatsRepository struct {
	Conn *gorm.DB
}

// func NewMySQLRepository(conn *gorm.DB) laundromats.Repository {
// 	return &mysqlLaundromatsRepository{
// 		Conn: conn,
// 	}
// }

func (mysqlRepo *mysqlLaundromatsRepository) Insert(laundroData *laundromats.Domain, addressData *addresses.Domain) (laundromats.Domain, error){
	recLaundro := fromDomain(*laundroData)
	recAddress := repoAddr.FromDomain(*addressData)
	
	queryString := "street = ? AND postal_code = ? AND city = ? AND province = ?"
	err := mysqlRepo.Conn.First(&recAddress, queryString ,recAddress.Street, recAddress.PostalCode, recAddress.City, recAddress.Province).Error
	if err != nil {
		if err := mysqlRepo.Conn.Create(&recAddress).Error; err != nil {
			return laundromats.Domain{}, err
		}
	}

	recLaundro.AddressID = recAddress.ID
	err = mysqlRepo.Conn.Create(&recLaundro).Error
	if err != nil {
		return laundromats.Domain{}, err
	}
	return recLaundro.toDomain(), nil
}

// func (mysqlRepo *mysqlLaundromatsRepository) GetLaundromatsByIP() ([]laundromats.Domain, error){

// }

// func (mysqlRepo *mysqlLaundromatsRepository) GetLaundromatsByName(name string) ([]laundromats.Domain, error){

// }

// func (mysqlRepo *mysqlLaundromatsRepository) UpdateLaundromat(id uint, laundroData *laundromats.Domain, addressData *addresses.Domain) (laundromats.Domain, error){

// }

// func (mysqlRepo *mysqlLaundromatsRepository) DeleteLaundromat(id uint) (string, error){

// }