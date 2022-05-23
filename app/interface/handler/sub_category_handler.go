package handler

import (
	"backend/app/common/dto"
	"backend/app/domain/service"
	"encoding/json"
	"net/http"
	"path"
)

type ISubCategoryHandler interface {
	GetAll(w http.ResponseWriter, r *http.Request) error
	Create(w http.ResponseWriter, r *http.Request) error
	Update(w http.ResponseWriter, r *http.Request) error
	Delete(w http.ResponseWriter, r *http.Request) error
}

type SubCategoryHandler struct {
	service.ISubCategoryService
}

func NewSubCategoryHandler(srv service.ISubCategoryService) (iSubCategoryHandler ISubCategoryHandler) {
	iSubCategoryHandler = &SubCategoryHandler{srv}
	return
}

func (h *SubCategoryHandler) GetAll(w http.ResponseWriter, r *http.Request) (err error) {
	subCategories, err := h.ISubCategoryService.GetAll()
	if err != nil {
		return
	}
	output, err := json.MarshalIndent(&subCategories, "", "\t")
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}

func (h *SubCategoryHandler) Create(w http.ResponseWriter, r *http.Request) (err error) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	var subCategoryDto dto.SubCategoryModel
	json.Unmarshal(body, &subCategoryDto)
	err = h.ISubCategoryService.Create(subCategoryDto)
	return
}

func (h *SubCategoryHandler) Update(w http.ResponseWriter, r *http.Request) (err error) {
	slug := path.Base(r.URL.Path)
	subCategoryDto, err := h.ISubCategoryService.GetBySlug(slug)
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	json.Unmarshal(body, &subCategoryDto)
	err = h.ISubCategoryService.Update(subCategoryDto)
	return
}

func (h *SubCategoryHandler) Delete(w http.ResponseWriter, r *http.Request) (err error) {
	slug := path.Base(r.URL.Path)
	subCategoryDto, err := h.ISubCategoryService.GetBySlug(slug)
	err = h.ISubCategoryService.Delete(subCategoryDto)
	return
}
