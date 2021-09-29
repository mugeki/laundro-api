package products

import (
	"laundro-api-ca/business/products"

	"gorm.io/gorm"
)

type mysqlProductRepository struct {
	Conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) products.Repository {
	return &mysqlProductRepository{
		Conn: conn,
	}
}

func (mysqlRepo *mysqlProductRepository) Insert(productData *products.Domain) (products.Domain, error){
	rec := FromDomain(*productData)
	err := mysqlRepo.Conn.Create(&rec).Error
	if err != nil {
		return products.Domain{}, err
	}
	return rec.toDomain(), nil
}

func (mysqlRepo *mysqlProductRepository) GetAllByLaundromat(laundroID uint) ([]products.Domain, error){
	rec := []Products{}
	err := mysqlRepo.Conn.Joins("Category").Find(&rec, "laundromat_id = ?", laundroID).Error
	if err != nil {
		return []products.Domain{}, err
	}
	products := toDomainArray(rec)
	return products, nil
}

func (mysqlRepo *mysqlProductRepository) Update(id uint, productData *products.Domain) (products.Domain, error){
	rec := FromDomain(*productData)
	recData := *rec
	
	mysqlRepo.Conn.First(&rec, "id = ?", id).Association("Category").Replace(&Category{
		productData.CategoryID,
		productData.CategoryName,
	})

	err := mysqlRepo.Conn.Joins("Category").First(&rec, id).Updates(recData).Error
	if err != nil {
		return products.Domain{}, nil
	}
	return rec.toDomain(), nil
}

func (mysqlRepo *mysqlProductRepository) Delete(id uint) (string, error){
	rec := Products{}
	err := mysqlRepo.Conn.Delete(&rec, "id = ?",id).Error
	if err != nil {
		return "", err
	}
	return "Product Deleted", nil
}

func (mysqlRepo *mysqlProductRepository) GetCategoryID(name string) (int, error){
	rec := Category{}
	err := mysqlRepo.Conn.First(&rec, "name = ?", name).Error
	if err != nil {
		return -1, err
	}
	return rec.ID, nil
}

func (mysqlRepo *mysqlProductRepository) GetLaundromatID(id uint) uint{
	rec := Products{}
	err := mysqlRepo.Conn.First(&rec, "id = ?",id).Error
	if err != nil {
		return 0
	}
	return rec.LaundromatID
}