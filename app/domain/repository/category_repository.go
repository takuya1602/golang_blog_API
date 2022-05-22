package repository

import "backend/app/infrastructure/postgresql/entity"

type ICategoryRepositry interface {
	GetAll() (categories []entity.Category, err error)
	GetBySlug(slug string) (category entity.Category, err error)
	Create(category entity.Category) (err error)
	Update(entity.Category) (err error)
	Delete(entity.Category) (err error)
}
