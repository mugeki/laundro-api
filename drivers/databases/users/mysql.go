package users

import (
	"laundro-api-ca/business/users"

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

func (mysqlRepo *mysqlUsersRepository) Register(userData *users.Domain) (users.Domain, error){
	recUser := fromDomain(*userData)
	err := mysqlRepo.Conn.Create(&recUser).Error
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