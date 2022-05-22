package service

import (
	"backend/app/common/dto"
	"backend/app/domain/repository"
	"backend/app/infrastructure/postgresql/entity"
)

type ISubCategoryService interface {
	GetAll() ([]dto.SubCategoryModel, error)
	GetBySlug(string) (dto.SubCategoryModel, error)
	Create(dto.SubCategoryModel) error
	Update(dto.SubCategoryModel) error
	Delete(dto.SubCategoryModel) error
}

type SubCategoryService struct {
	repository.ISubCategoryRepository
}

func NewSubCategoryService(repo repository.ISubCategoryRepository) (subCategoryService ISubCategoryService) {
	subCategoryService = &SubCategoryService{repo}
	return
}

func (s *SubCategoryService) convertToDtoFromEntity(subCategory entity.SubCategory) (subCategoryDto dto.SubCategoryModel) {
	subCategoryDto = dto.NewSubCategoryModel(subCategory.Id, subCategory.Name, subCategory.Slug, subCategory.ParentCategoryId)
	return
}

func (s *SubCategoryService) convertToDtosFromEntities(subCategories []entity.SubCategory) (subCategoryDtos []dto.SubCategoryModel) {
	for _, subCategory := range subCategories {
		subCategoryDto := dto.NewSubCategoryModel(subCategory.Id, subCategory.Name, subCategory.Slug, subCategory.ParentCategoryId)
		subCategoryDtos = append(subCategoryDtos, subCategoryDto)
	}
	return
}

func (s *SubCategoryService) convertToEntityFromDto(subCategoryDto dto.SubCategoryModel) (subCategory entity.SubCategory) {
	subCategory = entity.NewSubCategory(subCategoryDto.Id, subCategoryDto.Name, subCategoryDto.Slug, subCategoryDto.ParentCategoryId)
	return
}

func (s *SubCategoryService) convertToEntitiesFromDtos(subCategoryDtos []dto.SubCategoryModel) (subCategories []entity.SubCategory) {
	for _, subCategoryDto := range subCategoryDtos {
		subCategory := entity.NewSubCategory(subCategoryDto.Id, subCategoryDto.Name, subCategoryDto.Slug, subCategoryDto.ParentCategoryId)
		subCategories = append(subCategories, subCategory)
	}
	return
}

func (s *SubCategoryService) GetAll() (subCategoryDtos []dto.SubCategoryModel, err error) {
	subCategories, err := s.ISubCategoryRepository.GetAll()
	if err != nil {
		return
	}
	subCategoryDtos = s.convertToDtosFromEntities(subCategories)
	return
}

func (s *SubCategoryService) GetBySlug(slug string) (subCategoryDto dto.SubCategoryModel, err error) {
	subCategory, err := s.ISubCategoryRepository.GetBySlug(slug)
	if err != nil {
		return
	}
	subCategoryDto = s.convertToDtoFromEntity(subCategory)
	return
}

func (s *SubCategoryService) Create(subCategoryDto dto.SubCategoryModel) (err error) {
	subCategory := s.convertToEntityFromDto(subCategoryDto)
	err = s.ISubCategoryRepository.Create(subCategory)
	return
}

func (s *SubCategoryService) Update(subCategoryDto dto.SubCategoryModel) (err error) {
	subCategory := s.convertToEntityFromDto(subCategoryDto)
	err = s.ISubCategoryRepository.Update(subCategory)
	return
}

func (s *SubCategoryService) Delete(subCategoryDto dto.SubCategoryModel) (err error) {
	subCategory := s.convertToEntityFromDto(subCategoryDto)
	err = s.ISubCategoryRepository.Delete(subCategory)
	return
}
