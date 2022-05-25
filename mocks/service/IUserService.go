package service

import (
	"backend/app/common/dto"

	"github.com/stretchr/testify/mock"
)

type IUserService struct {
	mock.Mock
}

func (_m *IUserService) GetAll() (userDtos []dto.UserModel, err error) {
	ret := _m.Called()

	if rf, ok := ret.Get(0).(func() []dto.UserModel); ok {
		userDtos = rf()
	} else {
		if ret.Get(0) != nil {
			userDtos = ret.Get(0).([]dto.UserModel)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		err = rf()
	} else {
		err = ret.Error(1)
	}
	return
}

func (_m *IUserService) ValidateUser(credsDto dto.CredentialsModel) (userDto dto.UserModel, err error) {
	ret := _m.Called(credsDto)

	if rf, ok := ret.Get(0).(func(dto.CredentialsModel) dto.UserModel); ok {
		userDto = rf(credsDto)
	} else {
		if ret.Get(0) != nil {
			userDto = ret.Get(0).(dto.UserModel)
		}
	}

	if rf, ok := ret.Get(1).(func(dto.CredentialsModel) error); ok {
		err = rf(credsDto)
	} else {
		err = ret.Error(1)
	}
	return
}

func (_m *IUserService) Create(userDto dto.UserModel) (err error) {
	ret := _m.Called(userDto)

	if rf, ok := ret.Get(0).(func(dto.UserModel) error); ok {
		err = rf(userDto)
	} else {
		err = ret.Error(0)
	}
	return
}

func (_m *IUserService) Update(userDto dto.UserModel) (err error) {
	ret := _m.Called(userDto)

	if rf, ok := ret.Get(0).(func(dto.UserModel) error); ok {
		err = rf(userDto)
	} else {
		err = ret.Error(0)
	}
	return
}

func (_m *IUserService) Delete(userDto dto.UserModel) (err error) {
	ret := _m.Called(userDto)

	if rf, ok := ret.Get(0).(func(dto.UserModel) error); ok {
		err = rf(userDto)
	} else {
		err = ret.Error(0)
	}
	return
}

func (_m *IUserService) IssueToken(id int) (authTokenDto dto.AuthTokenModel, err error) {
	ret := _m.Called(id)

	if rf, ok := ret.Get(0).(func(int) dto.AuthTokenModel); ok {
		authTokenDto = rf(id)
	} else {
		if ret.Get(0) != nil {
			authTokenDto = ret.Get(0).(dto.AuthTokenModel)
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		err = rf(id)
	} else {
		err = ret.Error(1)
	}
	return
}

func (_m *IUserService) ValidateToken(authTokenDto dto.AuthTokenModel) (isAdmin bool, err error) {
	ret := _m.Called(authTokenDto)

	if rf, ok := ret.Get(0).(func(dto.AuthTokenModel) bool); ok {
		isAdmin = rf(authTokenDto)
	} else {
		if ret.Get(0) != nil {
			isAdmin = ret.Get(0).(bool)
		}
	}

	if rf, ok := ret.Get(1).(func(dto.AuthTokenModel) error); ok {
		err = rf(authTokenDto)
	} else {
		err = ret.Error(1)
	}
	return
}
