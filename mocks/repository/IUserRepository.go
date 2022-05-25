package repository

import (
	"backend/app/domain/entity"

	mock "github.com/stretchr/testify/mock"
)

type IUserRepository struct {
	mock.Mock
}

func (_m *IUserRepository) GetAll() (users []entity.User, err error) {
	ret := _m.Called()

	if rf, ok := ret.Get(0).(func() []entity.User); ok {
		users = rf()
	} else {
		if ret.Get(0) != nil {
			users = ret.Get(0).([]entity.User)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		err = rf()
	} else {
		err = ret.Error(1)
	}
	return
}

func (_m *IUserRepository) ValidateUser(creds entity.Credentials) (user entity.User, err error) {
	ret := _m.Called(creds)

	if rf, ok := ret.Get(0).(func(entity.Credentials) entity.User); ok {
		user = rf(creds)
	} else {
		if ret.Get(0) != nil {
			user = ret.Get(0).(entity.User)
		}
	}

	if rf, ok := ret.Get(1).(func(entity.Credentials) error); ok {
		err = rf(creds)
	} else {
		err = ret.Error(1)
	}
	return
}

func (_m *IUserRepository) Create(user entity.User) (err error) {
	ret := _m.Called(user)

	if rf, ok := ret.Get(0).(func(entity.User) error); ok {
		err = rf(user)
	} else {
		err = ret.Error(0)
	}
	return
}

func (_m *IUserRepository) Update(user entity.User) (err error) {
	ret := _m.Called(user)

	if rf, ok := ret.Get(0).(func(entity.User) error); ok {
		err = rf(user)
	} else {
		err = ret.Error(0)
	}
	return
}

func (_m *IUserRepository) Delete(user entity.User) (err error) {
	ret := _m.Called(user)

	if rf, ok := ret.Get(0).(func(entity.User) error); ok {
		err = rf(user)
	} else {
		err = ret.Error(0)
	}
	return
}

func (_m *IUserRepository) IsAdmin(id int) (isAdmin bool, err error) {
	ret := _m.Called(id)

	if rf, ok := ret.Get(0).(func(int) bool); ok {
		isAdmin = rf(id)
	} else {
		if ret.Get(0) != nil {
			isAdmin = ret.Get(0).(bool)
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		err = rf(id)
	} else {
		err = ret.Error(1)
	}
	return
}
