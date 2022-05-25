package service

import (
	"backend/app/common/dto"
	"backend/app/domain/entity"
	mocks "backend/mocks/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserService_GetAll(t *testing.T) {
	users := []entity.User{
		{
			Id:       1,
			Name:     "testuser1",
			Password: "testpass1",
			IsAdmin:  true,
		},
		{
			Id:       2,
			Name:     "testuser2",
			Password: "testpass2",
			IsAdmin:  true,
		},
	}

	r := new(mocks.IUserRepository)

	r.On("GetAll").Return(users, nil)

	s := NewUserService(r)

	ret, err := s.GetAll()

	assert.NoError(t, err)
	for i, r := range ret {
		assert.Equal(t, r.Id, users[i].Id)
		assert.Equal(t, r.Name, users[i].Name)
		assert.Equal(t, r.Password, users[i].Password)
		assert.Equal(t, r.IsAdmin, users[i].IsAdmin)
	}
	r.AssertExpectations(t)
}

func TestUserService_ValidateUser(t *testing.T) {
	user := entity.NewUser(1, "testuser1", "testpass1")

	creds := entity.NewCreds("testuser1", "testpass1")
	credsDto := dto.NewCredsModel("testuser1", "testpass1")

	r := new(mocks.IUserRepository)

	r.On("ValidateUser", creds).Return(user, nil)

	s := NewUserService(r)

	ret, err := s.ValidateUser(credsDto)

	assert.NoError(t, err)
	assert.Equal(t, ret.Id, user.Id)
	assert.Equal(t, ret.Name, user.Name)
	assert.Equal(t, ret.Password, user.Password)
	assert.Equal(t, ret.IsAdmin, user.IsAdmin)
	r.AssertExpectations(t)
}

func TestUserService_Create(t *testing.T) {
	user := entity.NewUser(1, "testuser1", "testpass1")
	userDto := dto.NewUserModel(1, "testuser1", "testpass1")

	r := new(mocks.IUserRepository)

	r.On("Create", user).Return(nil)

	s := NewUserService(r)

	err := s.Create(userDto)

	assert.NoError(t, err)
	r.AssertExpectations(t)
}

func TestUserService_Update(t *testing.T) {
	user := entity.NewUser(1, "testuser1", "testpass1")
	userDto := dto.NewUserModel(1, "testuser1", "testpass1")

	r := new(mocks.IUserRepository)

	r.On("Update", user).Return(nil)

	s := NewUserService(r)

	err := s.Update(userDto)

	assert.NoError(t, err)
	r.AssertExpectations(t)
}

func TestUserService_Delete(t *testing.T) {
	user := entity.NewUser(1, "testuser1", "testpass1")
	userDto := dto.NewUserModel(1, "testuser1", "testpass1")

	r := new(mocks.IUserRepository)

	r.On("Delete", user).Return(nil)

	s := NewUserService(r)

	err := s.Delete(userDto)

	assert.NoError(t, err)
	r.AssertExpectations(t)
}
