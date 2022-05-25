package service

import (
	"backend/app/common/dto"

	"github.com/stretchr/testify/mock"
)

type IPostService struct {
	mock.Mock
}

func (_m *IPostService) GetAll() (postDtos []dto.PostModel, err error) {
	ret := _m.Called()

	if rf, ok := ret.Get(0).(func() []dto.PostModel); ok {
		postDtos = rf()
	} else {
		if ret.Get(0) != nil {
			postDtos = ret.Get(0).([]dto.PostModel)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		err = rf()
	} else {
		err = ret.Error(1)
	}
	return
}

func (_m *IPostService) GetBySlug(slug string) (postDto dto.PostModel, err error) {
	ret := _m.Called(slug)

	if rf, ok := ret.Get(0).(func(string) dto.PostModel); ok {
		postDto = rf(slug)
	} else {
		postDto = ret.Get(0).(dto.PostModel)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		err = rf(slug)
	} else {
		err = ret.Error(1)
	}
	return
}

func (_m *IPostService) GetWithCategoryQuery(categoryName string) (postDtos []dto.PostModel, err error) {
	ret := _m.Called(categoryName)

	if rf, ok := ret.Get(0).(func(string) []dto.PostModel); ok {
		postDtos = rf(categoryName)
	} else {
		postDtos = ret.Get(0).([]dto.PostModel)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		err = rf(categoryName)
	} else {
		err = ret.Error(1)
	}
	return
}

func (_m *IPostService) GetWithSubCategoryQuery(subCategoryName string) (postDtos []dto.PostModel, err error) {
	ret := _m.Called(subCategoryName)

	if rf, ok := ret.Get(0).(func(string) []dto.PostModel); ok {
		postDtos = rf(subCategoryName)
	} else {
		postDtos = ret.Get(0).([]dto.PostModel)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		err = rf(subCategoryName)
	} else {
		err = ret.Error(1)
	}
	return
}

func (_m *IPostService) Create(postDto dto.PostModel) (err error) {
	ret := _m.Called(postDto)

	if rf, ok := ret.Get(0).(func(dto.PostModel) error); ok {
		err = rf(postDto)
	} else {
		err = ret.Error(0)
	}
	return
}

func (_m *IPostService) Update(postDto dto.PostModel) (err error) {
	ret := _m.Called(postDto)

	if rf, ok := ret.Get(0).(func(dto.PostModel) error); ok {
		err = rf(postDto)
	} else {
		err = ret.Error(0)
	}
	return
}

func (_m *IPostService) Delete(postDto dto.PostModel) (err error) {
	ret := _m.Called(postDto)

	if rf, ok := ret.Get(0).(func(dto.PostModel) error); ok {
		err = rf(postDto)
	} else {
		err = ret.Error(0)
	}
	return
}
