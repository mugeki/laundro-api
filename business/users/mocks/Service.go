// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	addresses "laundro-api-ca/business/addresses"

	mock "github.com/stretchr/testify/mock"

	users "laundro-api-ca/business/users"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// GetByID provides a mock function with given fields: id
func (_m *Service) GetByID(id uint) (users.Domain, error) {
	ret := _m.Called(id)

	var r0 users.Domain
	if rf, ok := ret.Get(0).(func(uint) users.Domain); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(users.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Login provides a mock function with given fields: username, password
func (_m *Service) Login(username string, password string) (string, error) {
	ret := _m.Called(username, password)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(username, password)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(username, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Register provides a mock function with given fields: userData, addressData
func (_m *Service) Register(userData *users.Domain, addressData *addresses.Domain) (users.Domain, error) {
	ret := _m.Called(userData, addressData)

	var r0 users.Domain
	if rf, ok := ret.Get(0).(func(*users.Domain, *addresses.Domain) users.Domain); ok {
		r0 = rf(userData, addressData)
	} else {
		r0 = ret.Get(0).(users.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*users.Domain, *addresses.Domain) error); ok {
		r1 = rf(userData, addressData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
