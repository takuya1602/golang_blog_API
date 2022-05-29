package repository

import (
	"backend/app/domain/entity"

	mock "github.com/stretchr/testify/mock"
)

type ISubCategoryRepository struct {
	mock.Mock
}

func (_m *ISubCategoryRepository) GetAll() (subCategories []entity.SubCategory, err error) {
	ret := _m.Called()

	if rf, ok := ret.Get(0).(func() []entity.SubCategory); ok {
		subCategories = rf()
	} else {
		if ret.Get(0) != nil {
			subCategories = ret.Get(0).([]entity.SubCategory)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		err = rf()
	} else {
		err = ret.Error(1)
	}
	return
}

func (_m *ISubCategoryRepository) GetFilterParentCategory(parentCategoryName string) (subCategories []entity.SubCategory, err error) {
	ret := _m.Called(parentCategoryName)

	if rf, ok := ret.Get(0).(func(string) []entity.SubCategory); ok {
		subCategories = rf(parentCategoryName)
	} else {
		if ret.Get(0) != nil {
			subCategories = ret.Get(0).([]entity.SubCategory)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		err = rf(parentCategoryName)
	} else {
		err = ret.Error(1)
	}
	return
}

func (_m *ISubCategoryRepository) GetBySlug(slug string) (subCategory entity.SubCategory, err error) {
	ret := _m.Called(slug)

	if rf, ok := ret.Get(0).(func(string) entity.SubCategory); ok {
		subCategory = rf(slug)
	} else {
		if ret.Get(0) != nil {
			subCategory = ret.Get(0).(entity.SubCategory)
		}
	}
	return
}

func (_m *ISubCategoryRepository) Create(subCategory entity.SubCategory) (err error) {
	ret := _m.Called(subCategory)

	if rf, ok := ret.Get(0).(func(entity.SubCategory) error); ok {
		err = rf(subCategory)
	} else {
		err = ret.Error(0)
	}
	return
}

func (_m *ISubCategoryRepository) Update(subCategory entity.SubCategory) (err error) {
	ret := _m.Called(subCategory)

	if rf, ok := ret.Get(0).(func(entity.SubCategory) error); ok {
		err = rf(subCategory)
	} else {
		err = ret.Error(0)
	}
	return
}

func (_m *ISubCategoryRepository) Delete(subCategory entity.SubCategory) (err error) {
	ret := _m.Called(subCategory)

	if rf, ok := ret.Get(0).(func(entity.SubCategory) error); ok {
		err = rf(subCategory)
	} else {
		err = ret.Error(0)
	}
	return
}

func (_m *ISubCategoryRepository) GetIdFromParentCategoryName(name string) (id int) {
	ret := _m.Called(name)

	if rf, ok := ret.Get(0).(func(string) int); ok {
		id = rf(name)
	} else {
		if ret.Get(0) != nil {
			id = ret.Get(0).(int)
		}
	}
	return
}

func (_m *ISubCategoryRepository) GetNameFromParentCategoryId(id int) (name string) {
	ret := _m.Called(id)

	if rf, ok := ret.Get(0).(func(int) string); ok {
		name = rf(id)
	} else {
		if ret.Get(0) != nil {
			name = ret.Get(0).(string)
		}
	}
	return
}
