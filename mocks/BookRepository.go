// Code generated by mockery v2.10.4. DO NOT EDIT.

package mocks

import (
	dto "git.garena.com/sea-labs-id/bootcamp/batch-03/frederik-hutabarat/exercise-library-api/dto"
	entity "git.garena.com/sea-labs-id/bootcamp/batch-03/frederik-hutabarat/exercise-library-api/entity"

	mock "github.com/stretchr/testify/mock"
)

// BookRepository is an autogenerated mock type for the BookRepository type
type BookRepository struct {
	mock.Mock
}

// CreateBook provides a mock function with given fields: body
func (_m *BookRepository) CreateBook(body dto.CreateBookBody) (*entity.Book, error) {
	ret := _m.Called(body)

	var r0 *entity.Book
	if rf, ok := ret.Get(0).(func(dto.CreateBookBody) *entity.Book); ok {
		r0 = rf(body)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Book)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(dto.CreateBookBody) error); ok {
		r1 = rf(body)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindAll provides a mock function with given fields:
func (_m *BookRepository) FindAll() ([]entity.Book, error) {
	ret := _m.Called()

	var r0 []entity.Book
	if rf, ok := ret.Get(0).(func() []entity.Book); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Book)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindOneBookByTitle provides a mock function with given fields: title
func (_m *BookRepository) FindOneBookByTitle(title string) ([]entity.Book, error) {
	ret := _m.Called(title)

	var r0 []entity.Book
	if rf, ok := ret.Get(0).(func(string) []entity.Book); ok {
		r0 = rf(title)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Book)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(title)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
