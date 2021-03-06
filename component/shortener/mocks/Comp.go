// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	shortener "shorturl/component/shortener"

	mock "github.com/stretchr/testify/mock"
)

// Comp is an autogenerated mock type for the Comp type
type Comp struct {
	mock.Mock
}

// AccessURL provides a mock function with given fields: input
func (_m *Comp) AccessURL(input *shortener.AccessInput) (*shortener.ShortenOutput, error) {
	ret := _m.Called(input)

	var r0 *shortener.ShortenOutput
	if rf, ok := ret.Get(0).(func(*shortener.AccessInput) *shortener.ShortenOutput); ok {
		r0 = rf(input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*shortener.ShortenOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*shortener.AccessInput) error); ok {
		r1 = rf(input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ShortenURL provides a mock function with given fields: input
func (_m *Comp) ShortenURL(input *shortener.ShortenInput) (*shortener.ShortenOutput, error) {
	ret := _m.Called(input)

	var r0 *shortener.ShortenOutput
	if rf, ok := ret.Get(0).(func(*shortener.ShortenInput) *shortener.ShortenOutput); ok {
		r0 = rf(input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*shortener.ShortenOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*shortener.ShortenInput) error); ok {
		r1 = rf(input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
