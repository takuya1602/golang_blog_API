package postgresql

import (
	"backend/app/domain/entity"
	"backend/app/domain/repository"
	"database/sql"
	"fmt"
)

type PostRepository struct {
	*sql.DB
}

func NewPostRepository(db *sql.DB) (postRepository repository.IPostRepository) {
	postRepository = &PostRepository{db}
	return
}

func (r *PostRepository) GetAll() (posts []entity.Post, err error) {
	rows, err := r.Query("select * from posts")
	if err != nil {
		return
	}
	for rows.Next() {
		var post entity.Post
		rows.Scan(&post.Id, &post.CategoryId, &post.SubCategoryId, &post.Title, &post.Slug, &post.EyeCatchingImg, &post.Content, &post.MetaDescription, &post.IsPublic, &post.CreatedAt, &post.UpdatedAt)
		posts = append(posts, post)
	}
	return
}

func (r *PostRepository) GetFilterCategory(categoryName string) (posts []entity.Post, err error) {
	var categoryId int
	err = r.QueryRow("select id from categories where slug = $1", categoryName).Scan(&categoryId)
	if err != nil {
		return
	}
	rows, err := r.Query("select * from posts where category_id = $1", categoryId)
	if err != nil {
		return
	}
	for rows.Next() {
		var post entity.Post
		rows.Scan(&post.Id, &post.CategoryId, &post.SubCategoryId, &post.Title, &post.Slug, &post.EyeCatchingImg, &post.Content, &post.MetaDescription, &post.IsPublic, &post.CreatedAt, &post.UpdatedAt)
		posts = append(posts, post)
	}
	return
}

func (r *PostRepository) GetFilterSubCategory(subCategoryName string) (posts []entity.Post, err error) {
	var subCategoryId int
	err = r.QueryRow("select id from sub_categories where slug = $1", subCategoryName).Scan(&subCategoryId)
	if err != nil {
		return
	}
	rows, err := r.Query("select * from posts where sub_category_id = $1", subCategoryId)
	if err != nil {
		return
	}
	for rows.Next() {
		var post entity.Post
		rows.Scan(&post.Id, &post.CategoryId, &post.SubCategoryId, &post.Title, &post.Slug, &post.EyeCatchingImg, &post.Content, &post.MetaDescription, &post.IsPublic, &post.CreatedAt, &post.UpdatedAt)
		posts = append(posts, post)
	}
	return
}

func (r *PostRepository) GetBySlug(slug string) (post entity.Post, err error) {
	err = r.QueryRow("select * from posts where slug = $1", slug).
		Scan(&post.Id, &post.CategoryId, &post.SubCategoryId, &post.Title, &post.Slug, &post.EyeCatchingImg, &post.Content, &post.MetaDescription, &post.IsPublic, &post.CreatedAt, &post.UpdatedAt)
	return
}

func (r *PostRepository) Create(post entity.Post) (err error) {
	_, err = r.Exec("insert into posts (category_id, sub_category_id, title, slug, eye_catching_img, content, meta_description, is_public) values ($1, $2, $3, $4, $5, $6, $7, $8)",
		post.CategoryId, post.SubCategoryId, post.Title, post.Slug, post.EyeCatchingImg, post.Content, post.MetaDescription, post.IsPublic)
	return
}

func (r *PostRepository) Update(post entity.Post) (err error) {
	_, err = r.Exec("update posts set category_id = $2, sub_category_id = $3, title = $4, slug = $5, eye_catching_img = $6, content = $7, meta_description = $8, is_public = $9 where id = $1",
		post.Id, post.CategoryId, post.SubCategoryId, post.Title, post.Slug, post.EyeCatchingImg, post.Content, post.MetaDescription, post.IsPublic)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func (r *PostRepository) Delete(post entity.Post) (err error) {
	_, err = r.Exec("delete from posts where id = $1", post.Id)
	return
}
