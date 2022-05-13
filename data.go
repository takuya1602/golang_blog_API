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

func retrieveCategory(category_id int) (category Category, err error) {
	var parentCategoryId int
	category = Category{}
	err = Db.QueryRow("select id, category_name, slug, parent_category_id from categories where id = $1", category_id).Scan(&category.Id, &category.CategoryName, &category.Slug, &parentCategoryId)
	parentCategory, err := retrieveParentCategory(parentCategoryId)
	if err != nil {
		return
	}
	category.ParentCategory = &parentCategory
	return
}

func retrievePost(id int) (post Post, err error) {
	var categoryId int
	post = Post{}
	err = Db.QueryRow("select id, slug, title, content, eye_catching_img, category_id from posts where id = $1", id).Scan(&post.Id, &post.Slug, &post.Title, &post.Content, &post.EyeCatchingImg, &categoryId)
	category, err := retrieveCategory(categoryId)
	if err != nil {
		return
	}
	post.Category = &category
	return
}

func (post *Post) create() (err error) {
	err = Db.QueryRow("insert into posts (slug, category_id, title, content, eye_catching_img) values ($1, $2, $3, $4, $5) returning id",
		post.Slug, post.Category.Id, post.Title, post.Content, post.EyeCatchingImg).Scan(&post.Id)
	return
}

func (post *Post) update() (err error) {
	_, err = Db.Exec("update posts set slug = $2, category_id = $3, title = $4, content = $5, eye_catching_img = $6 where id = $1",
		post.Slug, post.Category.Id, post.Id, post.Title, post.Content, post.EyeCatchingImg)
	return
}

func (post *Post) delete() (err error) {
	_, err = Db.Exec("delete from posts where id = $1", post.Id)
	return
}
