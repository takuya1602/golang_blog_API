package repository

import "backend/app/domain/entity"

type ISubCategoryRepository interface {
	GetAll() ([]entity.SubCategory, error)
	GetBySlug(string) (entity.SubCategory, error)
	Create(entity.SubCategory) error
	Update(entity.SubCategory) error
	Delete(entity.SubCategory) error
}
