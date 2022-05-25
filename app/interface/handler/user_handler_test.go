package handler

import (
	"backend/app/common/dto"
	"net/http/httptest"
	"strings"
	"testing"

	mocks "backend/mocks/service"

	"github.com/stretchr/testify/assert"
)

func TestUserHandler_IssueToken(t *testing.T) {
	credsDto := dto.NewCredsModel("testuser1", "testpass1")
	json := strings.NewReader(`{
		"username": "testuser1",
		"password": "testpass1"
	}`)
	userDto := dto.NewUserModel(1, "testuser1", "testpass1")
	authTokenDto := dto.NewAuthTokenModel("token")

	s := new(mocks.IUserService)

	s.On("ValidateUser", credsDto).Return(userDto, nil)
	s.On("IssueToken", userDto.Id).Return(authTokenDto, nil)

	h := NewUserHandler(s)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/admin/", json)

	err := h.IssueToken(w, r)

	assert.NoError(t, err)
	s.AssertExpectations(t)
}
