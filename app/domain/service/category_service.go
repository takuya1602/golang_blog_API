package service

import (
	"backend/app/common/dto"
	"backend/app/domain/repository"
	"backend/app/infrastructure/postgresql/entity"
)

type ICategoryService interface {
	GetAll() (categories []dto.CategoryModel, err error)
	GetBySlug(slug string) (category dto.CategoryModel, err error)
	Create(categoryDto dto.CategoryModel) (err error)
	Update(dto.CategoryModel) (err error)
	Delete(dto.CategoryModel) (err error)
}

type CategoryService struct {
	repository.ICategoryRepositry
}

func NewCategoryService(repo repository.ICategoryRepositry) (categoryService ICategoryService) {
	categoryService = &CategoryService{repo}
	return
}

func (s *CategoryService) convertToDtoFromEntity(category entity.Category) (categoryDto dto.CategoryModel) {
	categoryDto = dto.NewCategoryModel(category.Id, category.Name, category.Slug)
	return
}

func (s *CategoryService) convertToDtosFromEntities(categories []entity.Category) (categoryDtos []dto.CategoryModel) {
	for _, category := range categories {
		categoryDto := dto.NewCategoryModel(category.Id, category.Name, category.Slug)
		categoryDtos = append(categoryDtos, categoryDto)
	}
	return
}

func (s *CategoryService) convertToEntityFromDto(categoryDto dto.CategoryModel) (category entity.Category) {
	category = entity.NewCategory(categoryDto.Id, categoryDto.Name, categoryDto.Slug)
	return
}

func (s *CategoryService) convertToEntitiesFromDtos(categoryDtos []dto.CategoryModel) (categories []entity.Category) {
	for _, categoryDto := range categoryDtos {
		category := entity.NewCategory(categoryDto.Id, categoryDto.Name, categoryDto.Slug)
		categories = append(categories, category)
	}
	return
}

func (s *CategoryService) GetBySlug(slug string) (categoryDto dto.CategoryModel, err error) {
	category, err := s.ICategoryRepositry.GetBySlug(slug)
	if err != nil {
		return
	}
	categoryDto = s.convertToDtoFromEntity(category)
	return
}

func (s *CategoryService) GetAll() (categoryDtos []dto.CategoryModel, err error) {
	categories, err := s.ICategoryRepositry.GetAll()
	if err != nil {
		return
	}
	categoryDtos = s.convertToDtosFromEntities(categories)
	return
}

func (s *CategoryService) Create(categoryDto dto.CategoryModel) (err error) {
	category := s.convertToEntityFromDto(categoryDto)
	err = s.ICategoryRepositry.Create(category)
	return
}

func (s *CategoryService) Update(categoryDto dto.CategoryModel) (err error) {
	category := s.convertToEntityFromDto(categoryDto)
	err = s.ICategoryRepositry.Update(category)
	return
}

func (s *CategoryService) Delete(categoryDto dto.CategoryModel) (err error) {
	category := s.convertToEntityFromDto(categoryDto)
	err = s.ICategoryRepositry.Delete(category)
	return
}
