package repository

import "backend/app/domain/entity"

type IUserRepository interface {
	GetAll() ([]entity.User, error)
	ValidateUser(entity.Credentials) (entity.User, error)
	Create(entity.User) error
	Update(entity.User) error
	Delete(entity.User) error
}
