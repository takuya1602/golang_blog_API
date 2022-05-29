package service

import (
	"backend/app/common/dto"
	"backend/app/domain/entity"
	"backend/app/domain/repository"
)

type ISubCategoryService interface {
	GetSubCategories(map[string][]string) ([]dto.SubCategoryModel, error)
	GetSubCategoryBySlug(string) (dto.SubCategoryModel, error)
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
	subCategoryDto = dto.SubCategoryModel{
		Id:                 subCategory.Id,
		Name:               subCategory.Name,
		Slug:               subCategory.Slug,
		ParentCategoryId:   subCategory.ParentCategoryId,
		ParentCategoryName: subCategory.ParentCategoryName,
		ParentCategorySlug: subCategory.ParentCategorySlug,
	}
	return
}

func (s *SubCategoryService) convertToDtosFromEntities(subCategories []entity.SubCategory) (subCategoryDtos []dto.SubCategoryModel) {
	for _, subCategory := range subCategories {
		subCategoryDto := s.convertToDtoFromEntity(subCategory)
		subCategoryDtos = append(subCategoryDtos, subCategoryDto)
	}
	return
}

func (s *SubCategoryService) convertToEntityFromDto(subCategoryDto dto.SubCategoryModel) (subCategory entity.SubCategory) {
	subCategory = entity.SubCategory{
		Id:                 subCategoryDto.Id,
		Name:               subCategoryDto.Name,
		Slug:               subCategoryDto.Slug,
		ParentCategoryId:   subCategoryDto.ParentCategoryId,
		ParentCategoryName: subCategoryDto.ParentCategoryName,
		ParentCategorySlug: subCategoryDto.ParentCategorySlug,
	}
	return
}

func (s *SubCategoryService) convertToEntitiesFromDtos(subCategoryDtos []dto.SubCategoryModel) (subCategories []entity.SubCategory) {
	for _, subCategoryDto := range subCategoryDtos {
		subCategory := s.convertToEntityFromDto(subCategoryDto)
		subCategories = append(subCategories, subCategory)
	}
	return
}

func (s *SubCategoryService) GetSubCategories(queryParams map[string][]string) (subCategoryDtos []dto.SubCategoryModel, err error) {
	subCategories, err := s.ISubCategoryRepository.GetSubCategories(queryParams)
	if err != nil {
		return
	}
	subCategoryDtos = s.convertToDtosFromEntities(subCategories)
	return
}

func (s *SubCategoryService) GetSubCategoryBySlug(slug string) (subCategoryDto dto.SubCategoryModel, err error) {
	subCategory, err := s.ISubCategoryRepository.GetSubCategoryBySlug(slug)
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
