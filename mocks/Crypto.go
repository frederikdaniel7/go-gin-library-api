// Code generated by mockery v2.10.4. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Crypto is an autogenerated mock type for the Crypto type
type Crypto struct {
	mock.Mock
}

// CheckPassword provides a mock function with given fields: password, hash
func (_m *Crypto) CheckPassword(password string, hash []byte) (bool, error) {
	ret := _m.Called(password, hash)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, []byte) bool); ok {
		r0 = rf(password, hash)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, []byte) error); ok {
		r1 = rf(password, hash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HashPassword provides a mock function with given fields: password, cost
func (_m *Crypto) HashPassword(password string, cost int) ([]byte, error) {
	ret := _m.Called(password, cost)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(string, int) []byte); ok {
		r0 = rf(password, cost)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, int) error); ok {
		r1 = rf(password, cost)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}