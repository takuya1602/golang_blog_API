package service

import (
	"backend/app/common/dto"

	"github.com/stretchr/testify/mock"
)

type ICategoryService struct {
	mock.Mock
}

func (_m *ICategoryService) GetAll() (categories []dto.CategoryModel, err error) {
	ret := _m.Called()

	if rf, ok := ret.Get(0).(func() []dto.CategoryModel); ok {
		categories = rf()
	} else {
		if ret.Get(0) != nil {
			categories = ret.Get(0).([]dto.CategoryModel)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		err = rf()
	} else {
		err = ret.Error(1)
	}
	return
}

func (_m *ICategoryService) GetBySlug(slug string) (category dto.CategoryModel, err error) {
	ret := _m.Called(slug)

	if rf, ok := ret.Get(0).(func(string) dto.CategoryModel); ok {
		category = rf(slug)
	} else {
		if ret.Get(0) != nil {
			category = ret.Get(0).(dto.CategoryModel)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		err = rf(slug)
	} else {
		err = ret.Error(1)
	}
	return
}

func (_m *ICategoryService) Create(categoryDto dto.CategoryModel) (err error) {
	ret := _m.Called(categoryDto)

	if rf, ok := ret.Get(0).(func(dto.CategoryModel) error); ok {
		err = rf(categoryDto)
	} else {
		err = ret.Error(0)
	}
	return
}

func (_m *ICategoryService) Update(categoryDto dto.CategoryModel) (err error) {
	ret := _m.Called(categoryDto)

	if rf, ok := ret.Get(0).(func(dto.CategoryModel) error); ok {
		err = rf(categoryDto)
	} else {
		err = ret.Error(0)
	}
	return
}

func (_m *ICategoryService) Delete(categoryDto dto.CategoryModel) (err error) {
	ret := _m.Called(categoryDto)

	if rf, ok := ret.Get(0).(func(dto.CategoryModel) error); ok {
		err = rf(categoryDto)
	} else {
		err = ret.Error(0)
	}
	return
}
