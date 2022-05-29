package service

import (
	"backend/app/common/dto"

	"github.com/stretchr/testify/mock"
)

type IPostService struct {
	mock.Mock
}

func (_m *IPostService) GetPosts(queryParams map[string][]string) (postDtos []dto.PostModel, err error) {
	ret := _m.Called(queryParams)

	if rf, ok := ret.Get(0).(func(map[string][]string) []dto.PostModel); ok {
		postDtos = rf(queryParams)
	} else {
		if ret.Get(0) != nil {
			postDtos = ret.Get(0).([]dto.PostModel)
		} else {
		}
	}

	if rf, ok := ret.Get(1).(func(map[string][]string) error); ok {
		err = rf(queryParams)
	} else {
		err = ret.Error(1)
	}
	return
}

func (_m *IPostService) GetPostBySlug(slug string) (postDto dto.PostModel, err error) {
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
