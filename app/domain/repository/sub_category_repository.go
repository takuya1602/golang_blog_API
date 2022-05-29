package repository

import "backend/app/domain/entity"

type ISubCategoryRepository interface {
	GetSubCategories(map[string][]string) ([]entity.SubCategory, error)
	GetSubCategoryBySlug(string) (entity.SubCategory, error)
	Create(entity.SubCategory) error
	Update(entity.SubCategory) error
	Delete(entity.SubCategory) error
}
