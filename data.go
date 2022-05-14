package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", "user=gwp password=gwp dbname=go_blog sslmode=disable")
	if err != nil {
		panic(err)
	}
}

func retrieveParentCategory(parent_category_id int) (parentCategory ParentCategory, err error) {
	parentCategory = ParentCategory{}
	err = Db.QueryRow("select id, category_name, slug from parent_categories where id = $1", parent_category_id).Scan(&parentCategory.Id, &parentCategory.CategoryName, &parentCategory.Slug)
	return
}

func retrieveParentCategories() (parentCategories []ParentCategory, err error) {
	rows, err := Db.Query("select * from parent_categories")
	if err != nil {
		return
	}
	for rows.Next() {
		parentCategory := ParentCategory{}
		rows.Scan(&parentCategory.Id, &parentCategory.CategoryName, &parentCategory.Slug)
		parentCategories = append(parentCategories, parentCategory)
	}
	return
}

func retrieveCategories() (categories []Category, err error) {
	rows, err := Db.Query("select * from categories")
	if err != nil {
		return
	}
	for rows.Next() {
		category := Category{}
		rows.Scan(&category.Id, &category.CategoryName, &category.Slug, &category.ParentCategoryId)
		categories = append(categories, category)
	}
	return
}

func retrieveCategory(category_id int) (category Category, err error) {
	category = Category{}
	err = Db.QueryRow("select id, category_name, slug, parent_category_id from categories where id = $1", category_id).Scan(&category.Id, &category.CategoryName, &category.Slug, &category.ParentCategoryId)
	if err != nil {
		return
	}
	return
}

func retrieveCategoryPosts(slug string) (posts []Post, err error) {
	var categoryId int
	err = Db.QueryRow("select id from categories where slug = $1", slug).Scan(&categoryId)
	if err != nil {
		return
	}
	rows, err := Db.Query("select * from posts where category_id = $1", categoryId)
	if err != nil {
		return
	}
	for rows.Next() {
		post := Post{}
		rows.Scan(&post.Id, &post.CategoryId, &post.Title, &post.Content, &post.Slug, &post.EyeCatchingImg)
		posts = append(posts, post)
	}
	return
}

func retrievePosts() (posts []Post, err error) {
	rows, err := Db.Query("select * from posts")
	if err != nil {
		return
	}
	for rows.Next() {
		post := Post{}
		rows.Scan(&post.Id, &post.CategoryId, &post.Title, &post.Content, &post.Slug, &post.EyeCatchingImg)
		posts = append(posts, post)
	}
	return
}

func retrievePost(slug string) (post Post, err error) {
	post = Post{}
	err = Db.QueryRow("select id, slug, title, content, eye_catching_img, category_id from posts where slug = $1", slug).Scan(&post.Id, &post.Slug, &post.Title, &post.Content, &post.EyeCatchingImg, &post.CategoryId)
	if err != nil {
		return
	}
	return
}

func (parentCategory *ParentCategory) create() (err error) {
	err = Db.QueryRow("insert into parent_categories (category_name, slug) values ($1, $2) returning id", parentCategory.CategoryName, parentCategory.Slug).Scan(&parentCategory.Id)
	return
}

func (parentCategory *ParentCategory) update() (err error) {
	_, err = Db.Exec("update parent_categories set category_name = $2, slug = $3 where id = $1",
		parentCategory.Id, parentCategory.CategoryName, parentCategory.Slug)
	return
}

func (parentCategory *ParentCategory) delete() (err error) {
	_, err = Db.Exec("delete from parent_categories where id = $1", parentCategory.Id)
	return
}

func (category *Category) create() (err error) {
	err = Db.QueryRow("insert into categories (category_name, slug, parent_category_id) values ($1, $2, $3) returning id",
		category.CategoryName, category.Slug, category.ParentCategoryId).Scan(&category.Id)
	return
}

func (category *Category) update() (err error) {
	_, err = Db.Exec("update categories set category_name = $2, slug = $3, parent_category_id = $4 where id = $1",
		category.Id, category.CategoryName, category.Slug, category.ParentCategoryId)
	return
}

func (category *Category) delete() (err error) {
	_, err = Db.Exec("delete from categories where id = $1", category.Id)
	return
}

func (post *Post) create() (err error) {
	err = Db.QueryRow("insert into posts (slug, category_id, title, content, eye_catching_img) values ($1, $2, $3, $4, $5) returning id",
		post.Slug, post.CategoryId, post.Title, post.Content, post.EyeCatchingImg).Scan(&post.Id)
	return
}

func (post *Post) update() (err error) {
	_, err = Db.Exec("update posts set slug = $2, category_id = $3, title = $4, content = $5, eye_catching_img = $6 where id = $1",
		post.Slug, post.CategoryId, post.Id, post.Title, post.Content, post.EyeCatchingImg)
	return
}

func (post *Post) delete() (err error) {
	_, err = Db.Exec("delete from posts where id = $1", post.Id)
	return
}
