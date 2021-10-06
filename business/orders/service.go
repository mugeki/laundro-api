package orders

import (
	"laundro-api-ca/business"
	"laundro-api-ca/business/laundromats"
	"laundro-api-ca/business/products"
	"time"
)

type ordersService struct {
	orderRepository		 Repository
	productRepository	 products.Repository
	laundroRepository laundromats.Repository
}

func NewOrderService(orderRepo Repository, productRepo products.Repository, laundroRepo laundromats.Repository) Service {
	return &ordersService{
		orderRepository: orderRepo,
		productRepository: productRepo,
		laundroRepository: laundroRepo,
	}
}

func (service *ordersService) Create(userId uint, orderData *Domain) (Domain, error) {
	laundroOpen := service.laundroRepository.GetStatusByID(orderData.LaundromatID)
	if !laundroOpen {
		return Domain{}, business.ErrLaundromatNotAvailable
	}

	orderData.UserID = userId
	productData, err := service.productRepository.GetByCategoryName(orderData.ProductName)
	if err != nil {
		return Domain{}, business.ErrInvalidProductCategory
	}

	if orderData.Weight > productData.KgLimit{
		return Domain{}, business.ErrWeightExceed
	}
	
	orderData.PaymentGateway, err = service.orderRepository.GetPaymentGateway(orderData.PaymentID)
	if err != nil {
		return Domain{}, business.ErrInvalidPayment
	}

	finishedTime := time.Now().Add(time.Hour * time.Duration(productData.EstimatedHour))
	orderData.EstimatedFinishTime = finishedTime
	orderData.TotalPrice = productData.KgPrice * orderData.Weight 
	orderData.ProductID = productData.Id
	orderData.ProductName = productData.CategoryName
	orderData.LaundromatID = productData.LaundromatID
	orderData.LaundromatName = productData.LaundromatName

	res, err := service.orderRepository.Create(orderData)
	if err != nil {
		return Domain{}, business.ErrInternalServer
	}
	return res, nil
}

func (service *ordersService) GetByUserID(userId uint) ([]Domain, error) {
	res, err := service.orderRepository.GetByUserID(userId)
	if err != nil {
		return []Domain{}, business.ErrOrdersNotFound
	}

	for i, val := range res{
		res[i].ProductName = service.productRepository.GetCategoryNameByProductID(val.ProductID)
	}

	return res, nil
}