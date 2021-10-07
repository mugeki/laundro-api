package users_test

import (
	"laundro-api-ca/app/middleware"
	"laundro-api-ca/business/addresses"
	_addressMock "laundro-api-ca/business/addresses/mocks"
	"laundro-api-ca/business/users"
	_userMock "laundro-api-ca/business/users/mocks"
	"laundro-api-ca/helper/encrypt"
	_encryptMock "laundro-api-ca/helper/encrypt/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var(
	mockUserRepository _userMock.Repository
	mockAddrRepository _addressMock.Repository
	mockEncrypt _encryptMock.Helper
	userService	users.Service
	hashedPassword string
	userDomain	users.Domain
	addressDomain  addresses.Domain
)

func TestMain(m *testing.M){
	userService = users.NewUserService(&mockUserRepository, &mockAddrRepository, &middleware.ConfigJWT{})
	hashedPassword, _ = encrypt.Hash("test")
	userDomain = users.Domain{
		Username	: "testUser",
		Password	: hashedPassword,
		Email		: "test@gmail.com",
		Fullname	: "Test John",
		DateOfBirth	: time.Now(),
		PhoneNumber	: "123456789",
		RoleID		: 1,
		AddressID	: 1,
	}
	addressDomain = addresses.Domain{
		ID         : 1,
		Street     : "Test Street",
		PostalCode : 12345,
		City       : "Test City",
		Province   : "Test Province",
	}
	m.Run()
}

func TestRegister(t *testing.T){
	t.Run("Valid Test", func(t *testing.T){
		mockAddrRepository.On("Insert", mock.Anything).Return(addressDomain, nil).Once()
		mockUserRepository.On("Register", mock.Anything).Return(userDomain, nil).Once()

		inputUser := users.Domain{
			Username    : "testUser",
			Password    : "test",
			Email       : "test@gmail.com",
			Fullname    : "Test John",
			DateOfBirth : time.Now(),
			PhoneNumber : "123456789",
			RoleID      : 1,
		}

		inputAddr := addresses.Domain{
			Street     : "Test Street",
			PostalCode : 12345,
			City       : "Test City",
			Province   : "Test Province",
		}
		
		resp, err := userService.Register(&inputUser, &inputAddr)

		assert.Nil(t, err)
		assert.Equal(t, userDomain, resp)
	})
	t.Run("Invalid Test | Duplicate User", func(t *testing.T){
		mockAddrRepository.On("Insert", mock.Anything).Return(addressDomain, nil).Once()
		mockUserRepository.On("Register", mock.Anything).Return(users.Domain{}, assert.AnError).Once()

		inputUser := users.Domain{
			Username    : "testUser",
			Password    : "test",
			Email       : "test@gmail.com",
			Fullname    : "Test John",
			DateOfBirth : time.Now(),
			PhoneNumber : "123456789",
			RoleID      : 1,
		}

		inputAddr := addresses.Domain{
			Street     : "Test Street",
			PostalCode : 12345,
			City       : "Test City",
			Province   : "Test Province",
		}
		
		resp, err := userService.Register(&inputUser, &inputAddr)

		assert.NotNil(t, err)
		assert.NotEqual(t, userDomain, resp)
	})
	t.Run("Invalid Test | Internal Error", func(t *testing.T){
		mockAddrRepository.On("Insert", mock.Anything).Return(addressDomain, nil).Once()
		mockUserRepository.On("Register", mock.Anything).Return(users.Domain{}, assert.AnError).Once()

		inputUser := users.Domain{
			Username    : "testUser",
			Password    : "test",
			Email       : "test@gmail.com",
			Fullname    : "Test John",
			DateOfBirth : time.Now(),
			PhoneNumber : "123456789",
			RoleID      : 1,
		}

		inputAddr := addresses.Domain{
			Street     : "Test Street",
			PostalCode : 12345,
			City       : "Test City",
			Province   : "Test Province",
		}
		
		resp, err := userService.Register(&inputUser, &inputAddr)

		assert.NotNil(t, err)
		assert.NotEqual(t, userDomain, resp)
	})
}

func TestLogin(t *testing.T){
	t.Run("Valid Test", func(t *testing.T){
		mockUserRepository.On("GetByUsername", mock.AnythingOfType("string")).Return(userDomain, nil).Once()
		
		input := users.Domain{
			Username    : "testUser",
			Password    : "test",
		}

		resp, err := userService.Login(input.Username, input.Password)

		assert.Nil(t, err)
		assert.NotEmpty(t, resp)
	})
	t.Run("Invalid Test | Wrong Username", func(t *testing.T){
		mockUserRepository.On("GetByUsername", mock.AnythingOfType("string")).Return(users.Domain{}, assert.AnError).Once()
		
		input := users.Domain{
			Username    : "testUser",
			Password    : "test",
		}

		resp, err := userService.Login(input.Username, input.Password)

		assert.NotNil(t, err)
		assert.Empty(t, resp)
	})
	t.Run("Invalid Test | Wrong Password", func(t *testing.T){
		mockUserRepository.On("GetByUsername", mock.AnythingOfType("string")).Return(userDomain, nil).Once()
		mockEncrypt.On("ValidateHash", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return("", assert.AnError)

		input := users.Domain{
			Username    : "testUser",
			Password    : "wrong",
		}

		resp, err := userService.Login(input.Username, input.Password)

		assert.NotNil(t, err)
		assert.Empty(t, resp)
	})
}

func TestGetByID(t *testing.T){
	t.Run("Valid Test", func(t *testing.T){
		mockUserRepository.On("GetByID", mock.AnythingOfType("uint")).Return(userDomain, nil).Once()

		resp, err := userService.GetByID(1)

		assert.Nil(t, err)
		assert.Equal(t, userDomain, resp)
	})
	t.Run("Invalid Test | User Not Found", func(t *testing.T){
		mockUserRepository.On("GetByID", mock.AnythingOfType("uint")).Return(users.Domain{}, assert.AnError).Once()

		resp, err := userService.GetByID(2)

		assert.NotNil(t, err)
		assert.NotEqual(t, userDomain, resp)
	})
}