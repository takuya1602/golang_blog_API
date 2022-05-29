package handler

import (
	"backend/app/common/dto"
	mocks "backend/mocks/service"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubCategoryHandler_Get(t *testing.T) {
	t.Run(
		"with query param: category-name",
		func(t *testing.T) {
			subCategoryDtos := []dto.SubCategoryModel{
				{
					Id:                 1,
					Name:               "testSubCategory1",
					Slug:               "test-sub-category-1",
					ParentCategoryName: "testCategory1",
				},
				{
					Id:                 2,
					Name:               "testSubCategory2",
					Slug:               "test-sub-category-2",
					ParentCategoryName: "testCategory1",
				},
			}
			categorySlug := "test-category-1"

			s := new(mocks.ISubCategoryService)

			s.On("GetWithQuery", categorySlug).Return(subCategoryDtos, nil)

			h := NewSubCategoryHandler(s)

			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/sub-categories?category-name=test-category-1", nil)

			err := h.Get(w, r)

			assert.NoError(t, err)
			s.AssertExpectations(t)
		},
	)
	t.Run(
		"without query param",
		func(t *testing.T) {
			subCategoryDtos := []dto.SubCategoryModel{
				{
					Id:                 1,
					Name:               "testSubCategory1",
					Slug:               "test-sub-category-1",
					ParentCategoryName: "testCategory1",
				},
				{
					Id:                 2,
					Name:               "testSubCategory2",
					Slug:               "test-sub-category-2",
					ParentCategoryName: "testCategory1",
				},
			}
			s := new(mocks.ISubCategoryService)

			s.On("GetAll").Return(subCategoryDtos, nil)

			h := NewSubCategoryHandler(s)

			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/sub-categories/", nil)

			err := h.Get(w, r)

			assert.NoError(t, err)
			s.AssertExpectations(t)
		},
	)
}

func TestSubCategoryHandler_Create(t *testing.T) {
	subCategoryDto := dto.NewSubCategoryModel(1, "testSubCategory1", "test-sub-category-1", "testCategory1")
	json := strings.NewReader(`{
		"id": 1,
		"name": "testSubCategory1",
		"slug": "test-sub-category-1",
		"parent_category": "testCategory1"
	}`)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/sub-categories/", json)

	s := new(mocks.ISubCategoryService)

	s.On("Create", subCategoryDto).Return(nil)

	h := NewSubCategoryHandler(s)

	err := h.Create(w, r)

	assert.NoError(t, err)
	s.AssertExpectations(t)
}

func TestSubCategoryHandler_Update(t *testing.T) {
	subCategoryDto := dto.NewSubCategoryModel(1, "testSubCategory1", "test-sub-category-1", "testCategory1")
	json := strings.NewReader(`{
		"id": 1,
		"name": "testSubCategory1",
		"slug": "test-sub-category-1",
		"parent_category": "testCategory1"
	}`)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("PUT", "/sub-categories/test-sub-category-1/", json)

	s := new(mocks.ISubCategoryService)

	s.On("GetBySlug", subCategoryDto.Slug).Return(subCategoryDto, nil)
	s.On("Update", subCategoryDto).Return(nil)

	h := NewSubCategoryHandler(s)

	err := h.Update(w, r)

	assert.NoError(t, err)
	s.AssertExpectations(t)
}

func TestSubCategoryHandler_Delete(t *testing.T) {
	subCategoryDto := dto.NewSubCategoryModel(1, "testSubCategory1", "test-sub-category-1", "testCategory1")

	w := httptest.NewRecorder()
	r := httptest.NewRequest("DELETE", "/sub-categories/test-sub-category-1/", nil)

	s := new(mocks.ISubCategoryService)

	s.On("GetBySlug", subCategoryDto.Slug).Return(subCategoryDto, nil)
	s.On("Delete", subCategoryDto).Return(nil)

	h := NewSubCategoryHandler(s)

	err := h.Delete(w, r)

	assert.NoError(t, err)
	s.AssertExpectations(t)
}
