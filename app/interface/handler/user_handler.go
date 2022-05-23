package handler

import (
	"backend/app/common/dto"
	"backend/app/domain/service"
	"encoding/json"
	"net/http"
)

type IUserHandler interface {
	IssueToken(w http.ResponseWriter, r *http.Request) error
	ValidateToken(w http.ResponseWriter, r *http.Request) (bool, error)
}

type UserHandler struct {
	service.IUserService
}

func NewUserHandler(srv service.IUserService) (iUserHandler IUserHandler) {
	iUserHandler = &UserHandler{srv}
	return
}

func (h *UserHandler) IssueToken(w http.ResponseWriter, r *http.Request) (err error) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	var credsDto dto.CredentialsModel
	json.Unmarshal(body, &credsDto)

	userDto, err := h.IUserService.ValidateUser(credsDto)
	if err != nil {
		return
	}

	authTokenDto, err := h.IUserService.IssueToken(userDto.Id)
	if err != nil {
		return
	}

	output, err := json.MarshalIndent(&authTokenDto, "", "\t")
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	w.WriteHeader(200)
	return
}

func (h *UserHandler) ValidateToken(w http.ResponseWriter, r *http.Request) (isAdmin bool, err error) {
	authToken := r.Header.Get("Authorization")
	if authToken == "" {
		isAdmin = false
		return
	}
	authTokenDto := dto.NewAuthTokenModel(authToken)
	isAdmin, err = h.IUserService.ValidateToken(authTokenDto)
	return
}
