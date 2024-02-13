// Code generated by mockery v2.10.4. DO NOT EDIT.

package mocks

import (
	dto "git.garena.com/sea-labs-id/bootcamp/batch-03/frederik-hutabarat/exercise-library-api/dto"
	mock "github.com/stretchr/testify/mock"
)

// BookUseCase is an autogenerated mock type for the BookUseCase type
type BookUseCase struct {
	mock.Mock
}

// CreateBook provides a mock function with given fields: body, authorId
func (_m *BookUseCase) CreateBook(body dto.CreateBookBody, authorId int64) (*dto.Book, error) {
	ret := _m.Called(body, authorId)

	var r0 *dto.Book
	if rf, ok := ret.Get(0).(func(dto.CreateBookBody, int64) *dto.Book); ok {
		r0 = rf(body, authorId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.Book)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(dto.CreateBookBody, int64) error); ok {
		r1 = rf(body, authorId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBooks provides a mock function with given fields: title
func (_m *BookUseCase) GetBooks(title string) ([]dto.BookDetail, error) {
	ret := _m.Called(title)

	var r0 []dto.BookDetail
	if rf, ok := ret.Get(0).(func(string) []dto.BookDetail); ok {
		r0 = rf(title)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]dto.BookDetail)
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
