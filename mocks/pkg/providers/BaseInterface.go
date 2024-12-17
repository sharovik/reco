// Code generated by mockery v2.50.0. DO NOT EDIT.

package mocks

import (
	dto "github.com/reco/pkg/dto"
	mock "github.com/stretchr/testify/mock"
)

// BaseInterface is an autogenerated mock type for the BaseInterface type
type BaseInterface struct {
	mock.Mock
}

// GetUsersList provides a mock function with given fields: page, limit
func (_m *BaseInterface) GetUsersList(page string, limit int) ([]dto.UserDataItem, string, error) {
	ret := _m.Called(page, limit)

	if len(ret) == 0 {
		panic("no return value specified for GetUsersList")
	}

	var r0 []dto.UserDataItem
	var r1 string
	var r2 error
	if rf, ok := ret.Get(0).(func(string, int) ([]dto.UserDataItem, string, error)); ok {
		return rf(page, limit)
	}
	if rf, ok := ret.Get(0).(func(string, int) []dto.UserDataItem); ok {
		r0 = rf(page, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]dto.UserDataItem)
		}
	}

	if rf, ok := ret.Get(1).(func(string, int) string); ok {
		r1 = rf(page, limit)
	} else {
		r1 = ret.Get(1).(string)
	}

	if rf, ok := ret.Get(2).(func(string, int) error); ok {
		r2 = rf(page, limit)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// NewBaseInterface creates a new instance of BaseInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewBaseInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *BaseInterface {
	mock := &BaseInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}