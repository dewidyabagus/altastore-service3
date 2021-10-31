// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	checkoutpayment "AltaStore/business/checkoutpayment"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// CheckHasCheckoutId provides a mock function with given fields: id
func (_m *Repository) CheckHasCheckoutId(id string) (bool, error) {
	ret := _m.Called(id)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPaymentByCheckoutId provides a mock function with given fields: id
func (_m *Repository) GetPaymentByCheckoutId(id string) (*checkoutpayment.CheckoutPayment, error) {
	ret := _m.Called(id)

	var r0 *checkoutpayment.CheckoutPayment
	if rf, ok := ret.Get(0).(func(string) *checkoutpayment.CheckoutPayment); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*checkoutpayment.CheckoutPayment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertPayment provides a mock function with given fields: payment
func (_m *Repository) InsertPayment(payment *checkoutpayment.CheckoutPayment) (*checkoutpayment.CheckoutPayment, error) {
	ret := _m.Called(payment)

	var r0 *checkoutpayment.CheckoutPayment
	if rf, ok := ret.Get(0).(func(*checkoutpayment.CheckoutPayment) *checkoutpayment.CheckoutPayment); ok {
		r0 = rf(payment)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*checkoutpayment.CheckoutPayment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*checkoutpayment.CheckoutPayment) error); ok {
		r1 = rf(payment)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
