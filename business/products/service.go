package products

import (
	"laundro-api-ca/business"
	"laundro-api-ca/business/laundromats"
)

type productService struct {
	productRepository Repository
}

func NewProductService(productRepo Repository) Service {
	return &productService{
		productRepository: productRepo,
	}
}

func (service *productService) Insert(laundroID uint, productData *Domain) (Domain, error) {
	var err error

	productData.LaundromatID = laundroID
	productData.CategoryID, err = service.productRepository.GetCategoryID(productData.CategoryName)
	if err != nil {
		return Domain{}, business.ErrInternalServer
	}
	res, err := service.productRepository.Insert(productData)
	if err != nil {
		return Domain{}, business.ErrInternalServer
	}
	return res, nil
}

func (service *productService) GetAllByLaundromat(laundroID uint) ([]Domain, error) {
	res, err := service.productRepository.GetAllByLaundromat(laundroID)
	if err != nil {
		return []Domain{}, nil
	}
	return res, nil
}

func (service *productService) Update(id uint, productData *Domain) (Domain, error) {
	var err error

	productData.CategoryID, err = service.productRepository.GetCategoryID(productData.CategoryName)
	if err != nil {
		return Domain{}, business.ErrInvalidProductCategory
	}
	
	res, err := service.productRepository.Update(id, productData)
	if err != nil {
		return Domain{}, business.ErrInternalServer
	}
	return res, nil
}

func (service *productService) Delete(id uint) (string, error) {
	res, err := service.productRepository.Delete(id)
	if err != nil {
		return "", business.ErrProductNotFound
	}
	return res, nil
}

func (service *productService) GetLaundromatID(id uint) uint{
	res := service.productRepository.GetLaundromatID(id)
	return res
}

func (service *productService) GetLaundromatByCategory(categoryId int) ([]laundromats.Domain, error){
	res, err := service.productRepository.GetLaundromatByCategory(categoryId)
	if err != nil {
		return []laundromats.Domain{}, business.ErrLaundromatNotFound
	}
	return res, nil
}