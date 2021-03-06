// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	admin "shorturl/component/admin"

	mock "github.com/stretchr/testify/mock"
)

// Comp is an autogenerated mock type for the Comp type
type Comp struct {
	mock.Mock
}

// Delete provides a mock function with given fields: input
func (_m *Comp) Delete(input *admin.DelInput) error {
	ret := _m.Called(input)

	var r0 error
	if rf, ok := ret.Get(0).(func(*admin.DelInput) error); ok {
		r0 = rf(input)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// List provides a mock function with given fields: input
func (_m *Comp) List(input *admin.ListInput) (*admin.ListOutput, error) {
	ret := _m.Called(input)

	var r0 *admin.ListOutput
	if rf, ok := ret.Get(0).(func(*admin.ListInput) *admin.ListOutput); ok {
		r0 = rf(input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.ListOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*admin.ListInput) error); ok {
		r1 = rf(input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
