package handler

import (
	"backend/app/common/dto"
	mocks "backend/mocks/service"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubCategoryHandler_GetSubCategories(t *testing.T) {
	subCategoryDtos := []dto.SubCategoryModel{
		{
			Id:                 1,
			Name:               "testSubCategory1",
			Slug:               "test-sub-category-1",
			ParentCategoryId:   1,
			ParentCategoryName: "testCategory1",
			ParentCategorySlug: "test-category-1",
		},
		{
			Id:                 2,
			Name:               "testSubCategory2",
			Slug:               "test-sub-category-2",
			ParentCategoryId:   1,
			ParentCategoryName: "testCategory1",
			ParentCategorySlug: "test-category-1",
		},
	}

	t.Run(
		"with query param: category-name",
		func(t *testing.T) {
			s := new(mocks.ISubCategoryService)

			queryParams := map[string][]string{
				"category-name": {
					subCategoryDtos[0].ParentCategorySlug,
				},
			}

			s.On("GetSubCategories", queryParams).Return(subCategoryDtos, nil)

			h := NewSubCategoryHandler(s)

			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/sub-categories?category-name=test-category-1", nil)

			err := h.GetSubCategories(w, r)

			assert.NoError(t, err)
			s.AssertExpectations(t)
		},
	)
	t.Run(
		"without query param",
		func(t *testing.T) {
			s := new(mocks.ISubCategoryService)

			queryParams := map[string][]string{}

			s.On("GetSubCategories", queryParams).Return(subCategoryDtos, nil)

			h := NewSubCategoryHandler(s)

			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/sub-categories/", nil)

			err := h.GetSubCategories(w, r)

			assert.NoError(t, err)
			s.AssertExpectations(t)
		},
	)
}

func TestSubCategoryHandler_CRUD(t *testing.T) {
	// Create, Update, Delete
	// allow (parent_category_name, parent_category_slug)
	// empty, because "sub_categories" table in postgresSQL database dont't have these fields.
	subCategoryDto := dto.SubCategoryModel{
		Id:               1,
		Name:             "testSubCategory1",
		Slug:             "test-sub-category-1",
		ParentCategoryId: 1,
	}

	json := strings.NewReader(`{
		"id": 1,
		"name": "testSubCategory1",
		"slug": "test-sub-category-1",
		"parent_category_id": 1
	}`)

	t.Run(
		"Create",
		func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest("POST", "/sub-categories/", json)

			s := new(mocks.ISubCategoryService)

			s.On("Create", subCategoryDto).Return(nil)

			h := NewSubCategoryHandler(s)

			err := h.Create(w, r)

			assert.NoError(t, err)
			s.AssertExpectations(t)
		},
	)

	t.Run(
		"Update",
		func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest("PUT", "/sub-categories/test-sub-category-1/", json)

			s := new(mocks.ISubCategoryService)

			s.On("GetSubCategoryBySlug", subCategoryDto.Slug).Return(subCategoryDto, nil)
			s.On("Update", subCategoryDto).Return(nil)

			h := NewSubCategoryHandler(s)

			err := h.Update(w, r)

			assert.NoError(t, err)
			s.AssertExpectations(t)
		},
	)

	t.Run(
		"Delete",
		func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest("DELETE", "/sub-categories/test-sub-category-1/", nil)

			s := new(mocks.ISubCategoryService)

			s.On("GetSubCategoryBySlug", subCategoryDto.Slug).Return(subCategoryDto, nil)
			s.On("Delete", subCategoryDto).Return(nil)

			h := NewSubCategoryHandler(s)

			err := h.Delete(w, r)

			assert.NoError(t, err)
			s.AssertExpectations(t)
		},
	)
}
