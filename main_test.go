package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"regexp"
	"strings"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestHandleGetCategory(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	e := Env{Db: db}

	row := sqlmock.NewRows([]string{"id", "category_name", "slug"}).
		AddRow(1, "testCategory1", "test-category-1").
		AddRow(2, "testCategory2", "test-category-2")

	mock.ExpectQuery(regexp.QuoteMeta("select * from categories")).
		WillReturnRows(row)

	mux := http.NewServeMux()
	mux.HandleFunc("/categories/", e.handleRequestCategory)

	writer := httptest.NewRecorder()

	request, err := http.NewRequest("GET", "/categories/", nil)
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Fatalf("Response code is %v\n", writer.Code)
	}

	categories := []Category{}
	json.Unmarshal(writer.Body.Bytes(), &categories)

	expectedCategories := []Category{
		{
			Id:           1,
			CategoryName: "testCategory1",
			Slug:         "test-category-1",
		},
		{
			Id:           2,
			CategoryName: "testCategory2",
			Slug:         "test-category-2",
		},
	}

	if !(reflect.DeepEqual(categories, expectedCategories)) {
		t.Fatalf("Wrong content, was expecting %v, but got %v\n", expectedCategories, categories)
	}
}

func TestHandlePostCategory(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	e := Env{Db: db}

	mock.ExpectQuery(regexp.QuoteMeta("insert into categories (category_name, slug) values ($1, $2) returning id")).
		WithArgs("testCategory1", "test-category-1").
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	mux := http.NewServeMux()
	mux.HandleFunc("/categories/", e.handleRequestCategory)

	writer := httptest.NewRecorder()

	json := strings.NewReader(`{
		"category_name": "testCategory1",
		"slug": "test-category-1"
	}`)

	request, err := http.NewRequest("POST", "/categories/", json)
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Fatalf("Response code is %v\n", writer.Code)
	}
}

func TestHandlePutCategory(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	e := Env{Db: db}

	row := sqlmock.NewRows([]string{"id", "category_name", "slug"}).
		AddRow(1, "testCategory1", "test-category-1")

	mock.ExpectQuery(regexp.QuoteMeta("select * from categories where slug = $1")).
		WithArgs("test-category-1").
		WillReturnRows(row)

	mock.ExpectExec(regexp.QuoteMeta("update categories set category_name = $2, slug = $3 where id = $1")).
		WithArgs(1, "testCategory2", "test-category-2").
		WillReturnResult(sqlmock.NewResult(1, 2))

	mux := http.NewServeMux()
	mux.HandleFunc("/categories/", e.handleRequestCategory)

	writer := httptest.NewRecorder()

	json := strings.NewReader(`{
		"category_name": "testCategory2",
		"slug": "test-category-2"
	}`)

	request, err := http.NewRequest("PUT", "/categories/test-category-1/", json)
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Fatalf("Response code is %v\n", writer.Code)
	}
}

func TestHandleDeleteCategory(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	e := Env{Db: db}

	row := sqlmock.NewRows([]string{"id", "category_name", "slug"}).
		AddRow(1, "testCategory1", "test-category-1")

	mock.ExpectQuery(regexp.QuoteMeta("select * from categories where slug = $1")).
		WithArgs("test-category-1").
		WillReturnRows(row)

	mock.ExpectExec(regexp.QuoteMeta("delete from categories where id = $1")).
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(1, 3))

	mux := http.NewServeMux()
	mux.HandleFunc("/categories/", e.handleRequestCategory)

	writer := httptest.NewRecorder()

	request, err := http.NewRequest("DELETE", "/categories/test-category-1/", nil)
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Fatalf("Response code is %v\n", writer.Code)
	}
}

func TestHandleGetSubCategory(t *testing.T) {
	t.Run(
		"without query params",
		func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			e := Env{Db: db}

			rows := sqlmock.NewRows([]string{"id", "category_name", "slug", "parent_category_id"}).
				AddRow(1, "testSubCategory1", "test-sub-category-1", 1).
				AddRow(2, "testSubCategory2", "test-sub-category-2", 1)

			mock.ExpectQuery(regexp.QuoteMeta("select * from sub_categories")).
				WillReturnRows(rows)

			mux := http.NewServeMux()
			mux.HandleFunc("/sub-categories/", e.handleRequestSubCategory)

			writer := httptest.NewRecorder()

			request, err := http.NewRequest("GET", "/sub-categories/", nil)
			mux.ServeHTTP(writer, request)

			if writer.Code != 200 {
				t.Fatalf("Response code is %v\n", writer.Code)
			}

			subCategories := []SubCategory{}
			json.Unmarshal(writer.Body.Bytes(), &subCategories)

			expectedSubCategories := []SubCategory{
				{
					Id:               1,
					CategoryName:     "testSubCategory1",
					Slug:             "test-sub-category-1",
					ParentCategoryId: 1,
				},
				{
					Id:               2,
					CategoryName:     "testSubCategory2",
					Slug:             "test-sub-category-2",
					ParentCategoryId: 1,
				},
			}

			if !(reflect.DeepEqual(subCategories, expectedSubCategories)) {
				t.Fatalf("Wrong content, was expecting %v, but got %v\n", expectedSubCategories, subCategories)
			}
		},
	)
	t.Run(
		"with query params: category-name",
		func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			e := Env{Db: db}

			mock.ExpectQuery(regexp.QuoteMeta("select id from categories where slug = $1")).
				WithArgs("test-category-1").
				WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

			rows := sqlmock.NewRows([]string{"id", "category_name", "slug", "parent_category_id"}).
				AddRow(1, "testSubCategory1", "test-sub-category-1", 1).
				AddRow(2, "testSubCategory2", "test-sub-category-2", 1)

			mock.ExpectQuery(regexp.QuoteMeta("select * from sub_categories where parent_category_id = $1")).
				WithArgs(1).
				WillReturnRows(rows)

			mux := http.NewServeMux()
			mux.HandleFunc("/sub-categories/", e.handleRequestSubCategory)

			writer := httptest.NewRecorder()

			request, err := http.NewRequest("GET", "/sub-categories/?category-name=test-category-1", nil)
			mux.ServeHTTP(writer, request)

			if writer.Code != 200 {
				t.Fatalf("Response code is %v\n", writer.Code)
			}

			subCategories := []SubCategory{}
			json.Unmarshal(writer.Body.Bytes(), &subCategories)

			expectedSubCategories := []SubCategory{
				{
					Id:               1,
					CategoryName:     "testSubCategory1",
					Slug:             "test-sub-category-1",
					ParentCategoryId: 1,
				},
				{
					Id:               2,
					CategoryName:     "testSubCategory2",
					Slug:             "test-sub-category-2",
					ParentCategoryId: 1,
				},
			}

			if !(reflect.DeepEqual(subCategories, expectedSubCategories)) {
				t.Fatalf("Wrong content, was expecting %v, but got %v\n", expectedSubCategories, subCategories)
			}
		},
	)
}

func TestHandlePostSubCategory(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	e := Env{Db: db}

	mock.ExpectQuery(regexp.QuoteMeta("insert into sub_categories (category_name, slug, parent_category_id) values ($1, $2, $3) returning id")).
		WithArgs("testSubCategory1", "test-sub-category-1", 1).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	mux := http.NewServeMux()
	mux.HandleFunc("/sub-categories/", e.handleRequestSubCategory)

	writer := httptest.NewRecorder()

	json := strings.NewReader(`{
		"category_name": "testSubCategory1",
		"slug": "test-sub-category-1",
		"parent_category_id": 1
	}`)

	request, err := http.NewRequest("POST", "/sub-categories/", json)
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Fatalf("Response code is %v\n", writer.Code)
	}
}

func TestHandlePutSubCategory(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	e := Env{Db: db}

	row := sqlmock.NewRows([]string{"id", "category_name", "slug", "parent_category_id"}).
		AddRow(1, "testSubCategory1", "test-sub-category-1", 1)

	mock.ExpectQuery(regexp.QuoteMeta("select * from sub_categories where slug = $1")).
		WithArgs("test-sub-category-1").
		WillReturnRows(row)

	mock.ExpectExec(regexp.QuoteMeta("update sub_categories set category_name = $2, slug = $3, parent_category_id = $4 where id = $1")).
		WithArgs(1, "testSubCategory2", "test-sub-category-2", 2).
		WillReturnResult(sqlmock.NewResult(1, 3))

	mux := http.NewServeMux()
	mux.HandleFunc("/sub-categories/", e.handleRequestSubCategory)

	writer := httptest.NewRecorder()

	json := strings.NewReader(`{
		"category_name": "testSubCategory2",
		"slug": "test-sub-category-2",
		"parent_category_id": 2
	}`)

	request, err := http.NewRequest("PUT", "/sub-categories/test-sub-category-1", json)
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Fatalf("Response code is %v\n", writer.Code)
	}
}

func TestHandleDeleteSubCategory(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	e := Env{Db: db}

	row := sqlmock.NewRows([]string{"id", "category_name", "slug", "parent_category_id"}).
		AddRow(1, "testSubCategory1", "test-sub-category-1", 1)

	mock.ExpectQuery(regexp.QuoteMeta("select * from sub_categories where slug = $1")).
		WithArgs("test-sub-category-1").
		WillReturnRows(row)

	mock.ExpectExec(regexp.QuoteMeta("delete from sub_categories where id = $1")).
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(1, 4))

	mux := http.NewServeMux()
	mux.HandleFunc("/sub-categories/", e.handleRequestSubCategory)

	writer := httptest.NewRecorder()

	request, err := http.NewRequest("DELETE", "/sub-categories/test-sub-category-1", nil)
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Fatalf("Response code is %v\n", writer.Code)
	}
}

func TestHandleGetPosts(t *testing.T) {
	postCreatedAt, _ := time.Parse("2006-01-02 15:04:05.999999-07", "2006-01-02 15:04:05.999999-07")
	postUpdatedAt, _ := time.Parse("2006-01-02 15:04:05.999999-07", "2006-01-02 15:04:05.999999-07")
	t.Run(
		"without query params",
		func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			e := Env{Db: db}

			rows := sqlmock.NewRows([]string{"id", "category_id", "sub_category_id", "title", "slug", "eye_catching_img", "content", "meta_description", "is_public", "created_at", "updated_at"}).
				AddRow(1, 1, 1, "testPost1", "test-post-1", "test_post_1.png", "This is 1st post", "This is 1st post", false, postCreatedAt, postUpdatedAt).
				AddRow(2, 2, 2, "testPost2", "test-post-2", "test_post_2.png", "This is 2nd post", "This is 2nd post", false, postCreatedAt, postUpdatedAt)

			mock.ExpectQuery(regexp.QuoteMeta("select * from posts")).
				WillReturnRows(rows)

			mux := http.NewServeMux()
			mux.HandleFunc("/posts/", e.handleRequestPosts)

			writer := httptest.NewRecorder()

			request, err := http.NewRequest("GET", "/posts/", nil)
			mux.ServeHTTP(writer, request)

			if writer.Code != 200 {
				t.Fatalf("Response code is %v\n", writer.Code)
			}

			expectedPosts := []Post{
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
			posts := []Post{}
			json.Unmarshal(writer.Body.Bytes(), &posts)

			if !(reflect.DeepEqual(posts, expectedPosts)) {
				t.Fatalf("Wrong content, was expecting %v, but got %v\n", expectedPosts, posts)
			}
		},
	)
	t.Run(
		"with query params: category-name",
		func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			e := Env{Db: db}

			mock.ExpectQuery(regexp.QuoteMeta("select id from categories where slug = $1")).
				WithArgs("test-category-1").
				WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

			rows := sqlmock.NewRows([]string{"id", "category_id", "sub_category_id", "title", "slug", "eye_catching_img", "content", "meta_description", "is_public", "created_at", "updated_at"}).
				AddRow(1, 1, 1, "testPost1", "test-post-1", "test_post_1.png", "This is 1st post", "This is 1st post", false, postCreatedAt, postUpdatedAt).
				AddRow(2, 1, 2, "testPost2", "test-post-2", "test_post_2.png", "This is 2nd post", "This is 2nd post", false, postCreatedAt, postUpdatedAt)

			mock.ExpectQuery(regexp.QuoteMeta("select * from posts where category_id = $1")).
				WithArgs(1).
				WillReturnRows(rows)

			mux := http.NewServeMux()
			mux.HandleFunc("/posts/", e.handleRequestPosts)

			writer := httptest.NewRecorder()

			request, err := http.NewRequest("GET", "/posts/?category-name=test-category-1", nil)
			mux.ServeHTTP(writer, request)

			if writer.Code != 200 {
				t.Fatalf("Response code is %v\n", writer.Code)
			}

			expectedPosts := []Post{
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

			posts := []Post{}
			json.Unmarshal(writer.Body.Bytes(), &posts)

			if !(reflect.DeepEqual(posts, expectedPosts)) {
				t.Fatalf("Wrong content. was expecting %v, but got %v\n", expectedPosts, posts)
			}
		},
	)
	t.Run(
		"with query params: sub-category-name",
		func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			e := Env{Db: db}

			mock.ExpectQuery(regexp.QuoteMeta("select id from sub_categories where slug = $1")).
				WithArgs("test-sub-category-1").
				WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

			rows := sqlmock.NewRows([]string{"id", "category_id", "sub_category_id", "title", "slug", "eye_catching_img", "content", "meta_description", "is_public", "created_at", "updated_at"}).
				AddRow(1, 1, 1, "testPost1", "test-post-1", "test_post_1.png", "This is 1st post", "This is 1st post", false, postCreatedAt, postUpdatedAt).
				AddRow(2, 1, 1, "testPost2", "test-post-2", "test_post_2.png", "This is 2nd post", "This is 2nd post", false, postCreatedAt, postUpdatedAt)

			mock.ExpectQuery(regexp.QuoteMeta("select * from posts where sub_category_id = $1")).
				WithArgs(1).
				WillReturnRows(rows)

			mux := http.NewServeMux()
			mux.HandleFunc("/posts/", e.handleRequestPosts)

			writer := httptest.NewRecorder()

			request, err := http.NewRequest("GET", "/posts/?sub-category-name=test-sub-category-1", nil)
			mux.ServeHTTP(writer, request)

			if writer.Code != 200 {
				t.Fatalf("Response code is %v\n", writer.Code)
			}

			expectedPosts := []Post{
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

			posts := []Post{}
			json.Unmarshal(writer.Body.Bytes(), &posts)

			if !(reflect.DeepEqual(posts, expectedPosts)) {
				t.Fatalf("Wrong content. was expecting %v, but got %v\n", expectedPosts, posts)
			}
		},
	)
}

func TestHandleGetPost(t *testing.T) {
	postCreatedAt, _ := time.Parse("2006-01-02 15:04:05.999999-07", "2006-01-02 15:04:05.999999-07")
	postUpdatedAt, _ := time.Parse("2006-01-02 15:04:05.999999-07", "2006-01-02 15:04:05.999999-07")

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	e := Env{Db: db}

	row := sqlmock.NewRows([]string{"id", "category_id", "sub_category_id", "title", "slug", "eye_catching_img", "content", "meta_description", "is_public", "created_at", "updated_at"}).
		AddRow(1, 1, 1, "testPost1", "test-post-1", "test_post_1.png", "This is 1st post", "This is 1st post", false, postCreatedAt, postUpdatedAt)

	mock.ExpectQuery(regexp.QuoteMeta("select * from posts where slug = $1")).
		WithArgs("test-post-1").
		WillReturnRows(row)

	mux := http.NewServeMux()
	mux.HandleFunc("/posts/", e.handleRequestPosts)

	writer := httptest.NewRecorder()

	request, err := http.NewRequest("GET", "/posts/test-post-1/", nil)
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Fatalf("Response code is %v\n", writer.Code)
	}

	post := Post{}
	json.Unmarshal(writer.Body.Bytes(), &post)

	expectedPost := Post{
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

func TestHandlePostPost(t *testing.T) {
	postCreatedAt, _ := time.Parse("2006-01-02 15:04:05.999999-07", "2006-01-02 15:04:05.999999-07")
	postUpdatedAt, _ := time.Parse("2006-01-02 15:04:05.999999-07", "2006-01-02 15:04:05.999999-07")
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	e := Env{Db: db}

	mock.ExpectQuery(regexp.QuoteMeta("insert into posts (category_id, sub_category_id, title, slug, eye_catching_img, content, meta_description, is_public) values ($1, $2, $3, $4, $5, $6, $7, $8) returning id, created_at, updated_at")).
		WithArgs(1, 1, "testPost1", "test-post-1", "test_post_1.png", "This is 1st post", "This is 1st post", false).
		WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at"}).AddRow(1, postCreatedAt, postUpdatedAt))

	mux := http.NewServeMux()
	mux.HandleFunc("/posts/", e.handleRequestPosts)

	writer := httptest.NewRecorder()

	json := strings.NewReader(`{
		"category_id": 1,
		"sub_category_id": 1,
		"title": "testPost1",
		"slug": "test-post-1",
		"eye_catching_img": "test_post_1.png",
		"content": "This is 1st post",
		"meta_description": "This is 1st post",
		"is_public": "false"
	}`)

	request, err := http.NewRequest("POST", "/posts/", json)
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Fatalf("Response code is %v\n", writer.Code)
	}
}

func TestHandlePutPost(t *testing.T) {
	postCreatedAt, _ := time.Parse("2006-01-02 15:04:05.999999-07", "2006-01-02 15:04:05.999999-07")
	postUpdatedAt, _ := time.Parse("2006-01-02 15:04:05.999999-07", "2006-01-02 15:04:05.999999-07")
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	e := Env{Db: db}

	row := sqlmock.NewRows([]string{"id", "category_id", "sub_category_id", "title", "slug", "eye_catching_img", "content", "meta_description", "is_public", "created_at", "updated_at"}).
		AddRow(1, 1, 1, "testPost1", "test-post-1", "test_post_1.png", "This is 1st post", "This is 1st post", false, postCreatedAt, postUpdatedAt)

	mock.ExpectQuery(regexp.QuoteMeta("select * from posts where slug = $1")).
		WithArgs("test-post-1").
		WillReturnRows(row)

	mock.ExpectExec(regexp.QuoteMeta("update posts set category_id = $2, sub_category_id = $3, title = $4, slug = $5, eye_catching_img = $6, content = $7, meta_description = $8, is_public = $9 where id = $1")).
		WithArgs(1, 1, 1, "testPost2", "test-post-2", "test_post_2.png", "This is 2nd post", "This is 2nd post", false).
		WillReturnResult(sqlmock.NewResult(1, 8))

	mux := http.NewServeMux()
	mux.HandleFunc("/posts/", e.handleRequestPosts)

	writer := httptest.NewRecorder()

	json := strings.NewReader(`{
		"category_id": 1,
		"sub_category_id": 1,
		"title": "testPost2",
		"slug": "test-post-2",
		"eye_catching_img": "test_post_2.png",
		"content": "This is 2nd post",
		"meta_description": "This is 2nd post",
		"is_public": "false"
	}`)

	request, err := http.NewRequest("PUT", "/posts/test-post-1/", json)
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Fatalf("Response code is %v\n", writer.Code)
	}
}

func TestHandleDeletePost(t *testing.T) {
	postCreatedAt, _ := time.Parse("2006-01-02 15:04:05.999999-07", "2006-01-02 15:04:05.999999-07")
	postUpdatedAt, _ := time.Parse("2006-01-02 15:04:05.999999-07", "2006-01-02 15:04:05.999999-07")
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	e := Env{Db: db}

	row := sqlmock.NewRows([]string{"id", "category_id", "sub_category_id", "title", "slug", "eye_catching_img", "content", "meta_description", "is_public", "created_at", "updated_at"}).
		AddRow(1, 1, 1, "testPost1", "test-post-1", "test_post_1.png", "This is 1st post", "This is 1st post", false, postCreatedAt, postUpdatedAt)

	mock.ExpectQuery(regexp.QuoteMeta("select * from posts where slug = $1")).
		WithArgs("test-post-1").
		WillReturnRows(row)

	mock.ExpectExec(regexp.QuoteMeta("delete from posts where id = $1")).
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(1, 11))

	mux := http.NewServeMux()
	mux.HandleFunc("/posts/", e.handleRequestPosts)

	writer := httptest.NewRecorder()

	request, err := http.NewRequest("DELETE", "/posts/test-post-1/", nil)
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Fatalf("Response code is %v\n", writer.Code)
	}
}
