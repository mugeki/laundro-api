package orders_test

import (
	"encoding/json"
	"laundro-api-ca/app/middleware"
	"laundro-api-ca/business"
	_order "laundro-api-ca/business/orders"
	_orderMock "laundro-api-ca/business/orders/mocks"
	controller "laundro-api-ca/controllers"
	"laundro-api-ca/controllers/orders"
	"laundro-api-ca/controllers/orders/response"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var(
	mockOrderService _orderMock.Service
	orderCtrl *orders.OrderController
	orderReq string
	orderReqInvalidBind string
	orderReqInvalidStruct string
	orderDomain _order.Domain
	orderRes  response.Orders
	claims *middleware.JwtCustomClaims
)

func TestMain(m *testing.M){
	orderCtrl = orders.NewOrderController(&mockOrderService)
	orderReq = `{
		"laundromat_id": 1,
		"payment_id": 1,
		"product_name": "Test Product",
		"weight": 4
	}`
	orderReqInvalidBind = `{
		"laundromat_id": 1,
		"payment_id": 1,
		"product_name": "Test Product"
		"weight": 4
	}`
	orderReqInvalidStruct = `{
		"laundromat_id": 1,
		"payment_id": 1,
		"weight": 4
	}`
	orderDomain = _order.Domain{
		ID                  : 1,
		CreatedAt           : time.Now(),
		UserID              : 1,
		LaundromatID        : 1,
		LaundromatName		: "Test Laundry",
		PaymentID           : 1,
		PaymentGateway		: "Test Payment",
		ProductID           : 1,
		ProductName			: "Test Product",
		EstimatedFinishTime : time.Now(),
		Weight              : 4,
		TotalPrice          : 20000,
	}
	orderRes = response.Orders{
		ID                  : 1,
		CreatedAt           : time.Now(),
		UserID              : 1,
		LaundromatID        : 1,
		LaundromatName		: "Test Laundry",
		PaymentID           : 1,
		PaymentGateway		: "Test Payment",
		ProductID           : 1,
		ProductName			: "Test Product",
		EstimatedFinishTime : time.Now(),
		Weight              : 4,
		TotalPrice          : 20000,
	}
	claims = &middleware.JwtCustomClaims{
		1,
		jwt.StandardClaims{},
	}
	m.Run()
}

func TestCreate(t *testing.T){
	t.Run("Valid Test", func(t *testing.T){
		req := httptest.NewRequest(http.MethodPost, "/orders", strings.NewReader(orderReq))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		c := e.NewContext(req,rec)
		c.Set("user", &jwt.Token{Claims: claims})

		mockOrderService.On("Create", mock.AnythingOfType("uint"), mock.Anything).Return(orderDomain, nil).Once()
	
		resp := controller.BaseResponse{}
		resp.Meta.Status = http.StatusOK
		resp.Meta.Message = "Success"
		resp.Data = orderDomain
		expected, _ := json.Marshal(resp)

		if assert.NoError(t, orderCtrl.Create(c)){
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.JSONEq(t, string(expected), rec.Body.String())
		}
	})
	t.Run("Invalid Test | Bind Error", func(t *testing.T){
		req := httptest.NewRequest(http.MethodPost, "/orders", strings.NewReader(orderReqInvalidBind))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		c := e.NewContext(req,rec)
	
		resp := controller.BaseResponse{}
		resp.Meta.Status = http.StatusBadRequest
		resp.Meta.Message = "Error"
		resp.Meta.Messages = []string{"code=400, message=Syntax error: offset=79, error=invalid character '\"' after object key:value pair, internal=invalid character '\"' after object key:value pair"}
		expected, _ := json.Marshal(resp)

		if assert.NoError(t, orderCtrl.Create(c)){
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.JSONEq(t, string(expected), rec.Body.String())
		}
	})
	t.Run("Invalid Test | Invalid Struct", func(t *testing.T){
		req := httptest.NewRequest(http.MethodPost, "/orders", strings.NewReader(orderReqInvalidStruct))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		c := e.NewContext(req,rec)
	
		resp := controller.BaseResponse{}
		resp.Meta.Status = http.StatusBadRequest
		resp.Meta.Message = "Error"
		resp.Meta.Messages = []string{"product_name: non zero value required"}
		expected, _ := json.Marshal(resp)

		if assert.NoError(t, orderCtrl.Create(c)){
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.JSONEq(t, string(expected), rec.Body.String())
		}
	})
	t.Run("Invalid Test | Internal Server Error", func(t *testing.T){
		req := httptest.NewRequest(http.MethodPost, "/orders", strings.NewReader(orderReq))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		c := e.NewContext(req,rec)
		c.Set("user", &jwt.Token{Claims: claims})

		mockOrderService.On("Create", mock.AnythingOfType("uint"), mock.Anything).Return(_order.Domain{}, assert.AnError).Once()

		resp := controller.BaseResponse{}
		resp.Meta.Status = http.StatusInternalServerError
		resp.Meta.Message = "Error"
		resp.Meta.Messages = []string{assert.AnError.Error()}
		expected, _ := json.Marshal(resp)

		if assert.NoError(t, orderCtrl.Create(c)){
			assert.Equal(t, http.StatusInternalServerError, rec.Code)
			assert.JSONEq(t, string(expected), rec.Body.String())
		}
	})
}

func TestGetByUserID(t *testing.T){
	t.Run("Valid Test", func(t *testing.T){
		req := httptest.NewRequest(http.MethodGet, "/orders", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		c := e.NewContext(req,rec)
		c.Set("user", &jwt.Token{Claims: claims})

		mockOrderService.On("GetByUserID", mock.AnythingOfType("uint")).Return([]_order.Domain{orderDomain}, nil).Once()
	
		resp := controller.BaseResponse{}
		resp.Meta.Status = http.StatusOK
		resp.Meta.Message = "Success"
		resp.Data = []response.Orders{orderRes}
		expected, _ := json.Marshal(resp)

		if assert.NoError(t, orderCtrl.GetByUserID(c)){
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.JSONEq(t, string(expected), rec.Body.String())
		}
	})
	t.Run("Invalid Test | No Order Found", func(t *testing.T){
		req := httptest.NewRequest(http.MethodGet, "/orders", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		c := e.NewContext(req,rec)
		c.Set("user", &jwt.Token{Claims: claims})

		mockOrderService.On("GetByUserID", mock.AnythingOfType("uint")).Return([]_order.Domain{}, business.ErrOrdersNotFound).Once()
	
		resp := controller.BaseResponse{}
		resp.Meta.Status = http.StatusNotFound
		resp.Meta.Message = "Error"
		resp.Meta.Messages = []string{business.ErrOrdersNotFound.Error()}
		expected, _ := json.Marshal(resp)

		if assert.NoError(t, orderCtrl.GetByUserID(c)){
			assert.Equal(t, http.StatusNotFound, rec.Code)
			assert.JSONEq(t, string(expected), rec.Body.String())
		}
	})
}