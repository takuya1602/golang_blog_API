package handler

import (
	"backend/app/common/dto"
	"backend/app/domain/service"
	"encoding/json"
	"net/http"
	"path"
)

type IPostHandler interface {
	GetAll(w http.ResponseWriter, r *http.Request) (err error)
	Create(w http.ResponseWriter, r *http.Request) (err error)
	Update(w http.ResponseWriter, r *http.Request) (err error)
	Delete(w http.ResponseWriter, r *http.Request) (err error)
}

type PostHandler struct {
	service.IPostService
}

func NewPostHandler(srv service.IPostService) (iPostHandler IPostHandler) {
	iPostHandler = &PostHandler{srv}
	return
}

func (h *PostHandler) GetAll(w http.ResponseWriter, r *http.Request) (err error) {
	postDtos, err := h.IPostService.GetAll()
	if err != nil {
		return
	}
	output, err := json.MarshalIndent(&postDtos, "", "\t")
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}

func (h *PostHandler) Create(w http.ResponseWriter, r *http.Request) (err error) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	var postDto dto.PostModel
	json.Unmarshal(body, &postDto)
	err = h.IPostService.Create(postDto)
	return
}

func (h *PostHandler) Update(w http.ResponseWriter, r *http.Request) (err error) {
	slug := path.Base(r.URL.Path)
	postDto, err := h.IPostService.GetBySlug(slug)
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	json.Unmarshal(body, &postDto)
	err = h.IPostService.Update(postDto)
	return
}

func (h *PostHandler) Delete(w http.ResponseWriter, r *http.Request) (err error) {
	slug := path.Base(r.URL.Path)
	postDto, err := h.IPostService.GetBySlug(slug)
	err = h.IPostService.Delete(postDto)
	return
}
