package orders_test

import (
	_laundroMock "laundro-api-ca/business/laundromats/mocks"
	"laundro-api-ca/business/orders"
	_orderMock "laundro-api-ca/business/orders/mocks"
	"laundro-api-ca/business/products"
	_productMock "laundro-api-ca/business/products/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	mockOrderRepository   _orderMock.Repository
	mockProductRepository _productMock.Repository
	mockLaundroRepository _laundroMock.Repository
	orderService          orders.Service
	orderDomain           orders.Domain
	productDomain         products.Domain
)

func TestMain(m *testing.M) {
	orderService = orders.NewOrderService(&mockOrderRepository, &mockProductRepository, &mockLaundroRepository)
	orderDomain = orders.Domain{
		ID:                  1,
		CreatedAt:           time.Now(),
		UserID:              1,
		LaundromatID:        1,
		PaymentID:           1,
		PaymentGateway:      "Test Gateway",
		ProductID:           1,
		ProductName:         "Test Product",
		EstimatedFinishTime: (time.Now()).Add(4 * time.Hour),
		Weight:              4,
		TotalPrice:          20000,
	}
	productDomain = products.Domain{
		Id:             1,
		KgLimit:        4,
		KgPrice:        5000,
		EstimatedHour:  1,
		CategoryName:   "Test",
		LaundromatName: "Test Laundro",
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	m.Run()
}

func TestCreate(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T) {
		mockLaundroRepository.On("GetStatusByID", mock.AnythingOfType("uint")).Return(true).Once()
		mockProductRepository.On("GetByCategoryName", mock.AnythingOfType("string")).Return(productDomain, nil).Once()
		mockOrderRepository.On("GetPaymentGateway", mock.AnythingOfType("int")).Return("Test Gateway", nil).Once()
		mockOrderRepository.On("Create", mock.Anything).Return(orderDomain, nil).Once()

		input := orders.Domain{
			LaundromatID: 1,
			PaymentID:    1,
			ProductName:  "Test Product",
			Weight:       4,
		}

		resp, err := orderService.Create(1, &input)

		assert.Nil(t, err)
		assert.Equal(t, orderDomain, resp)
	})
	t.Run("Invalid Test | Laundromat Closed", func(t *testing.T) {
		mockLaundroRepository.On("GetStatusByID", mock.AnythingOfType("uint")).Return(false).Once()

		input := orders.Domain{
			LaundromatID: 1,
			PaymentID:    1,
			ProductName:  "Test Product",
			Weight:       4,
		}

		resp, err := orderService.Create(1, &input)

		assert.NotNil(t, err)
		assert.NotEqual(t, orderDomain, resp)
	})
	t.Run("Invalid Test | Laundromat Closed", func(t *testing.T) {
		mockLaundroRepository.On("GetStatusByID", mock.AnythingOfType("uint")).Return(true).Once()
		mockProductRepository.On("GetByCategoryName", mock.AnythingOfType("string")).Return(products.Domain{}, assert.AnError).Once()

		input := orders.Domain{
			LaundromatID: 1,
			PaymentID:    1,
			ProductName:  "Not a Product",
			Weight:       4,
		}

		resp, err := orderService.Create(1, &input)

		assert.NotNil(t, err)
		assert.NotEqual(t, orderDomain, resp)
	})
	t.Run("Invalid Test | Weight Exceeds Limit", func(t *testing.T) {
		mockLaundroRepository.On("GetStatusByID", mock.AnythingOfType("uint")).Return(true).Once()
		mockProductRepository.On("GetByCategoryName", mock.AnythingOfType("string")).Return(productDomain, nil).Once()

		input := orders.Domain{
			LaundromatID: 1,
			PaymentID:    1,
			ProductName:  "Test Product",
			Weight:       5,
		}

		resp, err := orderService.Create(1, &input)

		assert.NotNil(t, err)
		assert.NotEqual(t, orderDomain, resp)
	})
	t.Run("Invalid Test | Invalid Payment", func(t *testing.T) {
		mockLaundroRepository.On("GetStatusByID", mock.AnythingOfType("uint")).Return(true).Once()
		mockProductRepository.On("GetByCategoryName", mock.AnythingOfType("string")).Return(productDomain, nil).Once()
		mockOrderRepository.On("GetPaymentGateway", mock.AnythingOfType("int")).Return("", assert.AnError).Once()

		input := orders.Domain{
			LaundromatID: 1,
			PaymentID:    99,
			ProductName:  "Test Product",
			Weight:       4,
		}

		resp, err := orderService.Create(1, &input)

		assert.NotNil(t, err)
		assert.NotEqual(t, orderDomain, resp)
	})
	t.Run("Invalid Test | Internal Error", func(t *testing.T) {
		mockLaundroRepository.On("GetStatusByID", mock.AnythingOfType("uint")).Return(true).Once()
		mockProductRepository.On("GetByCategoryName", mock.AnythingOfType("string")).Return(productDomain, nil).Once()
		mockOrderRepository.On("GetPaymentGateway", mock.AnythingOfType("int")).Return("Test Gateway", nil).Once()
		mockOrderRepository.On("Create", mock.Anything).Return(orders.Domain{}, assert.AnError).Once()

		input := orders.Domain{
			LaundromatID: 1,
			PaymentID:    1,
			ProductName:  "Test Product",
			Weight:       4,
		}

		resp, err := orderService.Create(1, &input)

		assert.NotNil(t, err)
		assert.NotEqual(t, orderDomain, resp)
	})
}

func TestGetByUserID(t *testing.T){
	t.Run("Valid Test", func(t *testing.T){
		mockOrderRepository.On("GetByUserID", mock.AnythingOfType("uint")).Return([]orders.Domain{orderDomain}, nil).Once()
		mockProductRepository.On("GetCategoryNameByProductID", mock.AnythingOfType("uint")).Return("Test Product", nil).Once()

		resp, err := orderService.GetByUserID(1)

		assert.Nil(t, err)
		assert.Contains(t, resp, orderDomain)
	})
	t.Run("Invalid Test | No Order", func(t *testing.T){
		mockOrderRepository.On("GetByUserID", mock.AnythingOfType("uint")).Return([]orders.Domain{}, assert.AnError).Once()

		resp, err := orderService.GetByUserID(1)

		assert.NotNil(t, err)
		assert.NotContains(t, resp, orderDomain)
	})
}