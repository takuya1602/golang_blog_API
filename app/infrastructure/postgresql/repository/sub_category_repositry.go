package repository

import (
	"backend/app/domain/repository"
	"backend/app/infrastructure/postgresql/entity"
	"database/sql"
)

type SubCategoryRepositry struct {
	*sql.DB
}

func NewSubcategoryRepository(db *sql.DB) (subCategoryRepositry repository.ISubCategoryRepository) {
	subCategoryRepositry = &SubCategoryRepositry{db}
	return
}

func (r *SubCategoryRepositry) GetAll() (subCategories []entity.SubCategory, err error) {
	rows, err := r.Query("select id, name, slug, parent_category_id from sub_categories")
	if err != nil {
		return
	}
	for rows.Next() {
		var subCategory entity.SubCategory
		rows.Scan(&subCategory.Id, &subCategory.Name, &subCategory.Slug, &subCategory.ParentCategoryId)
		subCategories = append(subCategories, subCategory)
	}
	return
}

func (r *SubCategoryRepositry) GetBySlug(slug string) (subCategory entity.SubCategory, err error) {
	err = r.QueryRow("select id, name, slug, parent_category_id from sub_categories where slug = $1", slug).
		Scan(&subCategory.Id, &subCategory.Name, &subCategory.Slug, &subCategory.ParentCategoryId)
	if err != nil {
		return
	}
	return
}

func (r *SubCategoryRepositry) Create(subCategory entity.SubCategory) (err error) {
	_, err = r.Exec("insert into sub_categories (name, slug, parent_category_id) values ($1, $2, $3)", subCategory.Name, subCategory.Slug, subCategory.ParentCategoryId)
	if err != nil {
		return
	}
	return
}

func (r *SubCategoryRepositry) Update(subCategory entity.SubCategory) (err error) {
	_, err = r.Exec("update sub_categories set name = $2, slug = $3, parent_category_id = $4 where id = $1",
		subCategory.Id, subCategory.Name, subCategory.Slug, subCategory.ParentCategoryId)
	if err != nil {
		return
	}
	return
}

func (r *SubCategoryRepositry) Delete(subCategory entity.SubCategory) (err error) {
	_, err = r.Exec("delete from sub_categories where id = $1", subCategory.Id)
	if err != nil {
		return
	}
	return
}
