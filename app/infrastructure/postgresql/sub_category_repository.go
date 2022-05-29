package postgresql

import (
	"backend/app/domain/entity"
	"backend/app/domain/repository"
	"database/sql"
	"log"
)

type SubCategoryRepository struct {
	*sql.DB
}

func NewSubcategoryRepository(db *sql.DB) (subCategoryRepository repository.ISubCategoryRepository) {
	subCategoryRepository = &SubCategoryRepository{db}
	return
}

func (r *SubCategoryRepository) GetSubCategories(queryParams map[string][]string) (subCategories []entity.SubCategory, err error) {
	var rows *sql.Rows
	// join 2 tables (sub_categories, categories)
	// on sub_categories.parent_category_id = categories.id
	if categorySlugs, ok := queryParams["category-name"]; ok { // given category-name as query-params
		categorySlug := categorySlugs[0]
		rows, err = r.Query(`
			select 
			sub_categories.id as id, sub_categories.name, sub_categories.slug,
			categories.id as parent_category_id, categories.name as parent_category_name, categories.slug as parent_category_slug 
			from sub_categories 
			inner join categories 
			on sub_categories.parent_category_id = categories.id
			where categories.slug = $1
	`, categorySlug)
	} else { // no query-params; return all sub-categories
		rows, err = r.Query(`
			select 
			sub_categories.id as id, sub_categories.name, sub_categories.slug, 
			categories.id as parent_category_id, categories.name as parent_category_name, categories.slug as parent_category_slug 
			from sub_categories 
			inner join categories
			on sub_categories.parent_category_id = categories.id
	`)
	}
	if err != nil {
		return
	}
	for rows.Next() {
		var subCategory entity.SubCategory
		rows.Scan(&subCategory.Id, &subCategory.Name, &subCategory.Slug, &subCategory.ParentCategoryId, &subCategory.ParentCategoryName, &subCategory.ParentCategorySlug)
		subCategories = append(subCategories, subCategory)
	}
	return
}

func (r *SubCategoryRepository) GetSubCategoryBySlug(slug string) (subCategory entity.SubCategory, err error) {
	err = r.QueryRow(`
		select
		sub_categories.id as id, sub_categories.name, sub_categories.slug,
		categories.name as parent_category_name, categories.slug as parent_category_slug 
		from sub_categories
		inner join categories
		on sub_categories.parent_category_id = categories.id
		where sub_categories.slug = $1
	`, slug).
		Scan(&subCategory.Id, &subCategory.Name, &subCategory.Slug, &subCategory.ParentCategoryId, &subCategory.ParentCategoryName, &subCategory.ParentCategorySlug)
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

func (r *SubCategoryRepository) GetIdFromParentCategoryName(name string) (id int) {
	err := r.QueryRow("select id from categories where name = $1", name).Scan(&id)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func (r *SubCategoryRepository) GetNameFromParentCategoryId(id int) (name string) {
	err := r.QueryRow("select name from categories where id = $1", id).Scan(&name)
	if err != nil {
		log.Fatal(err)
	}
	return
}
