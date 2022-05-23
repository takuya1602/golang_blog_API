package di

import (
	"backend/app/domain/service"
	"backend/app/handler"
	"database/sql"

	"backend/app/infrastructure/postgresql"
)

func InitCategory(db *sql.DB) handler.ICategoryHandler {
	r := postgresql.NewCategoryRepository(db)
	s := service.NewCategoryService(r)
	return handler.NewCategoryHandler(s)
}

func InitSubCategory(db *sql.DB) handler.ISubCategoryHandler {
	r := postgresql.NewSubcategoryRepository(db)
	s := service.NewSubCategoryService(r)
	return handler.NewSubCategoryHandler(s)
}

func InitPost(db *sql.DB) handler.IPostHandler {
	r := postgresql.NewPostRepository(db)
	s := service.NewPostService(r)
	return handler.NewPostHandler(s)
}

func InitUser(db *sql.DB) handler.IUserHandler {
	r := postgresql.NewUserRepository(db)
	s := service.NewUserService(r)
	return handler.NewUserHandler(s)
}
