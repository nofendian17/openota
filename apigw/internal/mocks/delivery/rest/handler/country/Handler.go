// Code generated by mockery v2.43.0. DO NOT EDIT.

package mocks

import (
	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// Handler is an autogenerated mock type for the Handler type
type Handler struct {
	mock.Mock
}

// GetByCode provides a mock function with given fields:
func (_m *Handler) GetByCode() http.HandlerFunc {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetByCode")
	}

	var r0 http.HandlerFunc
	if rf, ok := ret.Get(0).(func() http.HandlerFunc); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(http.HandlerFunc)
		}
	}

	return r0
}

// GetByID provides a mock function with given fields:
func (_m *Handler) GetByID() http.HandlerFunc {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetByID")
	}

	var r0 http.HandlerFunc
	if rf, ok := ret.Get(0).(func() http.HandlerFunc); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(http.HandlerFunc)
		}
	}

	return r0
}

// NewHandler creates a new instance of Handler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *Handler {
	mock := &Handler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
