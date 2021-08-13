// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	entity "shorturl/entity"

	mock "github.com/stretchr/testify/mock"
)

// URLCache is an autogenerated mock type for the URLCache type
type URLCache struct {
	mock.Mock
}

// Read provides a mock function with given fields: ID
func (_m *URLCache) Read(ID string) *entity.URL {
	ret := _m.Called(ID)

	var r0 *entity.URL
	if rf, ok := ret.Get(0).(func(string) *entity.URL); ok {
		r0 = rf(ID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.URL)
		}
	}

	return r0
}

// Write provides a mock function with given fields: ent
func (_m *URLCache) Write(ent *entity.URL) error {
	ret := _m.Called(ent)

	var r0 error
	if rf, ok := ret.Get(0).(func(*entity.URL) error); ok {
		r0 = rf(ent)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
