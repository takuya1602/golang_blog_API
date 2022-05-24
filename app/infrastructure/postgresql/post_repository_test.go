package postgresql

import (
	"backend/app/domain/entity"
	"reflect"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestPostRepositoryGetAll(t *testing.T) {
	postCreatedAt, _ := time.Parse("2006-01-02 15:04:05.999999-07", "2006-01-02 15:04:05.999999-07")
	postUpdatedAt, _ := time.Parse("2006-01-02 15:04:05.999999-07", "2006-01-02 15:04:05.999999-07")

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "category_id", "sub_category_id", "title", "slug", "eye_catching_img", "content", "meta_description", "is_public", "created_at", "updated_at"}).
		AddRow(1, 1, 1, "testPost1", "test-post-1", "test_post_1.png", "This is 1st post", "This is 1st post", false, postCreatedAt, postUpdatedAt).
		AddRow(2, 2, 2, "testPost2", "test-post-2", "test_post_2.png", "This is 2nd post", "This is 2nd post", false, postCreatedAt, postUpdatedAt)

	mock.ExpectQuery(regexp.QuoteMeta("select * from posts")).
		WillReturnRows(rows)

	r := NewPostRepository(db)

	posts, err := r.GetAll()
	if err != nil {
		t.Fatal(err)
	}

	expectedPosts := []entity.Post{
		{
			Id:              1,
			CategoryId:      1,
			SubCategoryId:   1,
			Title:           "testPost1",
			Slug:            "test-post-1",
			EyeCatchingImg:  "test_post_1.png",
			Content:         "This is 1st post",
			MetaDescription: "This is 1st post",
			IsPublic:        false,
			CreatedAt:       postCreatedAt,
			UpdatedAt:       postUpdatedAt,
		},
		{
			Id:              2,
			CategoryId:      2,
			SubCategoryId:   2,
			Title:           "testPost2",
			Slug:            "test-post-2",
			EyeCatchingImg:  "test_post_2.png",
			Content:         "This is 2nd post",
			MetaDescription: "This is 2nd post",
			IsPublic:        false,
			CreatedAt:       postCreatedAt,
			UpdatedAt:       postUpdatedAt,
		},
	}

	if !(reflect.DeepEqual(posts, expectedPosts)) {
		t.Fatalf("Wrong content, was expecting %v, but got %v\n", expectedPosts, posts)
	}
}

func TestPostRepositoryGetFilterCategory(t *testing.T) {
	postCreatedAt, _ := time.Parse("2006-01-02 15:04:05.999999-07", "2006-01-02 15:04:05.999999-07")
	postUpdatedAt, _ := time.Parse("2006-01-02 15:04:05.999999-07", "2006-01-02 15:04:05.999999-07")

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery(regexp.QuoteMeta("select id from categories where slug = $1")).
		WithArgs("test-category-1").
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	rows := sqlmock.NewRows([]string{"id", "category_id", "sub_category_id", "title", "slug", "eye_catching_img", "content", "meta_description", "is_public", "created_at", "updated_at"}).
		AddRow(1, 1, 1, "testPost1", "test-post-1", "test_post_1.png", "This is 1st post", "This is 1st post", false, postCreatedAt, postUpdatedAt).
		AddRow(2, 1, 2, "testPost2", "test-post-2", "test_post_2.png", "This is 2nd post", "This is 2nd post", false, postCreatedAt, postUpdatedAt)

	mock.ExpectQuery(regexp.QuoteMeta("select * from posts where category_id = $1")).
		WithArgs(1).
		WillReturnRows(rows)

	r := NewPostRepository(db)

	posts, err := r.GetFilterCategory("test-category-1")
	if err != nil {
		t.Fatal(err)
	}

	expectedPosts := []entity.Post{
		{
			Id:              1,
			CategoryId:      1,
			SubCategoryId:   1,
			Title:           "testPost1",
			Slug:            "test-post-1",
			EyeCatchingImg:  "test_post_1.png",
			Content:         "This is 1st post",
			MetaDescription: "This is 1st post",
			IsPublic:        false,
			CreatedAt:       postCreatedAt,
			UpdatedAt:       postUpdatedAt,
		},
		{
			Id:              2,
			CategoryId:      1,
			SubCategoryId:   2,
			Title:           "testPost2",
			Slug:            "test-post-2",
			EyeCatchingImg:  "test_post_2.png",
			Content:         "This is 2nd post",
			MetaDescription: "This is 2nd post",
			IsPublic:        false,
			CreatedAt:       postCreatedAt,
			UpdatedAt:       postUpdatedAt,
		},
	}

	if !(reflect.DeepEqual(posts, expectedPosts)) {
		t.Fatalf("Wrong content, was expecting %v, but got %v\n", expectedPosts, posts)
	}
}

func TestPostRepositoryGetFilterSubCategory(t *testing.T) {
	postCreatedAt, _ := time.Parse("2006-01-02 15:04:05.999999-07", "2006-01-02 15:04:05.999999-07")
	postUpdatedAt, _ := time.Parse("2006-01-02 15:04:05.999999-07", "2006-01-02 15:04:05.999999-07")

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery(regexp.QuoteMeta("select id from sub_categories where slug = $1")).
		WithArgs("test-sub-category-1").
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	rows := sqlmock.NewRows([]string{"id", "category_id", "sub_category_id", "title", "slug", "eye_catching_img", "content", "meta_description", "is_public", "created_at", "updated_at"}).
		AddRow(1, 1, 1, "testPost1", "test-post-1", "test_post_1.png", "This is 1st post", "This is 1st post", false, postCreatedAt, postUpdatedAt).
		AddRow(2, 2, 1, "testPost2", "test-post-2", "test_post_2.png", "This is 2nd post", "This is 2nd post", false, postCreatedAt, postUpdatedAt)

	mock.ExpectQuery(regexp.QuoteMeta("select * from posts where sub_category_id = $1")).
		WithArgs(1).
		WillReturnRows(rows)

	r := NewPostRepository(db)

	posts, err := r.GetFilterSubCategory("test-sub-category-1")
	if err != nil {
		t.Fatal(err)
	}

	expectedPosts := []entity.Post{
		{
			Id:              1,
			CategoryId:      1,
			SubCategoryId:   1,
			Title:           "testPost1",
			Slug:            "test-post-1",
			EyeCatchingImg:  "test_post_1.png",
			Content:         "This is 1st post",
			MetaDescription: "This is 1st post",
			IsPublic:        false,
			CreatedAt:       postCreatedAt,
			UpdatedAt:       postUpdatedAt,
		},
		{
			Id:              2,
			CategoryId:      2,
			SubCategoryId:   1,
			Title:           "testPost2",
			Slug:            "test-post-2",
			EyeCatchingImg:  "test_post_2.png",
			Content:         "This is 2nd post",
			MetaDescription: "This is 2nd post",
			IsPublic:        false,
			CreatedAt:       postCreatedAt,
			UpdatedAt:       postUpdatedAt,
		},
	}

	if !(reflect.DeepEqual(posts, expectedPosts)) {
		t.Fatalf("Wrong content, was expecting %v, but got %v\n", expectedPosts, posts)
	}
}

func TestPostRepositoryGetBySlug(t *testing.T) {
	postCreatedAt, _ := time.Parse("2006-01-02 15:04:05.999999-07", "2006-01-02 15:04:05.999999-07")
	postUpdatedAt, _ := time.Parse("2006-01-02 15:04:05.999999-07", "2006-01-02 15:04:05.999999-07")

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "category_id", "sub_category_id", "title", "slug", "eye_catching_img", "content", "meta_description", "is_public", "created_at", "updated_at"}).
		AddRow(1, 1, 1, "testPost1", "test-post-1", "test_post_1.png", "This is 1st post", "This is 1st post", false, postCreatedAt, postUpdatedAt)

	mock.ExpectQuery(regexp.QuoteMeta("select * from posts where slug = $1")).
		WithArgs("test-post-1").
		WillReturnRows(rows)

	r := NewPostRepository(db)

	post, err := r.GetBySlug("test-post-1")
	if err != nil {
		t.Fatal(err)
	}

	expectedPost := entity.Post{
		Id:              1,
		CategoryId:      1,
		SubCategoryId:   1,
		Title:           "testPost1",
		Slug:            "test-post-1",
		EyeCatchingImg:  "test_post_1.png",
		Content:         "This is 1st post",
		MetaDescription: "This is 1st post",
		IsPublic:        false,
		CreatedAt:       postCreatedAt,
		UpdatedAt:       postUpdatedAt,
	}

	if !(reflect.DeepEqual(post, expectedPost)) {
		t.Fatalf("Wrong content, was expecting %v, but got %v\n", expectedPost, post)
	}
}

func TestPostRepositoryCreate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectExec(regexp.QuoteMeta("insert into posts (category_id, sub_category_id, title, slug, eye_catching_img, content, meta_description, is_public) values ($1, $2, $3, $4, $5, $6, $7, $8)")).
		WithArgs(1, 1, "testPost1", "test-post-1", "test_post_1.png", "This is 1st post", "This is 1st post", false).
		WillReturnResult(sqlmock.NewResult(1, 11))

	r := NewPostRepository(db)

	post := entity.Post{
		CategoryId:      1,
		SubCategoryId:   1,
		Title:           "testPost1",
		Slug:            "test-post-1",
		EyeCatchingImg:  "test_post_1.png",
		Content:         "This is 1st post",
		MetaDescription: "This is 1st post",
		IsPublic:        false,
	}

	if err := r.Create(post); err != nil {
		t.Fatal(err)
	}
}

func TestPostRepositoryUpdate(t *testing.T) {
	postCreatedAt, _ := time.Parse("2006-01-02 15:04:05.999999-07", "2006-01-02 15:04:05.999999-07")
	postUpdatedAt, _ := time.Parse("2006-01-02 15:04:05.999999-07", "2006-01-02 15:04:05.999999-07")
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectExec(regexp.QuoteMeta("update posts set category_id = $2, sub_category_id = $3, title = $4, slug = $5, eye_catching_img = $6, content = $7, meta_description = $8, is_public = $9 where id = $1")).
		WithArgs(1, 1, 1, "testPost1", "test-post-1", "test_post_1.png", "This is 1st post", "This is 1st post", false).
		WillReturnResult(sqlmock.NewResult(1, 8))

	r := NewPostRepository(db)

	post := entity.Post{
		Id:              1,
		CategoryId:      1,
		SubCategoryId:   1,
		Title:           "testPost1",
		Slug:            "test-post-1",
		EyeCatchingImg:  "test_post_1.png",
		Content:         "This is 1st post",
		MetaDescription: "This is 1st post",
		IsPublic:        false,
		CreatedAt:       postCreatedAt,
		UpdatedAt:       postUpdatedAt,
	}

	if err := r.Update(post); err != nil {
		t.Fatal(err)
	}
}

func TestPostRepositoryDelete(t *testing.T) {
	postCreatedAt, _ := time.Parse("2006-01-02 15:04:05.999999-07", "2006-01-02 15:04:05.999999-07")
	postUpdatedAt, _ := time.Parse("2006-01-02 15:04:05.999999-07", "2006-01-02 15:04:05.999999-07")
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectExec(regexp.QuoteMeta("delete from posts where id = $1")).
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(1, 11))

	r := NewPostRepository(db)

	post := entity.Post{
		Id:              1,
		CategoryId:      1,
		SubCategoryId:   1,
		Title:           "testPost1",
		Slug:            "test-post-1",
		EyeCatchingImg:  "test_post_1.png",
		Content:         "This is 1st post",
		MetaDescription: "This is 1st post",
		IsPublic:        false,
		CreatedAt:       postCreatedAt,
		UpdatedAt:       postUpdatedAt,
	}

	if err := r.Delete(post); err != nil {
		t.Fatal(err)
	}
}
