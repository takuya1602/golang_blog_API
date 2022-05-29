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

func (r *PostRepository) GetPosts(queryParams map[string][]string) (posts []entity.Post, err error) {
	var rows *sql.Rows
	// join 3 tables (posts, categories, sub_categories)
	// on posts.sub_category_id = sub_categories.id
	// on sub_categories.parent_category_id = categories.id
	if categorySlugs, ok := queryParams["category-name"]; ok { // given category-name as query-params
		categorySlug := categorySlugs[0]
		rows, err = r.Query(`
			select
			posts.id as id, title, posts.slug, eye_catching_img, content, meta_description, is_public, created_at, updated_at,
			categories.id as category_id, categories.name as category_name, categories.slug as category_slug,
			sub_categories.id as sub_category_id, sub_categories.name as sub_category_name, sub_categories.slug as sub_category_slug
			from (
			(posts inner join sub_categories on posts.sub_category_id = sub_categories.id)
			inner join categories on sub_categories.parent_category_id = categories.id)
			where categories.slug = $1
	`, categorySlug)
	} else if subCategorySlugs, ok := queryParams["sub-category-name"]; ok { // given sub-category-name as query-params
		subCategorySlug := subCategorySlugs[0]
		rows, err = r.Query(`
			select
			posts.id as id, title, posts.slug, eye_catching_img, content, meta_description, is_public, created_at, updated_at,
			categories.id as category_id, categories.name as category_name, categories.slug as category_slug,
			sub_categories.id as sub_category_id, sub_categories.name as sub_category_name, sub_categories.slug as sub_category_slug
			from (
			(posts inner join sub_categories on posts.sub_category_id = sub_categories.id)
			inner join categories on sub_categories.parent_category_id = categories.id)
			where sub_categories.slug = $1
		`, subCategorySlug)
	} else { // no query-params; return all posts
		rows, err = r.Query(`
			select
			posts.id as id, title, posts.slug, eye_catching_img, content, meta_description, is_public, created_at, updated_at,
			categories.id as category_id, categories.name as category_name, categories.slug as category_slug,
			sub_categories.id as sub_category_id, sub_categories.name as sub_category_name, sub_categories.slug as sub_category_slug
			from (
			(posts inner join sub_categories on posts.sub_category_id = sub_categories.id)
			inner join categories on sub_categories.parent_category_id = categories.id)
		`)
	}
	if err != nil {
		return
	}
	for rows.Next() {
		var post entity.Post
		rows.Scan(
			&post.Id,
			&post.Title,
			&post.Slug,
			&post.EyeCatchingImg,
			&post.Content,
			&post.MetaDescription,
			&post.IsPublic,
			&post.CreatedAt,
			&post.UpdatedAt,
			&post.CategoryId,
			&post.CategoryName,
			&post.CategorySlug,
			&post.SubCategoryId,
			&post.SubCategoryName,
			&post.SubCategorySlug,
		)
		posts = append(posts, post)
	}
	return
}

func (r *PostRepository) GetPostBySlug(slug string) (post entity.Post, err error) {
	err = r.QueryRow(`
		select
		posts.id as id, title, posts.slug, eye_catching_img, content, meta_description, is_public, created_at, updated_at,
		categories.id as category_id, categories.name as category_name, categories.slug as category_slug,
		sub_categories.id as sub_category_id, sub_categories.name as sub_category_name, sub_categories.slug as sub_category_slug
		from (
		(posts inner join sub_categories on posts.sub_category_id = sub_categories.id)
		inner join categories on sub_categories.parent_category_id = categories.id)
		where posts.slug = $1
	`, slug).
		Scan(
			&post.Id,
			&post.Title,
			&post.Slug,
			&post.EyeCatchingImg,
			&post.Content,
			&post.MetaDescription,
			&post.IsPublic,
			&post.CreatedAt,
			&post.UpdatedAt,
			&post.CategoryId,
			&post.CategoryName,
			&post.CategorySlug,
			&post.SubCategoryId,
			&post.SubCategoryName,
			&post.SubCategorySlug,
		)
	return
}

func (r *PostRepository) Create(post entity.Post) (err error) {
	_, err = r.Exec("insert into posts (title, slug, eye_catching_img, content, meta_description, is_public, sub_category_id) values ($1, $2, $3, $4, $5, $6, $7)",
		post.Title, post.Slug, post.EyeCatchingImg, post.Content, post.MetaDescription, post.IsPublic, post.SubCategoryId)
	return
}

func (r *PostRepository) Update(post entity.Post) (err error) {
	_, err = r.Exec("update posts set title = $2, slug = $3, eye_catching_img = $4, content = $5, meta_description = $6, is_public = $7, sub_category_id = $8 where id = $1",
		post.Id, post.Title, post.Slug, post.EyeCatchingImg, post.Content, post.MetaDescription, post.IsPublic, post.SubCategoryId)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func (r *PostRepository) Delete(post entity.Post) (err error) {
	_, err = r.Exec("delete from posts where id = $1", post.Id)
	return
}
