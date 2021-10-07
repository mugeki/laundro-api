package products_test

import (
	"laundro-api-ca/business"
	"laundro-api-ca/business/laundromats"
	"laundro-api-ca/business/products"
	_productMock "laundro-api-ca/business/products/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var(
	mockProductRepository _productMock.Repository
	productService	products.Service
	productDomain	products.Domain
	productNewDomain products.Domain
	laundroDomain  laundromats.Domain
)

func TestMain(m *testing.M){
	productService = products.NewProductService(&mockProductRepository)
	productDomain = products.Domain{
		Id             : 1,
		KgLimit        : 4,
		KgPrice        : 5000,
		EstimatedHour  : 1,
		CategoryName   : "Test",
		LaundromatName : "Test Laundro",
		CreatedAt      : time.Now(),
		UpdatedAt      : time.Now(),
	}
	productNewDomain = products.Domain{
		Id             : 1,
		KgLimit        : 6,
		KgPrice        : 6000,
		EstimatedHour  : 2,
		CategoryName   : "Test New",
		LaundromatName : "Test Laundro",
		CreatedAt      : time.Now(),
		UpdatedAt      : time.Now(),
	}
	laundroDomain = laundromats.Domain{
		Id        : 1,
		Name      : "Test Laundro",
		OwnerID   : 1,
		AddressID : 1,
		Status    : true,
		CreatedAt : time.Now(),
		UpdatedAt : time.Now(),
	}
	m.Run()
}

func TestInsert(t *testing.T){
	t.Run("Valid Test", func(t *testing.T){
		mockProductRepository.On("GetCategoryID", mock.AnythingOfType("string")).Return(1, nil).Once()
		mockProductRepository.On("Insert", mock.Anything).Return(productDomain, nil).Once()

		input := products.Domain{
			KgLimit        : 4,
			KgPrice        : 5000,
			EstimatedHour  : 1,
			CategoryName   : "Test",
		}

		resp, err := productService.Insert(1, &input)
		
		assert.Nil(t, err)
		assert.Equal(t, productDomain, resp)
	})
	t.Run("Invalid Test | Invalid Category", func(t *testing.T){
		mockProductRepository.On("GetCategoryID", mock.AnythingOfType("string")).Return(-1, business.ErrInvalidProductCategory).Once()

		input := products.Domain{
			KgLimit        : 4,
			KgPrice        : 5000,
			EstimatedHour  : 1,
			CategoryName   : "Not a Category",
		}

		resp, err := productService.Insert(1, &input)
		
		assert.NotNil(t, err)
		assert.NotEqual(t, productDomain, resp)
	})
	t.Run("Invalid Test | Invalid Value", func(t *testing.T){
		mockProductRepository.On("GetCategoryID", mock.AnythingOfType("string")).Return(1, nil).Once()
		mockProductRepository.On("Insert", mock.Anything).Return(products.Domain{}, business.ErrInternalServer).Once()

		input := products.Domain{
			KgLimit        : -4,
			KgPrice        : -5000,
			EstimatedHour  : -1,
			CategoryName   : "Test",
		}

		resp, err := productService.Insert(1, &input)
		
		assert.NotNil(t, err)
		assert.NotEqual(t, productDomain, resp)
	})
}

func TestGetAllByLaundromat(t *testing.T){
	t.Run("Valid Test", func(t *testing.T){
		mockProductRepository.On("GetAllByLaundromat", mock.AnythingOfType("uint")).Return([]products.Domain{productDomain}, nil).Once()

		resp, err := productService.GetAllByLaundromat(1)

		assert.Nil(t, err)
		assert.Contains(t, resp, productDomain)
	})
	t.Run("Invalid Test | Invalid Laundromat/ No products", func(t *testing.T){
		mockProductRepository.On("GetAllByLaundromat", mock.AnythingOfType("uint")).Return([]products.Domain{}, business.ErrProductNotFound).Once()

		resp, err := productService.GetAllByLaundromat(2)

		assert.NotNil(t, err)
		assert.NotContains(t, resp, productDomain)
	})
}

func TestUpdate(t *testing.T){
	t.Run("Valid Test", func(t *testing.T){
		mockProductRepository.On("GetCategoryID", mock.AnythingOfType("string")).Return(1, nil).Once()
		mockProductRepository.On("Update", mock.AnythingOfType("uint"), mock.Anything).Return(productNewDomain, nil).Once()

		input := products.Domain{
			KgLimit        : 6,
			KgPrice        : 6000,
			EstimatedHour  : 2,
			CategoryName   : "Test New",
		}

		resp, err := productService.Update(1, &input)

		assert.Nil(t, err)
		assert.Equal(t, productNewDomain, resp)
	})
	t.Run("Invalid Test | Invalid New Category", func(t *testing.T){
		mockProductRepository.On("GetCategoryID", mock.AnythingOfType("string")).Return(-1, business.ErrInvalidProductCategory).Once()

		input := products.Domain{
			KgLimit        : 6,
			KgPrice        : 6000,
			EstimatedHour  : 2,
			CategoryName   : "Not a Category",
		}

		resp, err := productService.Update(1, &input)

		assert.NotNil(t, err)
		assert.NotEqual(t, productDomain, resp)
	})
	t.Run("Invalid Test | Invalid LaundromatID", func(t *testing.T){
		mockProductRepository.On("GetCategoryID", mock.AnythingOfType("string")).Return(1, nil).Once()
		mockProductRepository.On("Update", mock.AnythingOfType("uint"), mock.Anything).Return(products.Domain{}, business.ErrInternalServer).Once()

		input := products.Domain{
			KgLimit        : 6,
			KgPrice        : 6000,
			EstimatedHour  : 2,
			CategoryName   : "Test New",
		}

		resp, err := productService.Update(1, &input)

		assert.NotNil(t, err)
		assert.NotEqual(t, productDomain, resp)
	})
}

func TestDelete(t *testing.T){
	t.Run("Valid Test", func(t *testing.T){
		mockProductRepository.On("Delete", mock.AnythingOfType("uint")).Return("Product Deleted", nil).Once()

		resp, err := productService.Delete(1)

		assert.Nil(t, err)
		assert.Equal(t, "Product Deleted", resp)
	})
	t.Run("Invalid Test | Invalid LaundromatID", func(t *testing.T){
		mockProductRepository.On("Delete", mock.AnythingOfType("uint")).Return("", business.ErrProductNotFound).Once()

		resp, err := productService.Delete(2)

		assert.NotNil(t, err)
		assert.Equal(t, "", resp)
	})
}

func TestGetLaundromatID(t *testing.T){
	t.Run("Valid Test", func(t *testing.T){
		mockProductRepository.On("GetLaundromatID", mock.AnythingOfType("uint")).Return(uint(1)).Once()

		resp := productService.GetLaundromatID(1)

		assert.Equal(t, uint(1), resp)
	})
}

func TestGetLaundromatByCategory(t *testing.T){
	t.Run("Valid Test", func(t *testing.T){
		mockProductRepository.On("GetLaundromatByCategory", mock.AnythingOfType("int")).Return([]laundromats.Domain{laundroDomain}, nil).Once()
		
		resp, err := productService.GetLaundromatByCategory(1)

		assert.Nil(t, err)
		assert.Contains(t, resp, laundroDomain)
	})
	t.Run("Invalid Test | Invalid Category", func(t *testing.T){
		mockProductRepository.On("GetLaundromatByCategory", mock.AnythingOfType("int")).Return([]laundromats.Domain{}, business.ErrLaundromatNotFound).Once()
		
		resp, err := productService.GetLaundromatByCategory(2)

		assert.NotNil(t, err)
		assert.NotContains(t, resp, laundroDomain)
	})
}