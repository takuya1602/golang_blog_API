package repository

import (
	"backend/app/domain/entity"

	mock "github.com/stretchr/testify/mock"
)

type IPostRepository struct {
	mock.Mock
}

func (_m *IPostRepository) GetAll() (posts []entity.Post, err error) {
	ret := _m.Called()

	if rf, ok := ret.Get(0).(func() []entity.Post); ok {
		posts = rf()
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

func (_m *IPostRepository) GetBySlug(slug string) (post entity.Post, err error) {
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

func (_m *IPostRepository) GetFilterCategory(categoryName string) (posts []entity.Post, err error) {
	ret := _m.Called(categoryName)

	if rf, ok := ret.Get(0).(func(string) []entity.Post); ok {
		posts = rf(categoryName)
	} else {
		if ret.Get(0) != nil {
			posts = ret.Get(0).([]entity.Post)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		err = rf(categoryName)
	} else {
		err = ret.Error(1)
	}
	return
}

func (_m *IPostRepository) GetFilterSubCategory(subCategoryName string) (posts []entity.Post, err error) {
	ret := _m.Called(subCategoryName)

	if rf, ok := ret.Get(0).(func(string) []entity.Post); ok {
		posts = rf(subCategoryName)
	} else {
		if ret.Get(0) != nil {
			posts = ret.Get(0).([]entity.Post)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		err = rf(subCategoryName)
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
