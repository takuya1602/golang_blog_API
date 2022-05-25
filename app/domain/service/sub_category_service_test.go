package service

import (
	"backend/app/common/dto"
	"backend/app/domain/entity"
	mocks "backend/mocks/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubCategoryService_GetAll(t *testing.T) {
	subCategories := []entity.SubCategory{
		{
			Id:               1,
			Name:             "testSubCategory1",
			Slug:             "test-sub-category-1",
			ParentCategoryId: 1,
		},
		{
			Id:               2,
			Name:             "testSubCategory2",
			Slug:             "test-sub-category-2",
			ParentCategoryId: 1,
		},
	}

	r := new(mocks.ISubCategoryRepository)

	r.On("GetAll").Return(subCategories, nil)

	s := NewSubCategoryService(r)

	ret, err := s.GetAll()

	assert.NoError(t, err)
	for i, r := range ret {
		assert.Equal(t, r.Id, subCategories[i].Id)
		assert.Equal(t, r.Name, subCategories[i].Name)
		assert.Equal(t, r.Slug, subCategories[i].Slug)
		assert.Equal(t, r.ParentCategoryId, subCategories[i].ParentCategoryId)
	}
	r.AssertExpectations(t)
}

func TestSubCategoryService_GetWithQuery(t *testing.T) {
	subCategories := []entity.SubCategory{
		{
			Id:               1,
			Name:             "testSubCategory1",
			Slug:             "test-sub-category-1",
			ParentCategoryId: 1,
		},
		{
			Id:               2,
			Name:             "testSubCategory2",
			Slug:             "test-sub-category-2",
			ParentCategoryId: 1,
		},
	}
	slug := "test-category-1"

	r := new(mocks.ISubCategoryRepository)

	r.On("GetFilterParentCategory", slug).Return(subCategories, nil)

	s := NewSubCategoryService(r)

	ret, err := s.GetWithQuery(slug)

	assert.NoError(t, err)
	for i, r := range ret {
		assert.Equal(t, r.Id, subCategories[i].Id)
		assert.Equal(t, r.Name, subCategories[i].Name)
		assert.Equal(t, r.Slug, subCategories[i].Slug)
		assert.Equal(t, r.ParentCategoryId, subCategories[i].ParentCategoryId)
	}
	r.AssertExpectations(t)
}

func TestSubCategoryService_GetBySlug(t *testing.T) {
	subCategory := entity.NewSubCategory(1, "testSubCategory1", "test-sub-category-1", 1)
	slug := "test-sub-category-1"

	r := new(mocks.ISubCategoryRepository)

	r.On("GetBySlug", slug).Return(subCategory, nil)

	s := NewSubCategoryService(r)

	ret, err := s.GetBySlug(slug)

	assert.NoError(t, err)
	assert.Equal(t, ret.Id, subCategory.Id)
	assert.Equal(t, ret.Name, subCategory.Name)
	assert.Equal(t, ret.Slug, subCategory.Slug)
	assert.Equal(t, ret.ParentCategoryId, subCategory.ParentCategoryId)
	r.AssertExpectations(t)
}

func TestSubCategoryService_Create(t *testing.T) {
	subCategoryDto := dto.NewSubCategoryModel(1, "testSubCategory1", "test-sub-category-1", 1)
	subCategory := entity.NewSubCategory(1, "testSubCategory1", "test-sub-category-1", 1)

	r := new(mocks.ISubCategoryRepository)

	r.On("Create", subCategory).Return(nil)

	s := NewSubCategoryService(r)

	err := s.Create(subCategoryDto)

	assert.NoError(t, err)
	r.AssertExpectations(t)
}

func TestSubCategoryService_Update(t *testing.T) {
	subCategoryDto := dto.NewSubCategoryModel(1, "testSubCategory1", "test-sub-category-1", 1)
	subCategory := entity.NewSubCategory(1, "testSubCategory1", "test-sub-category-1", 1)

	r := new(mocks.ISubCategoryRepository)

	r.On("Update", subCategory).Return(nil)

	s := NewSubCategoryService(r)

	err := s.Update(subCategoryDto)

	assert.NoError(t, err)
	r.AssertExpectations(t)
}

func TestSubCategoryService_Delete(t *testing.T) {
	subCategoryDto := dto.NewSubCategoryModel(1, "testSubCategory1", "test-sub-category-1", 1)
	subCategory := entity.NewSubCategory(1, "testSubCategory1", "test-sub-category-1", 1)

	r := new(mocks.ISubCategoryRepository)

	r.On("Delete", subCategory).Return(nil)

	s := NewSubCategoryService(r)

	err := s.Delete(subCategoryDto)

	assert.NoError(t, err)
	r.AssertExpectations(t)
}
