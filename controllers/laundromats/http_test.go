package laundromats_test

import (
	"encoding/json"
	"laundro-api-ca/app/middleware"
	"laundro-api-ca/business"
	_laundro "laundro-api-ca/business/laundromats"
	_laundroMock "laundro-api-ca/business/laundromats/mocks"
	controller "laundro-api-ca/controllers"
	"laundro-api-ca/controllers/laundromats"
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
	mockLaundroService _laundroMock.Service
	laundroCtrl *laundromats.LaundromatController
	laundroReq string
	laundroReqInvalidBind string
	laundroReqInvalidStruct string
	laundroDomain _laundro.Domain
	claims *middleware.JwtCustomClaims
)
func TestMain(m *testing.M){
	laundroCtrl = laundromats.NewLaundromatController(&mockLaundroService)
	laundroReq = `{
		"name": "Test Laundro",
		"status": true,
		"address":{
			"street": "Test Street",
			"postal_code": 12345,
			"city": "Test City",
			"province": "Test Province"
		}
	}`
	laundroReqInvalidBind = `{
		"name": "Test Laundro"
		"status": true
		"address":{
			"street": "Test Street",
			"postal_code": 12345,
			"city": "Test City",
			"province": "Test Province"
		}
	}`
	laundroReqInvalidStruct = `{
		"name": "Test Laundro",
		"status": true
	}`
	laundroDomain = _laundro.Domain{
		Id        : 1,
		Name      : "Test Laundro",
		OwnerID   : 1,
		AddressID : 1,
		Status    : true,
		CreatedAt : time.Now(),
		UpdatedAt : time.Now(),
	}
	claims = &middleware.JwtCustomClaims{
		1,
		jwt.StandardClaims{},
	}
	m.Run()
}

func TestInsert(t *testing.T){
	t.Run("Valid Test", func(t *testing.T){
		req := httptest.NewRequest(http.MethodPost, "/laundro", strings.NewReader(laundroReq))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		c := e.NewContext(req,rec)
		c.Set("user", &jwt.Token{Claims: claims})

		mockLaundroService.On("Insert", mock.AnythingOfType("uint"), mock.Anything, mock.Anything).Return(laundroDomain, nil).Once()
		
		resp := controller.BaseResponse{}
		resp.Meta.Status = http.StatusOK
		resp.Meta.Message = "Success"
		resp.Data = laundroDomain
		expected, _ := json.Marshal(resp)

		if assert.NoError(t, laundroCtrl.Insert(c)){
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.JSONEq(t, string(expected), rec.Body.String())
		}
	})
	t.Run("Invalid Test | Bind Error", func(t *testing.T){
		req := httptest.NewRequest(http.MethodPost, "/laundro", strings.NewReader(laundroReqInvalidBind))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		c := e.NewContext(req,rec)
		
		resp := controller.BaseResponse{}
		resp.Meta.Status = http.StatusBadRequest
		resp.Meta.Message = "Error"
		resp.Meta.Messages = []string{"code=400, message=Syntax error: offset=30, error=invalid character '\"' after object key:value pair, internal=invalid character '\"' after object key:value pair"}
		expected, _ := json.Marshal(resp)

		if assert.NoError(t, laundroCtrl.Insert(c)){
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.JSONEq(t, string(expected), rec.Body.String())
		}
	})
	t.Run("Invalid Test | Invalid Struct", func(t *testing.T){
		req := httptest.NewRequest(http.MethodPost, "/laundro", strings.NewReader(laundroReqInvalidStruct))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		c := e.NewContext(req,rec)
		
		resp := controller.BaseResponse{}
		resp.Meta.Status = http.StatusBadRequest
		resp.Meta.Message = "Error"
		resp.Meta.Messages = []string{"address: non zero value required"}
		expected, _ := json.Marshal(resp)

		if assert.NoError(t, laundroCtrl.Insert(c)){
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.JSONEq(t, string(expected), rec.Body.String())
		}
	})
	t.Run("Invalid Test | Internal Server Error", func(t *testing.T){
		req := httptest.NewRequest(http.MethodPost, "/laundro", strings.NewReader(laundroReq))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		c := e.NewContext(req,rec)
		c.Set("user", &jwt.Token{Claims: claims})
		
		mockLaundroService.On("Insert", mock.AnythingOfType("uint"), mock.Anything, mock.Anything).Return(_laundro.Domain{}, business.ErrDuplicateData).Once()

		resp := controller.BaseResponse{}
		resp.Meta.Status = http.StatusInternalServerError
		resp.Meta.Message = "Error"
		resp.Meta.Messages = []string{business.ErrDuplicateData.Error()}
		expected, _ := json.Marshal(resp)

		if assert.NoError(t, laundroCtrl.Insert(c)){
			assert.Equal(t, http.StatusInternalServerError, rec.Code)
			assert.JSONEq(t, string(expected), rec.Body.String())
		}
	})
}

func TestGetByIP(t *testing.T){
	t.Run("Valid Test", func(t *testing.T){
		req := httptest.NewRequest(http.MethodGet, "/laundro/find-ip", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		c := e.NewContext(req,rec)

		mockLaundroService.On("GetByIP", mock.AnythingOfType("string")).Return([]_laundro.Domain{laundroDomain}, nil).Once()

		resp := controller.BaseResponse{}
		resp.Meta.Status = http.StatusOK
		resp.Meta.Message = "Success"
		resp.Data = []_laundro.Domain{laundroDomain}
		expected, _ := json.Marshal(resp)

		if assert.NoError(t, laundroCtrl.GetByIP(c)){
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.JSONEq(t, string(expected), rec.Body.String())
		}
	})
	t.Run("Invalid Test | Laundromat Not Found", func(t *testing.T){
		req := httptest.NewRequest(http.MethodGet, "/laundro/find-ip", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		c := e.NewContext(req,rec)

		mockLaundroService.On("GetByIP", mock.AnythingOfType("string")).Return([]_laundro.Domain{}, business.ErrNearestLaundromatNotFound).Once()

		resp := controller.BaseResponse{}
		resp.Meta.Status = http.StatusNotFound
		resp.Meta.Message = "Error"
		resp.Meta.Messages = []string{business.ErrNearestLaundromatNotFound.Error()}
		expected, _ := json.Marshal(resp)

		if assert.NoError(t, laundroCtrl.GetByIP(c)){
			assert.Equal(t, http.StatusNotFound, rec.Code)
			assert.JSONEq(t, string(expected), rec.Body.String())
		}
	})
}

func TestGetByName(t *testing.T){
	t.Run("Valid Test", func(t *testing.T){
		req := httptest.NewRequest(http.MethodGet, "/laundro/find-name/Test", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		c := e.NewContext(req,rec)

		mockLaundroService.On("GetByName", mock.AnythingOfType("string")).Return([]_laundro.Domain{laundroDomain}, nil).Once()
	
		resp := controller.BaseResponse{}
		resp.Meta.Status = http.StatusOK
		resp.Meta.Message = "Success"
		resp.Data = []_laundro.Domain{laundroDomain}
		expected, _ := json.Marshal(resp)

		if assert.NoError(t, laundroCtrl.GetByName(c)){
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.JSONEq(t, string(expected), rec.Body.String())
		}
	})
	t.Run("Invalid Test | Laundromat Not Found", func(t *testing.T){
		req := httptest.NewRequest(http.MethodGet, "/laundro/find-name/Test", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		c := e.NewContext(req,rec)

		mockLaundroService.On("GetByName", mock.AnythingOfType("string")).Return([]_laundro.Domain{}, business.ErrLaundromatNotFound).Once()
	
		resp := controller.BaseResponse{}
		resp.Meta.Status = http.StatusNotFound
		resp.Meta.Message = "Error"
		resp.Meta.Messages = []string{business.ErrLaundromatNotFound.Error()}
		expected, _ := json.Marshal(resp)

		if assert.NoError(t, laundroCtrl.GetByName(c)){
			assert.Equal(t, http.StatusNotFound, rec.Code)
			assert.JSONEq(t, string(expected), rec.Body.String())
		}
	})
}

func TestGetByID(t *testing.T){
	t.Run("Valid Test", func(t *testing.T){
		mockLaundroService.On("GetByID", mock.AnythingOfType("uint")).Return(laundroDomain, nil).Once()

		resp := laundroCtrl.GetByID(1)

		assert.Equal(t, laundroDomain, resp)
	})
	t.Run("Invalid Test | Laundromat Not Found", func(t *testing.T){
		mockLaundroService.On("GetByID", mock.AnythingOfType("uint")).Return(_laundro.Domain{}, business.ErrLaundromatNotFound).Once()

		resp := laundroCtrl.GetByID(1)

		assert.Empty(t, resp)
		assert.NotEqual(t, laundroDomain, resp)
	})
}

func TestUpdate(t *testing.T){
	t.Run("Valid Test", func(t *testing.T){
		req := httptest.NewRequest(http.MethodPut, "/laundro/edit/1", strings.NewReader(laundroReq))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		c := e.NewContext(req,rec)

		mockLaundroService.On("Update", mock.AnythingOfType("uint"), mock.Anything, mock.Anything).Return(laundroDomain, nil).Once()
		
		resp := controller.BaseResponse{}
		resp.Meta.Status = http.StatusOK
		resp.Meta.Message = "Success"
		resp.Data = laundroDomain
		expected, _ := json.Marshal(resp)

		if assert.NoError(t, laundroCtrl.Update(c)){
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.JSONEq(t, string(expected), rec.Body.String())
		}
	})
	t.Run("Invalid Test | Bind Error", func(t *testing.T){
		req := httptest.NewRequest(http.MethodPut, "/laundro/edit/1", strings.NewReader(laundroReqInvalidBind))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		c := e.NewContext(req,rec)
		
		resp := controller.BaseResponse{}
		resp.Meta.Status = http.StatusBadRequest
		resp.Meta.Message = "Error"
		resp.Meta.Messages = []string{"code=400, message=Syntax error: offset=30, error=invalid character '\"' after object key:value pair, internal=invalid character '\"' after object key:value pair"}
		expected, _ := json.Marshal(resp)

		if assert.NoError(t, laundroCtrl.Update(c)){
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.JSONEq(t, string(expected), rec.Body.String())
		}
	})
	t.Run("Invalid Test | Invalid Struct", func(t *testing.T){
		req := httptest.NewRequest(http.MethodPut, "/laundro/edit/1", strings.NewReader(laundroReqInvalidStruct))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		c := e.NewContext(req,rec)
		
		resp := controller.BaseResponse{}
		resp.Meta.Status = http.StatusBadRequest
		resp.Meta.Message = "Error"
		resp.Meta.Messages = []string{"address: non zero value required"}
		expected, _ := json.Marshal(resp)

		if assert.NoError(t, laundroCtrl.Update(c)){
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.JSONEq(t, string(expected), rec.Body.String())
		}
	})
	t.Run("Invalid Test | Internal Server Error", func(t *testing.T){
		req := httptest.NewRequest(http.MethodPut, "/laundro/edit/1", strings.NewReader(laundroReq))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		c := e.NewContext(req,rec)
		
		mockLaundroService.On("Update", mock.AnythingOfType("uint"), mock.Anything, mock.Anything).Return(_laundro.Domain{}, business.ErrLaundromatNotFound).Once()

		resp := controller.BaseResponse{}
		resp.Meta.Status = http.StatusInternalServerError
		resp.Meta.Message = "Error"
		resp.Meta.Messages = []string{business.ErrLaundromatNotFound.Error()}
		expected, _ := json.Marshal(resp)

		if assert.NoError(t, laundroCtrl.Update(c)){
			assert.Equal(t, http.StatusInternalServerError, rec.Code)
			assert.JSONEq(t, string(expected), rec.Body.String())
		}
	})
}

func TestDelete(t *testing.T){
	t.Run("Valid Test", func(t *testing.T){
		req := httptest.NewRequest(http.MethodDelete, "/laundro/1", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		c := e.NewContext(req,rec)

		mockLaundroService.On("Delete", mock.AnythingOfType("uint")).Return("Laundromat Deleted", nil).Once()
	
		resp := controller.BaseResponse{}
		resp.Meta.Status = http.StatusOK
		resp.Meta.Message = "Success"
		resp.Data = "Laundromat Deleted"
		expected, _ := json.Marshal(resp)

		if assert.NoError(t, laundroCtrl.Delete(c)){
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.JSONEq(t, string(expected), rec.Body.String())
		}
	})
	t.Run("Valid Test", func(t *testing.T){
		req := httptest.NewRequest(http.MethodDelete, "/laundro/1", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		c := e.NewContext(req,rec)

		mockLaundroService.On("Delete", mock.AnythingOfType("uint")).Return("", business.ErrLaundromatNotFound).Once()
	
		resp := controller.BaseResponse{}
		resp.Meta.Status = http.StatusInternalServerError
		resp.Meta.Message = "Error"
		resp.Meta.Messages = []string{business.ErrLaundromatNotFound.Error()}
		expected, _ := json.Marshal(resp)

		if assert.NoError(t, laundroCtrl.Delete(c)){
			assert.Equal(t, http.StatusInternalServerError, rec.Code)
			assert.JSONEq(t, string(expected), rec.Body.String())
		}
	})
}