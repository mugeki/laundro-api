package laundromats

import (
	"errors"
	"laundro-api-ca/business/laundromats"

	"gorm.io/gorm"
)

type mysqlLaundromatsRepository struct {
	Conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) laundromats.Repository {
	return &mysqlLaundromatsRepository{
		Conn: conn,
	}
}

func (mysqlRepo *mysqlLaundromatsRepository) Insert(laundroData *laundromats.Domain) (laundromats.Domain, error){
	rec := fromDomain(*laundroData)
	err := mysqlRepo.Conn.Create(&rec).Error
	if err != nil {
		return laundromats.Domain{}, err
	}
	return rec.toDomain(), nil
}

func (mysqlRepo *mysqlLaundromatsRepository) GetByAddress(addressID []uint) ([]laundromats.Domain, error){
	rec := []Laundromats{}
	err := mysqlRepo.Conn.Find(&rec, "address_id IN ?", addressID).Error
	if len(rec) == 0{
		err = errors.New("Not Found")
		return nil, err
	}
	laundro := ToDomainArray(rec)
	return laundro, nil
}

func (mysqlRepo *mysqlLaundromatsRepository) GetByName(name string) ([]laundromats.Domain, error){
	rec := []Laundromats{}
	if err := mysqlRepo.Conn.Find(&rec, "name LIKE ?", "%"+name+"%").Error
	err != nil{
		return nil, err
	}
	laundro := ToDomainArray(rec)
	return laundro, nil
}

func (mysqlRepo *mysqlLaundromatsRepository) GetByID(id uint) (laundromats.Domain, error){
	rec := Laundromats{}
	err := mysqlRepo.Conn.First(&rec, "id = ?", id).Error
	if err != nil {
		return laundromats.Domain{}, err
	}
	return rec.toDomain(), nil
}

func (mysqlRepo *mysqlLaundromatsRepository) Update(id uint, laundroData *laundromats.Domain) (laundromats.Domain, error){
	rec := fromDomain(*laundroData)
	recData := *rec
	if err := mysqlRepo.Conn.First(&rec, "id = ?",id).Updates(recData).Error
	err != nil{
		return laundromats.Domain{}, err
	}
	return rec.toDomain(), nil
}

func (mysqlRepo *mysqlLaundromatsRepository) Delete(id uint) (string, error){
	rec := laundromats.Domain{}
	if err := mysqlRepo.Conn.Delete(&rec, "id = ?",id).Error
	err != nil {
		return "", err
	}
	return "Laundromat Deleted", nil
}