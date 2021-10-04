package users_test

import (
	"encoding/json"
	"laundro-api-ca/business"
	_user "laundro-api-ca/business/users"
	_userMock "laundro-api-ca/business/users/mocks"
	controller "laundro-api-ca/controllers"
	"laundro-api-ca/controllers/users"
	"laundro-api-ca/helper/encrypt"
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
	mockUserService _userMock.Service
	userCtrl	*users.UserController
	userReq	string
	userReqInvalidBind string
	userReqInvalidStruct string
	userLoginReq string
	hashedPassword string
	userDomain _user.Domain
)

func TestMain(m *testing.M){
	userCtrl = users.NewUserController(&mockUserService)
	userReq = `{
		"username": "testUser",
		"password": "testPassword",
		"email": "test@gmail.com",
		"fullname": "Test John",
		"date_of_birth": "2001-01-01T00:00:00Z",
		"phone_number": "123456789",
		"role_id": 1,
		"address":{
			"street": "Test Street",
			"postal_code": 12345,
			"city": "Test City",
			"province": "Test Province"
		}
	}`
	userReqInvalidBind = `{
		"username": "testUser",
		"password": "test",,
		"email": "test",
		"fullname": "Test John",
		"date_of_birth": "2001-01-01T00:00:00Z",
		"phone_number": "123456789",
		"role_id": 1,
		"address":{
			"street": "Test Street",
			"postal_code": 12345,
			"city": "Test City",
			"province": "Test Province"
		}
	}`
	userReqInvalidStruct = `{
		"username": "testUser",
		"password": "test",
		"email": "test",
		"fullname": "Test John",
		"date_of_birth": "2001-01-01T00:00:00Z",
		"phone_number": "123456789",
		"role_id": 1,
		"address":{
			"street": "Test Street",
			"postal_code": 12345,
			"city": "Test City",
			"province": "Test Province"
		}
	}`
	userLoginReq = `{"username": "testUser","password": "testPassword"}`
	hashedPassword, _ = encrypt.Hash("testPassword")
	userDomain = _user.Domain{
		Username	: "testUser",
		Password	: hashedPassword,
		Email		: "test@gmail.com",
		Fullname	: "Test John",
		DateOfBirth	: time.Date(2001,time.January,1,0,0,0,0,time.Local),
		PhoneNumber	: "123456789",
		RoleID		: 1,
		AddressID	: 1,
	}
	m.Run()
}

func TestRegister(t *testing.T){
	t.Run("Valid Test", func(t *testing.T){
		req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(userReq))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		c := e.NewContext(req,rec)
		mockUserService.On("Register", mock.Anything, mock.Anything).Return(userDomain, nil).Once()
		
		resp := controller.BaseResponse{}
		resp.Meta.Status = http.StatusOK
		resp.Meta.Message = "Success"
		resp.Data = userDomain
		expected, _ := json.Marshal(resp)

		if assert.NoError(t, userCtrl.Register(c)){
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.JSONEq(t, string(expected), rec.Body.String())
		}
	})
	t.Run("Invalid Test | Bind Error", func(t *testing.T){	
		req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(userReqInvalidBind))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		c := e.NewContext(req,rec)

		resp := controller.BaseResponse{}
		resp.Meta.Status = http.StatusBadRequest
		resp.Meta.Message = "Error"
		resp.Meta.Messages = []string{"code=400, message=Syntax error: offset=50, error=invalid character ',' looking for beginning of object key string, internal=invalid character ',' looking for beginning of object key string"}
		expected, _ := json.Marshal(resp)
		
		if assert.NoError(t, userCtrl.Register(c)){
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.JSONEq(t, string(expected), rec.Body.String())
		}
	})
	t.Run("Invalid Test | Invalid Struct ", func(t *testing.T){	
		req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(userReqInvalidStruct))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		c := e.NewContext(req,rec)

		resp := controller.BaseResponse{}
		resp.Meta.Status = http.StatusBadRequest
		resp.Meta.Message = "Error"
		resp.Meta.Messages = []string{"email: test does not validate as email;password: test does not validate as minstringlength(6)"}
		expected, _ := json.Marshal(resp)
		
		if assert.NoError(t, userCtrl.Register(c)){
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.JSONEq(t, string(expected), rec.Body.String())
		}
	})
	t.Run("Invalid Test | Internal Server Error", func(t *testing.T){
		req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(userReq))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		c := e.NewContext(req,rec)
		mockUserService.On("Register", mock.Anything, mock.Anything).Return(_user.Domain{}, business.ErrDuplicateData).Once()
		
		resp := controller.BaseResponse{}
		resp.Meta.Status = http.StatusInternalServerError
		resp.Meta.Message = "Error"
		resp.Meta.Messages = []string{business.ErrDuplicateData.Error()}
		expected, _ := json.Marshal(resp)
		
		if assert.NoError(t, userCtrl.Register(c)){
			assert.Equal(t, http.StatusInternalServerError, rec.Code)
			assert.JSONEq(t, string(expected), rec.Body.String())
		}
	})
}

func TestLogin(t *testing.T){
	t.Run("Valid Test", func(t *testing.T){
		req := httptest.NewRequest(http.MethodPost, "/users/login", strings.NewReader(userLoginReq))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		c := e.NewContext(req,rec)
		mockUserService.On("Login", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return("token", nil).Once()

		token := struct {
			Token string `json:"token"`
		}{Token: "token"}

		resp := controller.BaseResponse{}
		resp.Meta.Status = http.StatusOK
		resp.Meta.Message = "Success"
		resp.Data = token
		expected, _ := json.Marshal(resp)

		if assert.NoError(t, userCtrl.Login(c)){
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.JSONEq(t, string(expected), rec.Body.String())
		}
	})
	t.Run("Invalid Test | Bind Error", func(t *testing.T){
		userLoginReqInvalid := `{"username": "testUser",,"password": "testPassword"}`
		
		req := httptest.NewRequest(http.MethodPost, "/users/login", strings.NewReader(userLoginReqInvalid))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		c := e.NewContext(req,rec)

		resp := controller.BaseResponse{}
		resp.Meta.Status = http.StatusBadRequest
		resp.Meta.Message = "Error"
		resp.Meta.Messages = []string{"code=400, message=Syntax error: offset=25, error=invalid character ',' looking for beginning of object key string, internal=invalid character ',' looking for beginning of object key string"}
		expected, _ := json.Marshal(resp)
		
		if assert.NoError(t, userCtrl.Login(c)){
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.JSONEq(t, string(expected), rec.Body.String())
		}
	})
	t.Run("Invalid Test | Invalid Struct", func(t *testing.T){
		userLoginReqInvalid := `{"username": "test","password": "testPassword"}`
		
		req := httptest.NewRequest(http.MethodPost, "/users/login", strings.NewReader(userLoginReqInvalid))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		c := e.NewContext(req,rec)

		resp := controller.BaseResponse{}
		resp.Meta.Status = http.StatusBadRequest
		resp.Meta.Message = "Error"
		resp.Meta.Messages = []string{"username: test does not validate as minstringlength(6)"}
		expected, _ := json.Marshal(resp)
		
		if assert.NoError(t, userCtrl.Login(c)){
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.JSONEq(t, string(expected), rec.Body.String())
		}
	})
	t.Run("Invalid Test | Internal Server Error", func(t *testing.T){
		userLoginReqInvalid := `{"username": "testUser","password": "testNotUsser"}`
		
		req := httptest.NewRequest(http.MethodPost, "/users/login", strings.NewReader(userLoginReqInvalid))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		c := e.NewContext(req,rec)
		mockUserService.On("Login", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return("", business.ErrInvalidLoginInfo).Once()

		resp := controller.BaseResponse{}
		resp.Meta.Status = http.StatusInternalServerError
		resp.Meta.Message = "Error"
		resp.Meta.Messages = []string{business.ErrInvalidLoginInfo.Error()}
		expected, _ := json.Marshal(resp)
		
		if assert.NoError(t, userCtrl.Login(c)){
			assert.Equal(t, http.StatusInternalServerError, rec.Code)
			assert.JSONEq(t, string(expected), rec.Body.String())
		}
	})
}

func TestGetRoleByID(t *testing.T){
	t.Run("Valid Test", func(t *testing.T){
		mockUserService.On("GetByID", mock.AnythingOfType("uint")).Return(userDomain, nil).Once()

		resp := userCtrl.GetRoleByID(1)
		
		assert.Equal(t, userDomain.RoleID, uint(resp))
	})
	t.Run("Invalid Test", func(t *testing.T){
		mockUserService.On("GetByID", mock.AnythingOfType("uint")).Return(_user.Domain{}, business.ErrUserNotFound).Once()

		resp := userCtrl.GetRoleByID(1)
		
		assert.NotEqual(t, userDomain.RoleID, uint(resp))
	})
}