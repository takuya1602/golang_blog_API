package di

import (
	"backend/app/domain/service"
	"backend/app/handler"
	"database/sql"

	"backend/app/infrastructure/postgresql/repository"
)

func InitCategory(db *sql.DB) handler.ICategoryHandler {
	r := repository.NewCategoryRepository(db)
	s := service.NewCategoryService(r)
	return handler.NewCategoryHandler(s)
}

func InitSubCategory(db *sql.DB) handler.ISubCategoryHandler {
	r := repository.NewSubcategoryRepository(db)
	s := service.NewSubCategoryService(r)
	return handler.NewSubCategoryHandler(s)
}

func InitPost(db *sql.DB) handler.IPostHandler {
	r := repository.NewPostRepository(db)
	s := service.NewPostService(r)
	return handler.NewPostHandler(s)
}
