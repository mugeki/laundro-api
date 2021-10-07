// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	orders "laundro-api-ca/business/orders"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Create provides a mock function with given fields: orderData
func (_m *Repository) Create(orderData *orders.Domain) (orders.Domain, error) {
	ret := _m.Called(orderData)

	var r0 orders.Domain
	if rf, ok := ret.Get(0).(func(*orders.Domain) orders.Domain); ok {
		r0 = rf(orderData)
	} else {
		r0 = ret.Get(0).(orders.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*orders.Domain) error); ok {
		r1 = rf(orderData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByUserID provides a mock function with given fields: userId
func (_m *Repository) GetByUserID(userId uint) ([]orders.Domain, error) {
	ret := _m.Called(userId)

	var r0 []orders.Domain
	if rf, ok := ret.Get(0).(func(uint) []orders.Domain); ok {
		r0 = rf(userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]orders.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPaymentGateway provides a mock function with given fields: paymentId
func (_m *Repository) GetPaymentGateway(paymentId int) (string, error) {
	ret := _m.Called(paymentId)

	var r0 string
	if rf, ok := ret.Get(0).(func(int) string); ok {
		r0 = rf(paymentId)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(paymentId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
