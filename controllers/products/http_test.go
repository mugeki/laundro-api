package products_test

import (
	"encoding/json"
	"laundro-api-ca/business"
	"laundro-api-ca/business/laundromats"
	_product "laundro-api-ca/business/products"
	_productMock "laundro-api-ca/business/products/mocks"
	controller "laundro-api-ca/controllers"
	"laundro-api-ca/controllers/products"
	"laundro-api-ca/controllers/products/response"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var(
	mockProductService _productMock.Service
	productCtrl *products.ProductController
	productReq string
	productReqInvalidBind string
	productReqInvalidStruct string
	productResp response.Products
	productDomain _product.Domain
	laundroDomain laundromats.Domain
)

func TestMain(m *testing.M){
	productCtrl = products.NewProductController(&mockProductService)
	productReq = `{
		"category": "Test Category",
		"kg_limit": 4,
		"kg_price": 5000,
		"estimated_hour": 2
	}`
	productReqInvalidBind = `{
		"category": "Test Category",
		"kg_limit": 4,
		"kg_price": 5000
		"estimated_hour": 2
	}`
	productReqInvalidStruct = `{
			"category": "Test Category",
			"kg_limit": 4,
			"kg_price": 5000
		}`
	productResp = response.Products{
		ID				: 1,
		Category    	: "Test Category",
		KgLimit  		: 4,
		KgPrice 		: 5000,
		EstimatedHour 	: 2,
	}
	productDomain = _product.Domain{
		Id             : 1,
		KgLimit        : 4,
		KgPrice        : 5000,
		EstimatedHour  : 2,
		CategoryID     : 1,
		CategoryName   : "Test Category",
		LaundromatID   : 1,
		LaundromatName : "Test Laundry",
		CreatedAt      : time.Now(),
		UpdatedAt      : time.Now(),
	}
	laundroDomain = laundromats.Domain{
		Id        : 1,
		Name      : "Test Laundry",
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
		req := httptest.NewRequest(http.MethodPost, "/products/to/1", strings.NewReader(productReq))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		c := e.NewContext(req,rec)
		mockProductService.On("Insert", mock.AnythingOfType("uint"), mock.Anything).Return(productDomain, nil).Once()
		
		resp := controller.BaseResponse{}
		resp.Meta.Status = http.StatusOK
		resp.Meta.Message = "Success"
		resp.Data = productResp
		expected, _ := json.Marshal(resp)

		if assert.NoError(t, productCtrl.Insert(c)){
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.JSONEq(t, string(expected), rec.Body.String())
		}
	})
	t.Run("Invalid Test | Bind Error", func(t *testing.T){
		req := httptest.NewRequest(http.MethodPost, "/products/to/1", strings.NewReader(productReqInvalidBind))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		c := e.NewContext(req,rec)
		
		resp := controller.BaseResponse{}
		resp.Meta.Status = http.StatusBadRequest
		resp.Meta.Message = "Error"
		resp.Meta.Messages = []string{"code=400, message=Syntax error: offset=72, error=invalid character '\"' after object key:value pair, internal=invalid character '\"' after object key:value pair"}
		expected, _ := json.Marshal(resp)

		if assert.NoError(t, productCtrl.Insert(c)){
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.JSONEq(t, string(expected), rec.Body.String())
		}
	})
	t.Run("Invalid Test | Invalid Struct", func(t *testing.T){
		req := httptest.NewRequest(http.MethodPost, "/products/to/1", strings.NewReader(productReqInvalidStruct))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		c := e.NewContext(req,rec)
		
		resp := controller.BaseResponse{}
		resp.Meta.Status = http.StatusBadRequest
		resp.Meta.Message = "Error"
		resp.Meta.Messages = []string{"estimated_hour: non zero value required"}
		expected, _ := json.Marshal(resp)

		if assert.NoError(t, productCtrl.Insert(c)){
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.JSONEq(t, string(expected), rec.Body.String())
		}
	})
	t.Run("Invalid Test | Insert Failed", func(t *testing.T){
		req := httptest.NewRequest(http.MethodPost, "/products/to/1", strings.NewReader(productReq))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		c := e.NewContext(req,rec)
		mockProductService.On("Insert", mock.AnythingOfType("uint"), mock.Anything).Return(_product.Domain{}, assert.AnError).Once()
		
		resp := controller.BaseResponse{}
		resp.Meta.Status = http.StatusNotFound
		resp.Meta.Message = "Error"
		resp.Meta.Messages = []string{assert.AnError.Error()}
		expected, _ := json.Marshal(resp)

		if assert.NoError(t, productCtrl.Insert(c)){
			assert.Equal(t, http.StatusNotFound, rec.Code)
			assert.JSONEq(t, string(expected), rec.Body.String())
		}
	})
}

func TestGetAllByLaundromat(t *testing.T){
	t.Run("Valid Test", func(t *testing.T){
		req := httptest.NewRequest(http.MethodGet, "/products/from/1", nil)
		rec := httptest.NewRecorder()
		e := echo.New()
		c := e.NewContext(req,rec)
		mockProductService.On("GetAllByLaundromat", mock.AnythingOfType("uint")).Return([]_product.Domain{productDomain}, nil).Once()

		resp := controller.BaseResponse{}
		resp.Meta.Status = http.StatusOK
		resp.Meta.Message = "Success"
		resp.Data = []response.Products{productResp}
		expected, _ := json.Marshal(resp)

		if assert.NoError(t, productCtrl.GetAllByLaundromat(c)){
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.JSONEq(t, string(expected), rec.Body.String())
		}
	})
	t.Run("Invalid Test | Product Not Found", func(t *testing.T){
		req := httptest.NewRequest(http.MethodGet, "/products/from/1", nil)
		rec := httptest.NewRecorder()
		e := echo.New()
		c := e.NewContext(req,rec)
		mockProductService.On("GetAllByLaundromat", mock.AnythingOfType("uint")).Return([]_product.Domain{}, business.ErrProductNotFound).Once()

		resp := controller.BaseResponse{}
		resp.Meta.Status = http.StatusNotFound
		resp.Meta.Message = "Error"
		resp.Meta.Messages = []string{business.ErrProductNotFound.Error()}
		expected, _ := json.Marshal(resp)

		if assert.NoError(t, productCtrl.GetAllByLaundromat(c)){
			assert.Equal(t, http.StatusNotFound, rec.Code)
			assert.JSONEq(t, string(expected), rec.Body.String())
		}
	})
}

func TestUpdate(t *testing.T){
	t.Run("Valid Test", func(t *testing.T){
		req := httptest.NewRequest(http.MethodPut, "/products/edit/1", strings.NewReader(productReq))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		c := e.NewContext(req,rec)
		mockProductService.On("Update", mock.AnythingOfType("uint"), mock.Anything).Return(productDomain, nil).Once()
		
		resp := controller.BaseResponse{}
		resp.Meta.Status = http.StatusOK
		resp.Meta.Message = "Success"
		resp.Data = productResp
		expected, _ := json.Marshal(resp)

		if assert.NoError(t, productCtrl.Update(c)){
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.JSONEq(t, string(expected), rec.Body.String())
		}
	})
	t.Run("Invalid Test | Bind Error", func(t *testing.T){
		req := httptest.NewRequest(http.MethodPut, "/products/edit/1", strings.NewReader(productReqInvalidBind))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		c := e.NewContext(req,rec)
		
		resp := controller.BaseResponse{}
		resp.Meta.Status = http.StatusBadRequest
		resp.Meta.Message = "Error"
		resp.Meta.Messages = []string{"code=400, message=Syntax error: offset=72, error=invalid character '\"' after object key:value pair, internal=invalid character '\"' after object key:value pair"}
		expected, _ := json.Marshal(resp)

		if assert.NoError(t, productCtrl.Update(c)){
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.JSONEq(t, string(expected), rec.Body.String())
		}
	})
	t.Run("Invalid Test | Invalid Struct", func(t *testing.T){
		req := httptest.NewRequest(http.MethodPut, "/products/edit/1", strings.NewReader(productReqInvalidStruct))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		c := e.NewContext(req,rec)
		
		resp := controller.BaseResponse{}
		resp.Meta.Status = http.StatusBadRequest
		resp.Meta.Message = "Error"
		resp.Meta.Messages = []string{"estimated_hour: non zero value required"}
		expected, _ := json.Marshal(resp)

		if assert.NoError(t, productCtrl.Update(c)){
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.JSONEq(t, string(expected), rec.Body.String())
		}
	})
	t.Run("Invalid Test | Internal Server Error", func(t *testing.T){
		req := httptest.NewRequest(http.MethodPut, "/products/edit/1", strings.NewReader(productReq))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		c := e.NewContext(req,rec)
		mockProductService.On("Update", mock.AnythingOfType("uint"), mock.Anything).Return(_product.Domain{}, assert.AnError).Once()
		
		resp := controller.BaseResponse{}
		resp.Meta.Status = http.StatusNotFound
		resp.Meta.Message = "Error"
		resp.Meta.Messages = []string{assert.AnError.Error()}
		expected, _ := json.Marshal(resp)

		if assert.NoError(t, productCtrl.Update(c)){
			assert.Equal(t, http.StatusNotFound, rec.Code)
			assert.JSONEq(t, string(expected), rec.Body.String())
		}
	})
}

func TestDelete(t *testing.T){
	t.Run("Valid Test", func(t *testing.T){
		req := httptest.NewRequest(http.MethodDelete, "/products/1", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		c := e.NewContext(req,rec)

		mockProductService.On("Delete", mock.AnythingOfType("uint")).Return("Laundromat Deleted", nil).Once()

		resp := controller.BaseResponse{}
		resp.Meta.Status = http.StatusOK
		resp.Meta.Message = "Success"
		resp.Data = "Laundromat Deleted"
		expected, _ := json.Marshal(resp)

		if assert.NoError(t, productCtrl.Delete(c)){
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.JSONEq(t, string(expected), rec.Body.String())
		}
	})
	t.Run("Invalid Test | Product Not Found", func(t *testing.T){
		req := httptest.NewRequest(http.MethodDelete, "/products/1", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		c := e.NewContext(req,rec)

		mockProductService.On("Delete", mock.AnythingOfType("uint")).Return("", business.ErrProductNotFound).Once()

		resp := controller.BaseResponse{}
		resp.Meta.Status = http.StatusNotFound
		resp.Meta.Message = "Error"
		resp.Meta.Messages = []string{business.ErrProductNotFound.Error()}
		expected, _ := json.Marshal(resp)

		if assert.NoError(t, productCtrl.Delete(c)){
			assert.Equal(t, http.StatusNotFound, rec.Code)
			assert.JSONEq(t, string(expected), rec.Body.String())
		}
	})
}

func TestGetLaundromatID(t *testing.T){
	t.Run("Valid Test", func(t *testing.T){
		mockProductService.On("GetLaundromatID", mock.AnythingOfType("uint")).Return(uint(1)).Once()
		
		resp := productCtrl.GetLaundromatID(uint(1))

		assert.Equal(t, laundroDomain.Id, resp)
	})
}

func TestGetLaundromatByCategory(t *testing.T){
	t.Run("Valid Test", func(t *testing.T){
		req := httptest.NewRequest(http.MethodGet, "/laundro/find-category/1", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		c := e.NewContext(req,rec)

		mockProductService.On("GetLaundromatByCategory", mock.AnythingOfType("int")).Return([]laundromats.Domain{laundroDomain}, nil).Once()

		resp := controller.BaseResponse{}
		resp.Meta.Status = http.StatusOK
		resp.Meta.Message = "Success"
		resp.Data = []laundromats.Domain{laundroDomain}
		expected, _ := json.Marshal(resp)

		if assert.NoError(t, productCtrl.GetLaundromatByCategory(c)){
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.JSONEq(t, string(expected), rec.Body.String())
		}
	})
	t.Run("Invalid Test | Laundromat Not Found", func(t *testing.T){
		req := httptest.NewRequest(http.MethodGet, "/laundro/find-category/1", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		c := e.NewContext(req,rec)

		mockProductService.On("GetLaundromatByCategory", mock.AnythingOfType("int")).Return([]laundromats.Domain{}, business.ErrLaundromatNotFound).Once()

		resp := controller.BaseResponse{}
		resp.Meta.Status = http.StatusNotFound
		resp.Meta.Message = "Error"
		resp.Meta.Messages = []string{business.ErrLaundromatNotFound.Error()}
		expected, _ := json.Marshal(resp)

		if assert.NoError(t, productCtrl.GetLaundromatByCategory(c)){
			assert.Equal(t, http.StatusNotFound, rec.Code)
			assert.JSONEq(t, string(expected), rec.Body.String())
		}
	})
}