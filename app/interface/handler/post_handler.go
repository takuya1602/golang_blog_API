package handler

import (
	"backend/app/common/dto"
	"backend/app/domain/service"
	"encoding/json"
	"net/http"
	"path"
)

type IPostHandler interface {
	Get(w http.ResponseWriter, r *http.Request) (err error)
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

func (h *PostHandler) Get(w http.ResponseWriter, r *http.Request) (err error) {
	var postDtos []dto.PostModel
	queryParams := r.URL.Query()
	if categorySlugs, ok := queryParams["category-name"]; ok {
		categorySlug := categorySlugs[0]
		postDtos, err = h.IPostService.GetWithCategoryQuery(categorySlug)
		if err != nil {
			return
		}
	} else if subCategorySlugs, ok := queryParams["sub-category-name"]; ok {
		subCategorySlug := subCategorySlugs[0]
		postDtos, err = h.IPostService.GetWithSubCategoryQuery(subCategorySlug)
		if err != nil {
			return
		}
	} else {
		postDtos, err = h.IPostService.GetAll()
		if err != nil {
			return
		}
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
