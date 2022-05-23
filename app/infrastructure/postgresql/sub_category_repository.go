package postgresql

import (
	"backend/app/domain/entity"
	"backend/app/domain/repository"
	"database/sql"
)

type SubCategoryRepository struct {
	*sql.DB
}

func NewSubcategoryRepository(db *sql.DB) (subCategoryRepository repository.ISubCategoryRepository) {
	subCategoryRepository = &SubCategoryRepository{db}
	return
}

func (r *SubCategoryRepository) GetAll() (subCategories []entity.SubCategory, err error) {
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

func (r *SubCategoryRepository) GetFilterParentCategory(parentCategoryId int) (subCategories []entity.SubCategory, er error) {
	rows, err := r.Query("select * from sub_categories where parent_category_id = $1", parentCategoryId)
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

func (r *SubCategoryRepository) GetBySlug(slug string) (subCategory entity.SubCategory, err error) {
	err = r.QueryRow("select id, name, slug, parent_category_id from sub_categories where slug = $1", slug).
		Scan(&subCategory.Id, &subCategory.Name, &subCategory.Slug, &subCategory.ParentCategoryId)
	if err != nil {
		return
	}
	return
}

func (r *SubCategoryRepository) Create(subCategory entity.SubCategory) (err error) {
	_, err = r.Exec("insert into sub_categories (name, slug, parent_category_id) values ($1, $2, $3)", subCategory.Name, subCategory.Slug, subCategory.ParentCategoryId)
	if err != nil {
		return
	}
	return
}

func (r *SubCategoryRepository) Update(subCategory entity.SubCategory) (err error) {
	_, err = r.Exec("update sub_categories set name = $2, slug = $3, parent_category_id = $4 where id = $1",
		subCategory.Id, subCategory.Name, subCategory.Slug, subCategory.ParentCategoryId)
	if err != nil {
		return
	}
	return
}

func (r *SubCategoryRepository) Delete(subCategory entity.SubCategory) (err error) {
	_, err = r.Exec("delete from sub_categories where id = $1", subCategory.Id)
	if err != nil {
		return
	}
	return
}

func (r *SubCategoryRepository) GetParentCategoryId(categoryName string) (parentCategoryId int, err error) {
	err = r.QueryRow("select id from categories where slug = $1", categoryName).
		Scan(&parentCategoryId)
	return
}
