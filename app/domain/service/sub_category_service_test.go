package service

import (
	"backend/app/common/dto"
	"backend/app/domain/entity"
	mocks "backend/mocks/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func assertSubCategories(t *testing.T, ret []dto.SubCategoryModel, subCategories []entity.SubCategory) {
	for i, r := range ret {
		assert.Equal(t, r.Id, subCategories[i].Id)
		assert.Equal(t, r.Name, subCategories[i].Name)
		assert.Equal(t, r.Slug, subCategories[i].Slug)
		assert.Equal(t, r.ParentCategoryId, subCategories[i].ParentCategoryId)
		assert.Equal(t, r.ParentCategoryName, subCategories[i].ParentCategoryName)
		assert.Equal(t, r.ParentCategorySlug, subCategories[i].ParentCategorySlug)
	}
}

func TestSubCategoryService_GetSubCategories(t *testing.T) {
	subCategories := []entity.SubCategory{
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
		"with query params: category-name",
		func(t *testing.T) {
			r := new(mocks.ISubCategoryRepository)

			queryParams := map[string][]string{
				"category-name": {
					subCategories[0].Slug,
				},
			}

			r.On("GetSubCategories", queryParams).Return(subCategories, nil)

			s := NewSubCategoryService(r)

			ret, err := s.GetSubCategories(queryParams)

			assert.NoError(t, err)
			assertSubCategories(t, ret, subCategories)
			r.AssertExpectations(t)
		},
	)

	t.Run(
		"without query params",
		func(t *testing.T) {
			r := new(mocks.ISubCategoryRepository)

			var queryParams map[string][]string

			r.On("GetSubCategories", queryParams).Return(subCategories, nil)

			s := NewSubCategoryService(r)

			ret, err := s.GetSubCategories(queryParams)

			assert.NoError(t, err)
			assertSubCategories(t, ret, subCategories)
			r.AssertExpectations(t)
		},
	)
}

func TestSubCategoryService_CRUD(t *testing.T) {
	subCategory := entity.SubCategory{
		Id:                 1,
		Name:               "testSubCategory1",
		Slug:               "test-sub-category-1",
		ParentCategoryId:   1,
		ParentCategoryName: "testCategory1",
		ParentCategorySlug: "test-category-1",
	}

	subCategoryDto := dto.SubCategoryModel{
		Id:                 1,
		Name:               "testSubCategory1",
		Slug:               "test-sub-category-1",
		ParentCategoryId:   1,
		ParentCategoryName: "testCategory1",
		ParentCategorySlug: "test-category-1",
	}

	t.Run(
		"GetSubCategoryBySlug",
		func(t *testing.T) {
			r := new(mocks.ISubCategoryRepository)

			r.On("GetSubCategoryBySlug", subCategory.Slug).Return(subCategory, nil)
			s := NewSubCategoryService(r)

			ret, err := s.GetSubCategoryBySlug(subCategory.Slug)

			assert.NoError(t, err)
			assert.Equal(t, ret.Id, subCategory.Id)
			assert.Equal(t, ret.Name, subCategory.Name)
			assert.Equal(t, ret.Slug, subCategory.Slug)
			assert.Equal(t, ret.ParentCategoryId, subCategory.ParentCategoryId)
			assert.Equal(t, ret.ParentCategoryName, subCategory.ParentCategoryName)
			assert.Equal(t, ret.ParentCategorySlug, subCategory.ParentCategorySlug)
			r.AssertExpectations(t)
		},
	)

	t.Run(
		"Create",
		func(t *testing.T) {
			r := new(mocks.ISubCategoryRepository)

			r.On("Create", subCategory).Return(nil)

			s := NewSubCategoryService(r)

			err := s.Create(subCategoryDto)

			assert.NoError(t, err)
			r.AssertExpectations(t)
		},
	)

	t.Run(
		"Update",
		func(t *testing.T) {
			r := new(mocks.ISubCategoryRepository)

			r.On("Update", subCategory).Return(nil)

			s := NewSubCategoryService(r)

			err := s.Update(subCategoryDto)

			assert.NoError(t, err)
			r.AssertExpectations(t)
		},
	)

	t.Run(
		"Delete",
		func(t *testing.T) {
			r := new(mocks.ISubCategoryRepository)

			r.On("Delete", subCategory).Return(nil)

			s := NewSubCategoryService(r)

			err := s.Delete(subCategoryDto)

			assert.NoError(t, err)
			r.AssertExpectations(t)
		},
	)
}
