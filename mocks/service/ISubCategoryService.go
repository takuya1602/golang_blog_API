package service

import (
	"backend/app/common/dto"

	"github.com/stretchr/testify/mock"
)

type ISubCategoryService struct {
	mock.Mock
}

func (_m *ISubCategoryService) GetAll() (subCategoryDtos []dto.SubCategoryModel, err error) {
	ret := _m.Called()

	if rf, ok := ret.Get(0).(func() []dto.SubCategoryModel); ok {
		subCategoryDtos = rf()
	} else {
		if ret.Get(0) != nil {
			subCategoryDtos = ret.Get(0).([]dto.SubCategoryModel)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		err = rf()
	} else {
		err = ret.Error(1)
	}
	return
}

func (_m *ISubCategoryService) GetWithQuery(categoryName string) (subCategoryDtos []dto.SubCategoryModel, err error) {
	ret := _m.Called(categoryName)

	if rf, ok := ret.Get(0).(func(string) []dto.SubCategoryModel); ok {
		subCategoryDtos = rf(categoryName)
	} else {
		if ret.Get(0) != nil {
			subCategoryDtos = ret.Get(0).([]dto.SubCategoryModel)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		err = rf(categoryName)
	} else {
		err = ret.Error(1)
	}
	return
}

func (_m *ISubCategoryService) GetBySlug(slug string) (subCategoryDto dto.SubCategoryModel, err error) {
	ret := _m.Called(slug)

	if rf, ok := ret.Get(0).(func(string) dto.SubCategoryModel); ok {
		subCategoryDto = rf(slug)
	} else {
		if ret.Get(0) != nil {
			subCategoryDto = ret.Get(0).(dto.SubCategoryModel)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		err = rf(slug)
	} else {
		err = ret.Error(1)
	}
	return
}

func (_m ISubCategoryService) Create(subCategoryDto dto.SubCategoryModel) (err error) {
	ret := _m.Called(subCategoryDto)

	if rf, ok := ret.Get(0).(func(dto.SubCategoryModel) error); ok {
		err = rf(subCategoryDto)
	} else {
		err = ret.Error(0)
	}
	return
}

func (_m ISubCategoryService) Update(subCategoryDto dto.SubCategoryModel) (err error) {
	ret := _m.Called(subCategoryDto)

	if rf, ok := ret.Get(0).(func(dto.SubCategoryModel) error); ok {
		err = rf(subCategoryDto)
	} else {
		err = ret.Error(0)
	}
	return
}

func (_m ISubCategoryService) Delete(subCategoryDto dto.SubCategoryModel) (err error) {
	ret := _m.Called(subCategoryDto)

	if rf, ok := ret.Get(0).(func(dto.SubCategoryModel) error); ok {
		err = rf(subCategoryDto)
	} else {
		err = ret.Error(0)
	}
	return
}
