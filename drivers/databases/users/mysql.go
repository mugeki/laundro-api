package users

import (
	"laundro-api-ca/business/addresses"
	"laundro-api-ca/business/users"
	repoAddr "laundro-api-ca/drivers/databases/addresses"

	"gorm.io/gorm"
)

type mysqlUsersRepository struct {
	Conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) users.Repository {
	return &mysqlUsersRepository{
		Conn: conn,
	}
}

func (mysqlRepo *mysqlUsersRepository) Register(userData *users.Domain, addressData *addresses.Domain) (users.Domain, error){
	recUser := fromDomain(*userData)
	recAddress := repoAddr.FromDomain(*addressData)
	
	queryString := "street = ? AND postal_code = ? AND city = ? AND province = ?"
	err := mysqlRepo.Conn.First(&recAddress, queryString ,recAddress.Street, recAddress.PostalCode, recAddress.City, recAddress.Province).Error
	if err != nil {
		if err := mysqlRepo.Conn.Create(&recAddress).Error; err != nil {
			return users.Domain{}, err
		}
	}

	recUser.AddressID = recAddress.ID
	err = mysqlRepo.Conn.Create(&recUser).Error
	if err != nil {
		return users.Domain{}, err
	}
	return recUser.toDomain(), nil
}

func (mysqlRepo *mysqlUsersRepository) GetByUsername(username string) (users.Domain, error){
	rec := Users{}
	err := mysqlRepo.Conn.First(&rec, "username = ?", username).Error
	if err != nil {
		return users.Domain{}, err
	}
	return rec.toDomain(), nil
}

func (mysqlRepo *mysqlUsersRepository) GetByID(id uint) (users.Domain, error){
	rec := Users{}
	err := mysqlRepo.Conn.First(&rec, "id = ?", id).Error
	if err != nil {
		return users.Domain{}, err
	}
	return rec.toDomain(), nil
}