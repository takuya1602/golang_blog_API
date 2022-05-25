package postgresql

import (
	"backend/app/domain/entity"
	"backend/app/domain/repository"
	"database/sql"
)

type CategoryRepository struct {
	*sql.DB
}

func NewCategoryRepository(db *sql.DB) (categoryRepository repository.ICategoryRepository) {
	categoryRepository = &CategoryRepository{db}
	return
}

func (r *CategoryRepository) GetAll() (categories []entity.Category, err error) {
	rows, err := r.Query("select * from categories")
	if err != nil {
		return
	}
	for rows.Next() {
		var category entity.Category
		rows.Scan(&category.Id, &category.Name, &category.Slug)
		categories = append(categories, category)
	}
	return
}

func (r *CategoryRepository) GetBySlug(slug string) (category entity.Category, err error) {
	err = r.QueryRow("select id, name, slug from categories where slug = $1", slug).
		Scan(&category.Id, &category.Name, &category.Slug)
	return
}

func (r *CategoryRepository) Create(category entity.Category) (err error) {
	_, err = r.Exec("insert into categories (name, slug) values ($1, $2)", category.Name, category.Slug)
	return
}

func (r *CategoryRepository) Update(category entity.Category) (err error) {
	_, err = r.Exec("update categories set name = $2, slug = $3 where id = $1", category.Id, category.Name, category.Slug)
	return
}

func (r *CategoryRepository) Delete(category entity.Category) (err error) {
	_, err = r.Exec("delete from categories where id = $1", category.Id)
	return
}
