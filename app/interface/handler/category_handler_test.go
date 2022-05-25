package handler

import (
	"backend/app/common/dto"
	"net/http/httptest"
	"strings"
	"testing"

	mocks "backend/mocks/service"

	"github.com/stretchr/testify/assert"
)

func TestCategoryHandler_GetAll(t *testing.T) {
	categoryDtos := []dto.CategoryModel{
		{
			Id:   1,
			Name: "testCategory1",
			Slug: "test-category-1",
		},
		{
			Id:   2,
			Name: "testCategory2",
			Slug: "test-category-2",
		},
	}

	s := new(mocks.ICategoryService)

	s.On("GetAll").Return(categoryDtos, nil)

	h := NewCategoryHandler(s)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/categories/", nil)

	err := h.GetAll(w, r)

	assert.NoError(t, err)
	s.AssertExpectations(t)
}

func TestCategoryHandler_Create(t *testing.T) {
	categoryDto := dto.NewCategoryModel(1, "testCategory1", "test-category-1")
	json := strings.NewReader(`{
		"id": 1,
		"name": "testCategory1",
		"slug": "test-category-1"
	}`)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/categories/", json)

	s := new(mocks.ICategoryService)

	s.On("Create", categoryDto).Return(nil)

	h := NewCategoryHandler(s)

	err := h.Create(w, r)

	assert.NoError(t, err)
	s.AssertExpectations(t)
}

func TestCategoryHandler_Update(t *testing.T) {
	categoryDto := dto.NewCategoryModel(1, "testCategory1", "test-category-1")
	json := strings.NewReader(`{
		"id": 1,
		"name": "testCategory1",
		"slug": "test-category-1"
	}`)
	slug := "test-category-1"

	w := httptest.NewRecorder()
	r := httptest.NewRequest("PUT", "/categories/test-category-1/", json)

	s := new(mocks.ICategoryService)

	s.On("GetBySlug", slug).Return(categoryDto, nil)
	s.On("Update", categoryDto).Return(nil)

	h := NewCategoryHandler(s)

	err := h.Update(w, r)

	assert.NoError(t, err)
	s.AssertExpectations(t)
}

func TestCategoryHandler_Delete(t *testing.T) {
	categoryDto := dto.NewCategoryModel(1, "testCategory1", "test-category-1")
	slug := "test-category-1"

	w := httptest.NewRecorder()
	r := httptest.NewRequest("DELETE", "/categories/test-category-1/", nil)

	s := new(mocks.ICategoryService)

	s.On("GetBySlug", slug).Return(categoryDto, nil)
	s.On("Delete", categoryDto).Return(nil)

	h := NewCategoryHandler(s)

	err := h.Delete(w, r)

	assert.NoError(t, err)
	s.AssertExpectations(t)
}
