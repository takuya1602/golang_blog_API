package repository

import (
	"backend/app/domain/entity"

	mock "github.com/stretchr/testify/mock"
)

type ICategoryRepository struct {
	mock.Mock
}

func (_m *ICategoryRepository) GetAll() (categories []entity.Category, err error) {
	ret := _m.Called()

	if rf, ok := ret.Get(0).(func() []entity.Category); ok {
		categories = rf()
	} else {
		if ret.Get(0) != nil {
			categories = ret.Get(0).([]entity.Category)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		err = rf()
	} else {
		err = ret.Error(1)
	}
	return
}

func (_m *ICategoryRepository) GetBySlug(slug string) (category entity.Category, err error) {
	ret := _m.Called(slug)

	if rf, ok := ret.Get(0).(func(string) entity.Category); ok {
		category = rf(slug)
	} else {
		if ret.Get(0) != nil {
			category = ret.Get(0).(entity.Category)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		err = rf(slug)
	} else {
		err = ret.Error(1)
	}
	return
}

func (_m *ICategoryRepository) Create(category entity.Category) (err error) {
	ret := _m.Called(category)

	if rf, ok := ret.Get(0).(func(entity.Category) error); ok {
		err = rf(category)
	} else {
		err = ret.Error(0)
	}
	return
}

func (_m *ICategoryRepository) Update(category entity.Category) (err error) {
	ret := _m.Called(category)

	if rf, ok := ret.Get(0).(func(entity.Category) error); ok {
		err = rf(category)
	} else {
		err = ret.Error(0)
	}
	return
}

func (_m *ICategoryRepository) Delete(category entity.Category) (err error) {
	ret := _m.Called(category)

	if rf, ok := ret.Get(0).(func(entity.Category) error); ok {
		err = rf(category)
	} else {
		err = ret.Error(0)
	}
	return
}
