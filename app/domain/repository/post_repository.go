package repository

import "backend/app/domain/entity"

type IPostRepository interface {
	GetAll() ([]entity.Post, error)
	GetBySlug(string) (entity.Post, error)
	Create(entity.Post) error
	Update(entity.Post) error
	Delete(entity.Post) error
}
