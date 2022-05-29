package repository

import "backend/app/domain/entity"

type IPostRepository interface {
	GetAll() ([]entity.Post, error)
	GetFilterCategory(string) ([]entity.Post, error)
	GetFilterSubCategory(string) ([]entity.Post, error)
	GetBySlug(string) (entity.Post, error)
	Create(entity.Post) error
	Update(entity.Post) error
	Delete(entity.Post) error
	GetIdFromCategoryName(string) int
	GetNameFromCategoryId(int) string
	GetIdFromSubCategoryName(string) int
	GetNameFromSubCategoryId(int) string
}
