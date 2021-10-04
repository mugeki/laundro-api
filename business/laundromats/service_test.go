package laundromats_test

import (
	"laundro-api-ca/business/addresses"
	_addressMock "laundro-api-ca/business/addresses/mocks"
	"laundro-api-ca/business/geolocation"
	_geoMock "laundro-api-ca/business/geolocation/mocks"
	"laundro-api-ca/business/laundromats"
	_laundroMock "laundro-api-ca/business/laundromats/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var(
	mockLaundroRepository _laundroMock.Repository
	mockAddressRepository _addressMock.Repository
	mockGeoRepository _geoMock.Repository
	laundroService	laundromats.Service	
	laundroDomain	laundromats.Domain
	laundroNewDomain laundromats.Domain
	laundroNewAddrDomain laundromats.Domain
	addressDomain  	addresses.Domain
	addressNewDomain  	addresses.Domain
	geoDomain  		geolocation.Domain
)

func TestMain(m *testing.M){
	laundroService = laundromats.NewLaundromatService(&mockLaundroRepository,&mockAddressRepository,&mockGeoRepository)
	laundroDomain = laundromats.Domain{
		Id        : 1,
		Name      : "Test Laundro",
		OwnerID   : 1,
		AddressID : 1,
		Status    : true,
		CreatedAt : time.Now(),
		UpdatedAt : time.Now(),
	}
	laundroNewDomain = laundromats.Domain{
		Id        : 1,
		Name      : "Test Laundro New",
		OwnerID   : 1,
		AddressID : 1,
		Status    : false,
		CreatedAt : time.Now(),
		UpdatedAt : time.Now(),
	}
	laundroNewAddrDomain = laundromats.Domain{
		Id        : 1,
		Name      : "Test Laundro New",
		OwnerID   : 1,
		AddressID : 2,
		Status    : false,
		CreatedAt : time.Now(),
		UpdatedAt : time.Now(),
	}
	addressDomain = addresses.Domain{
		ID         : 1,
		Street     : "Test Street",
		PostalCode : 12345,
		City       : "Test City",
		Province   : "Test Province",
	}
	addressNewDomain = addresses.Domain{
		ID         : 2,
		Street     : "Test Street 2",
		PostalCode : 99999,
		City       : "Test City 2",
		Province   : "Test Province 2",
	}
	geoDomain = geolocation.Domain{
		IP	: "0.0.0.0",
		City: "Test City",
	}
	m.Run()
}

func TestInsert(t *testing.T){
	t.Run("Valid Test", func(t *testing.T){
		mockAddressRepository.On("Insert", mock.Anything).Return(addressDomain, nil).Once()
		mockLaundroRepository.On("Insert", mock.Anything).Return(laundroDomain, nil).Once()

		input := laundromats.Domain{
			Name      : "Test Laundro",
			Status    : true,
		}

		resp, err := laundroService.Insert(1, &input, &addressDomain)

		assert.Nil(t, err)
		assert.Equal(t, resp, laundroDomain)
	})
	t.Run("Invalid Test", func(t *testing.T){
		mockAddressRepository.On("Insert", mock.Anything).Return(addressNewDomain, nil).Once()
		mockLaundroRepository.On("Insert", mock.Anything).Return(laundromats.Domain{}, assert.AnError).Once()

		input := laundromats.Domain{
			Name      : "Test Laundro",
			Status    : true,
		}

		resp, err := laundroService.Insert(1, &input, &addressNewDomain)

		assert.NotNil(t, err)
		assert.Equal(t, laundromats.Domain{}, resp)
	})
}

func TestGetByIP(t *testing.T){
	t.Run("Valid Test", func(t *testing.T){
		mockGeoRepository.On("GetLocationByIP").Return(geoDomain, nil).Once()
		mockAddressRepository.On("FindByCity", mock.AnythingOfType("string")).Return([]addresses.Domain{addressDomain}, nil).Once()
		mockLaundroRepository.On("GetByAddress", mock.AnythingOfType("[]uint")).Return([]laundromats.Domain{laundroDomain}, nil).Once()

		resp, err := laundroService.GetByIP()

		assert.Nil(t, err)
		assert.Contains(t, resp, laundroDomain)
	})
	t.Run("Invalid Test | Location Not Found", func(t *testing.T){
		mockGeoRepository.On("GetLocationByIP").Return(geolocation.Domain{}, assert.AnError).Once()

		resp, err := laundroService.GetByIP()

		assert.NotNil(t, err)
		assert.NotContains(t, resp, laundroDomain)
	})
	t.Run("Invalid Test | Addresses Not Found", func(t *testing.T){
		mockGeoRepository.On("GetLocationByIP").Return(geoDomain, nil).Once()
		mockAddressRepository.On("FindByCity", mock.AnythingOfType("string")).Return([]addresses.Domain{}, assert.AnError).Once()

		resp, err := laundroService.GetByIP()

		assert.NotNil(t, err)
		assert.NotContains(t, resp, laundroDomain)
	})
	t.Run("Invalid Test | Laundromat Not Found", func(t *testing.T){
		mockGeoRepository.On("GetLocationByIP").Return(geoDomain, nil).Once()
		mockAddressRepository.On("FindByCity", mock.AnythingOfType("string")).Return([]addresses.Domain{addressDomain}, nil).Once()
		mockLaundroRepository.On("GetByAddress", mock.AnythingOfType("[]uint")).Return([]laundromats.Domain{}, assert.AnError).Once()

		resp, err := laundroService.GetByIP()

		assert.NotNil(t, err)
		assert.NotContains(t, resp, laundroDomain)
	})
}

func TestGetByName(t *testing.T){
	t.Run("Valid Test", func(t *testing.T){
		mockLaundroRepository.On("GetByName", mock.AnythingOfType("string")).Return([]laundromats.Domain{laundroDomain}, nil).Once()

		resp, err := laundroService.GetByName("Laundry")

		assert.Nil(t, err)
		assert.Contains(t, resp, laundroDomain)
	})
	t.Run("Invalid Test", func(t *testing.T){
		mockLaundroRepository.On("GetByName", mock.AnythingOfType("string")).Return([]laundromats.Domain{}, assert.AnError).Once()

		resp, err := laundroService.GetByName("Laundry")

		assert.NotNil(t, err)
		assert.NotContains(t, resp, laundroDomain)
	})
}
func TestGetByID(t *testing.T){
	t.Run("Valid Test", func(t *testing.T){
		mockLaundroRepository.On("GetByID", mock.AnythingOfType("uint")).Return(laundroDomain, nil).Once()

		resp, err := laundroService.GetByID(1)

		assert.Nil(t, err)
		assert.Equal(t, laundroDomain, resp)
	})
	t.Run("Invalid Test", func(t *testing.T){
		mockLaundroRepository.On("GetByID", mock.AnythingOfType("uint")).Return(laundromats.Domain{}, assert.AnError).Once()

		resp, err := laundroService.GetByID(1)

		assert.NotNil(t, err)
		assert.Equal(t, laundromats.Domain{}, resp)
	})
}

func TestUpdate(t *testing.T){
	t.Run("Valid Test | Address Unchanged", func(t *testing.T){
		mockAddressRepository.On("Insert", mock.Anything).Return(addressDomain, nil).Once()
		mockLaundroRepository.On("Update", mock.AnythingOfType("uint"), mock.Anything).Return(laundroNewDomain, nil).Once()

		input := laundromats.Domain{
			Name      : "Test Laundro New",
			Status    : false,
		}

		resp, err := laundroService.Update(1, &input, &addressDomain)

		assert.Nil(t, err)
		assert.Equal(t, laundroNewDomain, resp)
	})
	t.Run("Valid Test | New Address", func(t *testing.T){
		mockAddressRepository.On("Insert", mock.Anything).Return(addressNewDomain, nil).Once()
		mockLaundroRepository.On("Update", mock.AnythingOfType("uint"), mock.Anything).Return(laundroNewAddrDomain, nil).Once()

		input := laundromats.Domain{
			Name      : "Test Laundro New",
			Status    : false,
		}

		resp, err := laundroService.Update(1, &input, &addressNewDomain)

		assert.Nil(t, err)
		assert.Equal(t, laundroNewAddrDomain, resp)
	})
	t.Run("Invalid Test", func(t *testing.T){
		mockAddressRepository.On("Insert", mock.Anything).Return(addressDomain, nil).Once()
		mockLaundroRepository.On("Update", mock.AnythingOfType("uint"), mock.Anything).Return(laundroDomain, assert.AnError).Once()

		resp, err := laundroService.Update(2, &laundroDomain, &addressDomain)

		assert.NotNil(t, err)
		assert.Equal(t, laundromats.Domain{}, resp)
	})
}
func TestDelete(t *testing.T){
	t.Run("Valid Test", func(t *testing.T){
		mockLaundroRepository.On("Delete", mock.Anything).Return("Laundromat Deleted", nil).Once()

		resp, err := laundroService.Delete(1)

		assert.Nil(t, err)
		assert.Equal(t, "Laundromat Deleted", resp)
	})
	t.Run("Invalid Test", func(t *testing.T){
		mockLaundroRepository.On("Delete", mock.Anything).Return("", assert.AnError).Once()

		resp, err := laundroService.Delete(1)

		assert.NotNil(t, err)
		assert.Equal(t, "", resp)
	})
}