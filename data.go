package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func retrieveCategories(db *sql.DB) (categories []Category, err error) {
	rows, err := db.Query("select * from categories")
	if err != nil {
		return
	}
	for rows.Next() {
		category := Category{}
		rows.Scan(&category.Id, &category.CategoryName, &category.Slug)
		categories = append(categories, category)
	}
	return
}

func retrieveCategory(db *sql.DB, slug string) (category Category, err error) {
	category = Category{}
	err = db.QueryRow("select * from categories where slug = $1", slug).Scan(&category.Id, &category.CategoryName, &category.Slug)
	return
}

func retrieveSubCategories(db *sql.DB, queryParams map[string][]string) (subCategories []SubCategory, err error) {
	var rows *sql.Rows
	if categorySlugs, ok := queryParams["category-name"]; ok {
		categorySlug := categorySlugs[0] // cannot receive multiple parameters
		var categoryId int
		err = db.QueryRow("select id from categories where slug = $1", categorySlug).Scan(&categoryId)
		if err != nil {
			return
		}
		rows, err = db.Query("select * from sub_categories where parent_category_id = $1", categoryId)
		if err != nil {
			return
		}
	} else {
		rows, err = db.Query("select * from sub_categories") // The case that there is no category-name in query params; return all sub-categories
		if err != nil {
			return
		}
	}
	for rows.Next() {
		subCategory := SubCategory{}
		rows.Scan(&subCategory.Id, &subCategory.CategoryName, &subCategory.Slug, &subCategory.ParentCategoryId)
		subCategories = append(subCategories, subCategory)
	}
	return
}

func retrieveSubCategory(db *sql.DB, slug string) (subCategory SubCategory, err error) {
	subCategory = SubCategory{}
	err = db.QueryRow("select * from sub_categories where slug = $1", slug).Scan(&subCategory.Id, &subCategory.CategoryName, &subCategory.Slug, &subCategory.ParentCategoryId)
	if err != nil {
		return
	}
	return
}

func retrievePosts(db *sql.DB, queryParams map[string][]string) (posts []Post, err error) {
	var rows *sql.Rows
	if categorySlugs, ok := queryParams["category-name"]; ok { // filtering posts by category-name
		categorySlug := categorySlugs[0] // cannot recieve multiple parameters
		var categoryId int
		err = db.QueryRow("select id from categories where slug = $1", categorySlug).Scan(&categoryId)
		if err != nil {
			return
		}
		rows, err = db.Query("select * from posts where category_id = $1", categoryId)
		if err != nil {
			return
		}
	} else if subCategorySlugs, ok := queryParams["sub-category-name"]; ok { // filtering posts by sub-category-name
		subCategorySlug := subCategorySlugs[0] // cannot recieve multiple parameters
		var subCategoryId int
		err = db.QueryRow("select id from sub_categories where slug = $1", subCategorySlug).Scan(&subCategoryId)
		if err != nil {
			return
		}
		rows, err = db.Query("select * from posts where sub_category_id = $1", subCategoryId)
		if err != nil {
			return
		}
	} else { // The case that there is no category-name and sub-category-name in query parameters; return all posts
		rows, err = db.Query("select * from posts")
		if err != nil {
			return
		}
	}
	for rows.Next() {
		post := Post{}
		rows.Scan(&post.Id, &post.CategoryId, &post.SubCategoryId, &post.Title, &post.Slug, &post.EyeCatchingImg, &post.Content, &post.MetaDescription, &post.IsPublic, &post.CreatedAt, &post.UpdatedAt)
		posts = append(posts, post)
	}
	return
}

func retrievePost(db *sql.DB, slug string) (post Post, err error) {
	post = Post{}
	err = db.QueryRow("select id, category_id, sub_category_id, title, slug, eye_catching_img, content, meta_description, is_public, created_at, updated_at from posts where slug = $1", slug).Scan(&post.Id, &post.CategoryId, &post.SubCategoryId, &post.Title, &post.Slug, &post.EyeCatchingImg, &post.Content, &post.MetaDescription, &post.IsPublic, &post.CreatedAt, &post.UpdatedAt)
	if err != nil {
		return
	}
	return
}

func (category *Category) create(db *sql.DB) (err error) {
	err = db.QueryRow("insert into categories (category_name, slug) values ($1, $2) returning id", category.CategoryName, category.Slug).Scan(&category.Id)
	return
}

func (category *Category) update(db *sql.DB) (err error) {
	_, err = db.Exec("update categories set category_name = $2, slug = $3 where id = $1",
		category.Id, category.CategoryName, category.Slug)
	return
}

func (category *Category) delete(db *sql.DB) (err error) {
	_, err = db.Exec("delete from categories where id = $1", category.Id)
	return
}

func (subCategory *SubCategory) create(db *sql.DB) (err error) {
	err = db.QueryRow("insert into sub_categories (category_name, slug, parent_category_id) values ($1, $2, $3) returning id",
		subCategory.CategoryName, subCategory.Slug, subCategory.ParentCategoryId).Scan(&subCategory.Id)
	return
}

func (subCategory *SubCategory) update(db *sql.DB) (err error) {
	_, err = db.Exec("update sub_categories set category_name = $2, slug = $3, parent_category_id = $4 where id = $1",
		subCategory.Id, subCategory.CategoryName, subCategory.Slug, subCategory.ParentCategoryId)
	return
}

func (subCategory *SubCategory) delete(db *sql.DB) (err error) {
	_, err = db.Exec("delete from sub_categories where id = $1", subCategory.Id)
	return
}

func (post *Post) create(db *sql.DB) (err error) {
	err = db.QueryRow("insert into posts (category_id, sub_category_id, title, slug, eye_catching_img, content, meta_description, is_public) values ($1, $2, $3, $4, $5, $6, $7, $8) returning id, created_at, updated_at",
		post.CategoryId, post.SubCategoryId, post.Title, post.Slug, post.EyeCatchingImg, post.Content, post.MetaDescription, post.IsPublic).Scan(&post.Id, &post.CreatedAt, &post.UpdatedAt)
	return
}

func (post *Post) update(db *sql.DB) (err error) {
	_, err = db.Exec("update posts set category_id = $2, sub_category_id = $3, title = $4, slug = $5, eye_catching_img = $6, content = $7, meta_description = $8, is_public = $9 where id = $1",
		post.Id, post.CategoryId, post.SubCategoryId, post.Title, post.Slug, post.EyeCatchingImg, post.Content, post.MetaDescription, post.IsPublic)
	return
}

func (post *Post) delete(db *sql.DB) (err error) {
	_, err = db.Exec("delete from posts where id = $1", post.Id)
	return
}
