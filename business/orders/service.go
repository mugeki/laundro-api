package orders

import (
	"laundro-api-ca/business"
	"laundro-api-ca/business/products"
	"time"
)

type ordersService struct {
	orderRepository		Repository
	productRepository	products.Repository
}

func NewOrderService(orderRepo Repository, productRepo products.Repository) Service {
	return &ordersService{
		orderRepository: orderRepo,
		productRepository: productRepo,
	}
}

func (service *ordersService) Create(userId uint, orderData *Domain) (Domain, error) {
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

	res, err := service.orderRepository.Create(orderData, &productData)
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