package handler

import (
	"path"

	"backend/app/common/dto"
	"backend/app/domain/service"
	"encoding/json"
	"net/http"
)

type ICategoryHandler interface {
	GetAll(w http.ResponseWriter, r *http.Request) (err error)
	Create(w http.ResponseWriter, r *http.Request) (err error)
	Update(w http.ResponseWriter, r *http.Request) (err error)
	Delete(w http.ResponseWriter, r *http.Request) (err error)
}

type CategoryHandler struct {
	service.ICategoryService
}

func NewCategoryHandler(srv service.ICategoryService) (iCategoryHandler ICategoryHandler) {
	iCategoryHandler = &CategoryHandler{srv}
	return
}

func (h *CategoryHandler) GetAll(w http.ResponseWriter, r *http.Request) (err error) {
	categories, err := h.ICategoryService.GetAll()
	if err != nil {
		return
	}
	output, err := json.MarshalIndent(&categories, "", "\t")
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}

func (h *CategoryHandler) Create(w http.ResponseWriter, r *http.Request) (err error) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	var categoryDto dto.CategoryModel
	json.Unmarshal(body, &categoryDto)
	err = h.ICategoryService.Create(categoryDto)
	return
}

func (h *CategoryHandler) Update(w http.ResponseWriter, r *http.Request) (err error) {
	slug := path.Base(r.URL.Path)
	categoryDto, err := h.ICategoryService.GetBySlug(slug)
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	json.Unmarshal(body, &categoryDto)
	err = h.ICategoryService.Update(categoryDto)
	return
}

func (h *CategoryHandler) Delete(w http.ResponseWriter, r *http.Request) (err error) {
	slug := path.Base(r.URL.Path)
	categoryDto, err := h.ICategoryService.GetBySlug(slug)
	err = h.ICategoryService.Delete(categoryDto)
	return
}
