package service

import (
	"backend/app/common/dto"
	"backend/app/domain/entity"
	mocks "backend/mocks/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCategoryService_GetAll(t *testing.T) {
	categories := []entity.Category{
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

	r := new(mocks.ICategoryRepository)

	r.On("GetAll").Return(categories, nil)

	s := NewCategoryService(r)

	ret, err := s.GetAll()

	assert.NoError(t, err)
	for i, r := range ret {
		assert.Equal(t, r.Id, categories[i].Id)
		assert.Equal(t, r.Name, categories[i].Name)
		assert.Equal(t, r.Slug, categories[i].Slug)
	}
	r.AssertExpectations(t)
}

func TestCategoryService_GetBySlug(t *testing.T) {
	category := entity.Category{
		Id:   1,
		Name: "testCategory1",
		Slug: "test-category-1",
	}

	r := new(mocks.ICategoryRepository)

	r.On("GetBySlug", category.Slug).Return(category, nil)

	s := NewCategoryService(r)

	ret, err := s.GetBySlug(category.Slug)

	assert.NoError(t, err)
	assert.Equal(t, ret.Id, category.Id)
	assert.Equal(t, ret.Name, category.Name)
	assert.Equal(t, ret.Slug, category.Slug)
	r.AssertExpectations(t)
}

func TestCategoryService_Create(t *testing.T) {
	category := entity.Category{
		Name: "testCategory1",
		Slug: "test-category-1",
	}
	categoryDto := dto.CategoryModel{
		Name: "testCategory1",
		Slug: "test-category-1",
	}

	r := new(mocks.ICategoryRepository)

	r.On("Create", category).Return(nil)

	s := NewCategoryService(r)

	assert.NoError(t, s.Create(categoryDto))
	r.AssertExpectations(t)
}

func TestCategoryService_Update(t *testing.T) {
	category := entity.NewCategory(1, "testCategory1", "test-category-1")
	categoryDto := dto.NewCategoryModel(1, "testCategory1", "test-category-1")

	r := new(mocks.ICategoryRepository)

	r.On("Update", category).Return(nil)

	s := NewCategoryService(r)

	assert.NoError(t, s.Update(categoryDto))
	r.AssertExpectations(t)
}

func TestCategoryService_Delete(t *testing.T) {
	category := entity.NewCategory(1, "testCategory1", "test-category-1")
	categoryDto := dto.NewCategoryModel(1, "testCategory1", "test-category-1")

	r := new(mocks.ICategoryRepository)

	r.On("Delete", category).Return(nil)

	s := NewCategoryService(r)

	assert.NoError(t, s.Delete(categoryDto))
	r.AssertExpectations(t)
}
