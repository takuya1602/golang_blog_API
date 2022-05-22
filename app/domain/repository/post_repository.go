package repository

import "backend/app/infrastructure/postgresql/entity"

type IPostRepository interface {
	GetAll() ([]entity.Post, error)
	GetBySlug(string) (entity.Post, error)
	Create(entity.Post) error
	Update(entity.Post) error
	Delete(entity.Post) error
}
