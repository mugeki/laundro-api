package products

import (
	"errors"
	"laundro-api-ca/business/laundromats"
	"laundro-api-ca/business/products"
	laundroRec "laundro-api-ca/drivers/databases/laundromats"

	"gorm.io/gorm"
)

type mysqlProductsRepository struct {
	Conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) products.Repository {
	return &mysqlProductsRepository{
		Conn: conn,
	}
}

func (mysqlRepo *mysqlProductsRepository) Insert(productData *products.Domain) (products.Domain, error){
	rec := FromDomain(*productData)
	if rec.KgLimit < 0 || rec.KgPrice < 0 || rec.EstimatedHour < 0 {
		err := errors.New("Invalid value")
		return products.Domain{}, err
	}
	err := mysqlRepo.Conn.Create(&rec).Error
	if err != nil {
		return products.Domain{}, err
	}
	return rec.toDomain(), nil
}

func (mysqlRepo *mysqlProductsRepository) GetAllByLaundromat(laundroID uint) ([]products.Domain, error){
	rec := []Products{}
	err := mysqlRepo.Conn.Joins("Category").Find(&rec, "laundromat_id = ?", laundroID).Error
	if err != nil {
		return []products.Domain{}, err
	}
	products := toDomainArray(rec)
	return products, nil
}

func (mysqlRepo *mysqlProductsRepository) Update(id uint, productData *products.Domain) (products.Domain, error){
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

func (mysqlRepo *mysqlProductsRepository) Delete(id uint) (string, error){
	rec := Products{}
	err := mysqlRepo.Conn.Delete(&rec, "id = ?",id).Error
	if err != nil {
		return "", err
	}
	return "Product Deleted", nil
}

func (mysqlRepo *mysqlProductsRepository) GetCategoryID(name string) (int, error){
	rec := Category{}
	err := mysqlRepo.Conn.First(&rec, "name = ?", name).Error
	if err != nil {
		return -1, err
	}
	return rec.ID, nil
}

func (mysqlRepo *mysqlProductsRepository) GetCategoryNameByProductID(id uint) string{
	rec := Products{}
	err := mysqlRepo.Conn.Joins("Category").First(&rec, id).Error
	if err != nil {
		return ""
	}
	return rec.Category.Name
}

func (mysqlRepo *mysqlProductsRepository) GetByCategoryName(categoryName string) (products.Domain, error){
	rec := Products{}
	err := mysqlRepo.Conn.Joins("Laundromat").Joins("Category").First(&rec, "Category.name = ?", categoryName).Error
	if err != nil {
		return products.Domain{}, err
	}
	return rec.toDomain(), nil
}

func (mysqlRepo *mysqlProductsRepository) GetLaundromatID(id uint) uint{
	rec := Products{}
	err := mysqlRepo.Conn.First(&rec, "id = ?",id).Error
	if err != nil {
		return 0
	}
	return rec.LaundromatID
}

func (mysqlRepo *mysqlProductsRepository) GetLaundromatByCategory(categoryId int) ([]laundromats.Domain, error){
	recLaundro := []laundroRec.Laundromats{}
	recProduct := []Products{}
	err := mysqlRepo.Conn.Joins("Laundromat").Find(&recProduct, "category_id = ?", categoryId).Error
	if len(recProduct) == 0{
		err = errors.New("Not Found")
		return []laundromats.Domain{}, err
	}

	idArray := []uint{}
	for _, val := range recProduct{
		idArray = append(idArray, val.LaundromatID)
	}

	err = mysqlRepo.Conn.Find(&recLaundro,idArray).Error
	if err != nil {
		return []laundromats.Domain{}, err
	}
	domain := laundroRec.ToDomainArray(recLaundro)
	return domain, nil
}
