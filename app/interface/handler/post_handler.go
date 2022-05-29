package handler

import (
	"backend/app/common/dto"
	"backend/app/domain/service"
	"encoding/json"
	"net/http"
	"path"
)

type IPostHandler interface {
	GetPosts(w http.ResponseWriter, r *http.Request) (err error)
	GetPostBySlug(w http.ResponseWriter, r *http.Request, slug string) (err error)
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

func (h *PostHandler) GetPosts(w http.ResponseWriter, r *http.Request) (err error) {
	queryParams := r.URL.Query()
	postDtos, err := h.IPostService.GetPosts(queryParams)
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

func (h *PostHandler) GetPostBySlug(w http.ResponseWriter, r *http.Request, slug string) (err error) {
	postDto, err := h.IPostService.GetPostBySlug(slug)
	if err != nil {
		return
	}
	output, err := json.MarshalIndent(&postDto, "", "\t")
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
	postDto, err := h.IPostService.GetPostBySlug(slug)
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	json.Unmarshal(body, &postDto)
	err = h.IPostService.Update(postDto)
	return
}

func (h *PostHandler) Delete(w http.ResponseWriter, r *http.Request) (err error) {
	slug := path.Base(r.URL.Path)
	postDto, err := h.IPostService.GetPostBySlug(slug)
	err = h.IPostService.Delete(postDto)
	return
}
