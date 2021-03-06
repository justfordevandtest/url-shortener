// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	auth "shorturl/component/auth"

	mock "github.com/stretchr/testify/mock"
)

// Comp is an autogenerated mock type for the Comp type
type Comp struct {
	mock.Mock
}

// ReadByCredential provides a mock function with given fields: input
func (_m *Comp) ReadByCredential(input *auth.CredentialInput) (*auth.UserOutput, error) {
	ret := _m.Called(input)

	var r0 *auth.UserOutput
	if rf, ok := ret.Get(0).(func(*auth.CredentialInput) *auth.UserOutput); ok {
		r0 = rf(input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*auth.UserOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*auth.CredentialInput) error); ok {
		r1 = rf(input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
