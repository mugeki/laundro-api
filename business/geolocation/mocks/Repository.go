// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	geolocation "laundro-api-ca/business/geolocation"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// GetLocationByIP provides a mock function with given fields: ip
func (_m *Repository) GetLocationByIP(ip string) (geolocation.Domain, error) {
	ret := _m.Called(ip)

	var r0 geolocation.Domain
	if rf, ok := ret.Get(0).(func(string) geolocation.Domain); ok {
		r0 = rf(ip)
	} else {
		r0 = ret.Get(0).(geolocation.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(ip)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
