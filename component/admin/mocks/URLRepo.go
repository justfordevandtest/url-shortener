// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	entity "shorturl/entity"

	mock "github.com/stretchr/testify/mock"
)

// URLRepo is an autogenerated mock type for the URLRepo type
type URLRepo struct {
	mock.Mock
}

// Delete provides a mock function with given fields: ID
func (_m *URLRepo) Delete(ID string) error {
	ret := _m.Called(ID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(ID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// List provides a mock function with given fields: page, perPage, filters
func (_m *URLRepo) List(page int, perPage int, filters map[string]interface{}) (int, []entity.URL, error) {
	ret := _m.Called(page, perPage, filters)

	var r0 int
	if rf, ok := ret.Get(0).(func(int, int, map[string]interface{}) int); ok {
		r0 = rf(page, perPage, filters)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 []entity.URL
	if rf, ok := ret.Get(1).(func(int, int, map[string]interface{}) []entity.URL); ok {
		r1 = rf(page, perPage, filters)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]entity.URL)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(int, int, map[string]interface{}) error); ok {
		r2 = rf(page, perPage, filters)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Read provides a mock function with given fields: ID
func (_m *URLRepo) Read(ID string) (*entity.URL, error) {
	ret := _m.Called(ID)

	var r0 *entity.URL
	if rf, ok := ret.Get(0).(func(string) *entity.URL); ok {
		r0 = rf(ID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.URL)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
