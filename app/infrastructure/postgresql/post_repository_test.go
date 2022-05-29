package postgresql

import (
	"backend/app/domain/entity"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func AssertPosts(t *testing.T, ret []entity.Post, posts []entity.Post) {
	for i, r := range ret {
		assert.Equal(t, r.Id, posts[i].Id)
		assert.Equal(t, r.CategoryId, posts[i].CategoryId)
		assert.Equal(t, r.SubCategoryId, posts[i].SubCategoryId)
		assert.Equal(t, r.Title, posts[i].Title)
		assert.Equal(t, r.Slug, posts[i].Slug)
		assert.Equal(t, r.EyeCatchingImg, posts[i].EyeCatchingImg)
		assert.Equal(t, r.Content, posts[i].Content)
		assert.Equal(t, r.MetaDescription, posts[i].MetaDescription)
		assert.Equal(t, r.IsPublic, posts[i].IsPublic)
		assert.Equal(t, r.CreatedAt, posts[i].CreatedAt)
		assert.Equal(t, r.UpdatedAt, posts[i].UpdatedAt)
	}
}

func TestPostRepository_GetPosts(t *testing.T) {
	postCreatedAt, _ := time.Parse("2006-01-02 15:04:05.999999-07", "2006-01-02 15:04:05.999999-07")
	postUpdatedAt, _ := time.Parse("2006-01-02 15:04:05.999999-07", "2006-01-02 15:04:05.999999-07")

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	posts := []entity.Post{
		{
			Id:              1,
			Title:           "testPost1",
			Slug:            "test-post-1",
			EyeCatchingImg:  "test_post_1.png",
			Content:         "This is 1st post",
			MetaDescription: "This is 1st post",
			IsPublic:        false,
			CreatedAt:       postCreatedAt,
			UpdatedAt:       postUpdatedAt,
			CategoryId:      1,
			CategoryName:    "testCategory1",
			CategorySlug:    "test-category-1",
			SubCategoryId:   1,
			SubCategoryName: "testSubCategory1",
			SubCategorySlug: "test-sub-category-1",
		},
		{
			Id:              2,
			Title:           "testPost2",
			Slug:            "test-post-2",
			EyeCatchingImg:  "test_post_2.png",
			Content:         "This is 2nd post",
			MetaDescription: "This is 2nd post",
			IsPublic:        false,
			CreatedAt:       postCreatedAt,
			UpdatedAt:       postUpdatedAt,
			CategoryId:      1,
			CategoryName:    "testCategory1",
			CategorySlug:    "test-category-1",
			SubCategoryId:   1,
			SubCategoryName: "testSubCategory1",
			SubCategorySlug: "test-sub-category-1",
		},
	}

	fields := []string{
		"id",
		"title",
		"slug",
		"eye_catching_img",
		"content",
		"meta_description",
		"is_public",
		"created_at",
		"updated_at",
		"category_id",
		"category_name",
		"category_slug",
		"sub_category_id",
		"sub_category_name",
		"sub_category_slug",
	}

	rows := sqlmock.NewRows(fields).
		AddRow(
			posts[0].Id,
			posts[0].Title,
			posts[0].Slug,
			posts[0].EyeCatchingImg,
			posts[0].Content,
			posts[0].MetaDescription,
			posts[0].IsPublic,
			posts[0].CreatedAt,
			posts[0].UpdatedAt,
			posts[0].CategoryId,
			posts[0].CategoryName,
			posts[0].CategorySlug,
			posts[0].SubCategoryId,
			posts[0].SubCategoryName,
			posts[0].SubCategorySlug,
		).
		AddRow(
			posts[1].Id,
			posts[1].Title,
			posts[1].Slug,
			posts[1].EyeCatchingImg,
			posts[1].Content,
			posts[1].MetaDescription,
			posts[1].IsPublic,
			posts[1].CreatedAt,
			posts[1].UpdatedAt,
			posts[1].CategoryId,
			posts[1].CategoryName,
			posts[1].CategorySlug,
			posts[1].SubCategoryId,
			posts[1].SubCategoryName,
			posts[1].SubCategorySlug,
		)

	t.Run(
		"with query-params: category-name",
		func(t *testing.T) {
			mock.ExpectQuery(regexp.QuoteMeta(`
				select
				posts.id as id, title, posts.slug, eye_catching_img, content, meta_description, is_public, created_at, updated_at,
				categories.id as category_id, categories.name as category_name, categories.slug as category_slug,
				sub_categories.id as sub_category_id, sub_categories.name as sub_category_name, sub_categories.slug as sub_category_slug
				from (
				(posts inner join sub_categories on posts.sub_category_id = sub_categories.id)
				inner join categories on sub_categories.parent_category_id = categories.id)
				where categories.slug = $1
			`)).WithArgs(posts[0].CategorySlug).WillReturnRows(rows)

			r := NewPostRepository(db)

			queryParams := map[string][]string{
				"category-name": {
					posts[0].CategorySlug,
				},
			}

			ret, err := r.GetPosts(queryParams)

			assert.NoError(t, err)
			AssertPosts(t, ret, posts)
		},
	)

	t.Run(
		"with query-params: sub-category-name",
		func(t *testing.T) {
			mock.ExpectQuery(regexp.QuoteMeta(`
				select
				posts.id as id, title, posts.slug, eye_catching_img, content, meta_description, is_public, created_at, updated_at,
				categories.id as category_id, categories.name as category_name, categories.slug as category_slug,
				sub_categories.id as sub_category_id, sub_categories.name as sub_category_name, sub_categories.slug as sub_category_slug
				from (
				(posts inner join sub_categories on posts.sub_category_id = sub_categories.id)
				inner join categories on sub_categories.parent_category_id = categories.id)
				where sub_categories.slug = $1
			`)).WithArgs(posts[0].SubCategorySlug).WillReturnRows(rows)

			r := NewPostRepository(db)

			queryParams := map[string][]string{
				"sub-category-name": {
					posts[0].SubCategorySlug,
				},
			}

			ret, err := r.GetPosts(queryParams)

			assert.NoError(t, err)
			AssertPosts(t, ret, posts)
		},
	)

	t.Run(
		"without query-params",
		func(t *testing.T) {
			mock.ExpectQuery(regexp.QuoteMeta(`
				select
				posts.id as id, title, posts.slug, eye_catching_img, content, meta_description, is_public, created_at, updated_at,
				categories.id as category_id, categories.name as category_name, categories.slug as category_slug,
				sub_categories.id as sub_category_id, sub_categories.name as sub_category_name, sub_categories.slug as sub_category_slug
				from (
				(posts inner join sub_categories on posts.sub_category_id = sub_categories.id)
				inner join categories on sub_categories.parent_category_id = categories.id)
			`)).WillReturnRows(rows)

			r := NewPostRepository(db)

			var queryParams map[string][]string

			ret, err := r.GetPosts(queryParams)

			assert.NoError(t, err)
			AssertPosts(t, ret, posts)
		},
	)
}

func TestPostRepository_CRUD(t *testing.T) {
	postCreatedAt, _ := time.Parse("2006-01-02 15:04:05.999999-07", "2006-01-02 15:04:05.999999-07")
	postUpdatedAt, _ := time.Parse("2006-01-02 15:04:05.999999-07", "2006-01-02 15:04:05.999999-07")

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	post := entity.Post{
		Id:              1,
		Title:           "testPost1",
		Slug:            "test-post-1",
		EyeCatchingImg:  "test_post_1.png",
		Content:         "This is 1st post",
		MetaDescription: "This is 1st post",
		IsPublic:        false,
		CreatedAt:       postCreatedAt,
		UpdatedAt:       postUpdatedAt,
		CategoryId:      1,
		CategoryName:    "testCategory1",
		CategorySlug:    "test-category-1",
		SubCategoryId:   1,
		SubCategoryName: "testSubCategory1",
		SubCategorySlug: "test-sub-category-1",
	}

	fields := []string{
		"id",
		"title",
		"slug",
		"eye_catching_img",
		"content",
		"meta_description",
		"is_public",
		"created_at",
		"updated_at",
		"category_id",
		"category_name",
		"category_slug",
		"sub_category_id",
		"sub_category_name",
		"sub_category_slug",
	}

	rows := sqlmock.NewRows(fields).
		AddRow(
			post.Id,
			post.Title,
			post.Slug,
			post.EyeCatchingImg,
			post.Content,
			post.MetaDescription,
			post.IsPublic,
			post.CreatedAt,
			post.UpdatedAt,
			post.CategoryId,
			post.CategoryName,
			post.CategorySlug,
			post.SubCategoryId,
			post.SubCategoryName,
			post.SubCategorySlug,
		)

	t.Run(
		"GetPostBySlug",
		func(t *testing.T) {
			mock.ExpectQuery(regexp.QuoteMeta(`
				select
				posts.id as id, title, posts.slug, eye_catching_img, content, meta_description, is_public, created_at, updated_at,
				categories.id as category_id, categories.name as category_name, categories.slug as category_slug,
				sub_categories.id as sub_category_id, sub_categories.name as sub_category_name, sub_categories.slug as sub_category_slug
				from (
				(posts inner join sub_categories on posts.sub_category_id = sub_categories.id)
				inner join categories on sub_categories.parent_category_id = categories.id)
				where posts.slug = $1
			`)).WithArgs(post.Slug).WillReturnRows(rows)

			r := NewPostRepository(db)

			ret, err := r.GetPostBySlug(post.Slug)

			assert.NoError(t, err)
			assert.Equal(t, ret.Id, post.Id)
			assert.Equal(t, ret.CategoryId, post.CategoryId)
			assert.Equal(t, ret.SubCategoryId, post.SubCategoryId)
			assert.Equal(t, ret.Title, post.Title)
			assert.Equal(t, ret.Slug, post.Slug)
			assert.Equal(t, ret.EyeCatchingImg, post.EyeCatchingImg)
			assert.Equal(t, ret.Content, post.Content)
			assert.Equal(t, ret.MetaDescription, post.MetaDescription)
			assert.Equal(t, ret.IsPublic, post.IsPublic)
			assert.Equal(t, ret.CreatedAt, post.CreatedAt)
			assert.Equal(t, ret.UpdatedAt, post.UpdatedAt)
		},
	)

	t.Run(
		"Create",
		func(t *testing.T) {
			mock.ExpectExec(regexp.QuoteMeta("insert into posts (title, slug, eye_catching_img, content, meta_description, is_public, sub_category_id) values ($1, $2, $3, $4, $5, $6, $7)")).
				WithArgs(post.Title, post.Slug, post.EyeCatchingImg, post.Content, post.MetaDescription, post.IsPublic, post.SubCategoryId).
				WillReturnResult(sqlmock.NewResult(1, 8))

			r := NewPostRepository(db)

			err := r.Create(post)

			assert.NoError(t, err)
		},
	)

	t.Run(
		"Update",
		func(t *testing.T) {
			mock.ExpectExec(regexp.QuoteMeta("update posts set title = $2, slug = $3, eye_catching_img = $4, content = $5, meta_description = $6, is_public = $7, sub_category_id = $8 where id = $1")).
				WithArgs(post.Id, post.Title, post.Slug, post.EyeCatchingImg, post.Content, post.MetaDescription, post.IsPublic, post.SubCategoryId).
				WillReturnResult(sqlmock.NewResult(1, 7))

			r := NewPostRepository(db)

			err := r.Update(post)

			assert.NoError(t, err)
		},
	)

	t.Run(
		"Delete",
		func(t *testing.T) {
			mock.ExpectExec(regexp.QuoteMeta("delete from posts where id = $1")).
				WithArgs(post.Id).
				WillReturnResult(sqlmock.NewResult(1, 8))

			r := NewPostRepository(db)

			err := r.Delete(post)

			assert.NoError(t, err)
		},
	)
}
