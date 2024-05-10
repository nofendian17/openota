// Code generated by mockery v2.43.0. DO NOT EDIT.

package mocks

import (
	context "context"

	airport "github.com/nofendian17/openota/apigw/internal/delivery/rest/model/request/airport"

	mock "github.com/stretchr/testify/mock"

	responseairport "github.com/nofendian17/openota/apigw/internal/delivery/rest/model/response/airport"
)

// UseCase is an autogenerated mock type for the UseCase type
type UseCase struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, _a1
func (_m *UseCase) Create(ctx context.Context, _a1 airport.Create) error {
	ret := _m.Called(ctx, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, airport.Create) error); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: ctx, ID
func (_m *UseCase) Delete(ctx context.Context, ID string) error {
	ret := _m.Called(ctx, ID)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, ID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields: ctx
func (_m *UseCase) GetAll(ctx context.Context) ([]*responseairport.Airport, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetAll")
	}

	var r0 []*responseairport.Airport
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]*responseairport.Airport, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []*responseairport.Airport); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*responseairport.Airport)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByCode provides a mock function with given fields: ctx, code
func (_m *UseCase) GetByCode(ctx context.Context, code string) (*responseairport.Airport, error) {
	ret := _m.Called(ctx, code)

	if len(ret) == 0 {
		panic("no return value specified for GetByCode")
	}

	var r0 *responseairport.Airport
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*responseairport.Airport, error)); ok {
		return rf(ctx, code)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *responseairport.Airport); ok {
		r0 = rf(ctx, code)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*responseairport.Airport)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, code)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: ctx, ID
func (_m *UseCase) GetByID(ctx context.Context, ID string) (*responseairport.Airport, error) {
	ret := _m.Called(ctx, ID)

	if len(ret) == 0 {
		panic("no return value specified for GetByID")
	}

	var r0 *responseairport.Airport
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*responseairport.Airport, error)); ok {
		return rf(ctx, ID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *responseairport.Airport); ok {
		r0 = rf(ctx, ID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*responseairport.Airport)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, ID, _a2
func (_m *UseCase) Update(ctx context.Context, ID string, _a2 airport.Update) error {
	ret := _m.Called(ctx, ID, _a2)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, airport.Update) error); ok {
		r0 = rf(ctx, ID, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewUseCase creates a new instance of UseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUseCase(t interface {
	mock.TestingT
	Cleanup(func())
}) *UseCase {
	mock := &UseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
