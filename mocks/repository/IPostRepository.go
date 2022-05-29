package repository

import (
	"backend/app/domain/entity"

	mock "github.com/stretchr/testify/mock"
)

type IPostRepository struct {
	mock.Mock
}

func (_m *IPostRepository) GetPosts(queryParams map[string][]string) (posts []entity.Post, err error) {
	ret := _m.Called(queryParams)

	if rf, ok := ret.Get(0).(func(map[string][]string) []entity.Post); ok {
		posts = rf(queryParams)
	} else {
		if ret.Get(0) != nil {
			posts = ret.Get(0).([]entity.Post)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		err = rf()
	} else {
		err = ret.Error(1)
	}
	return
}

func (_m *IPostRepository) GetPostBySlug(slug string) (post entity.Post, err error) {
	ret := _m.Called(slug)

	if rf, ok := ret.Get(0).(func(string) entity.Post); ok {
		post = rf(slug)
	} else {
		if ret.Get(0) != nil {
			post = ret.Get(0).(entity.Post)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		err = rf(slug)
	} else {
		err = ret.Error(1)
	}
	return
}

func (_m *IPostRepository) Create(post entity.Post) (err error) {
	ret := _m.Called(post)

	if rf, ok := ret.Get(0).(func(entity.Post) error); ok {
		err = rf(post)
	} else {
		err = ret.Error(0)
	}
	return
}

func (_m *IPostRepository) Update(post entity.Post) (err error) {
	ret := _m.Called(post)

	if rf, ok := ret.Get(0).(func(entity.Post) error); ok {
		err = rf(post)
	} else {
		err = ret.Error(0)
	}
	return
}

func (_m *IPostRepository) Delete(post entity.Post) (err error) {
	ret := _m.Called(post)

	if rf, ok := ret.Get(0).(func(entity.Post) error); ok {
		err = rf(post)
	} else {
		err = ret.Error(0)
	}
	return
}
