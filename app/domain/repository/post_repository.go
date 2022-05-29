package repository

import "backend/app/domain/entity"

type IPostRepository interface {
	GetPosts(map[string][]string) ([]entity.Post, error)
	GetPostBySlug(string) (entity.Post, error)
	Create(entity.Post) error
	Update(entity.Post) error
	Delete(entity.Post) error
	GetIdFromCategoryName(string) int
	GetNameFromCategoryId(int) string
	GetIdFromSubCategoryName(string) int
	GetNameFromSubCategoryId(int) string
}
