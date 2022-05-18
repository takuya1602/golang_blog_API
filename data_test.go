package main

import (
	"reflect"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestRetrieveCategories(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "category_name", "slug"}).
		AddRow(1, "testCategory1", "test-category-1").
		AddRow(2, "testCategory2", "test-category-2")

	mock.ExpectQuery(regexp.QuoteMeta("select * from categories")).WillReturnRows(rows)

	categories, err := retrieveCategories(db)
	if err != nil {
		t.Fatal(err)
	}

	expectedCategory := []Category{
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

	if !(reflect.DeepEqual(categories, expectedCategory)) {
		t.Fatalf("Wrong content, was expecting %v, but got %v\n", expectedCategory, categories)
	}
}

func TestRetrieveCategory(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	row := sqlmock.NewRows([]string{"id", "category_name", "slug"}).
		AddRow(1, "testCategory1", "test-category-1")

	mock.ExpectQuery(regexp.QuoteMeta("select * from categories where slug = $1")).
		WithArgs("test-category-1").
		WillReturnRows(row)

	category, err := retrieveCategory(db, "test-category-1")
	if err != nil {
		t.Fatal(err)
	}

	expectedCategory := Category{
		Id:           1,
		CategoryName: "testCategory1",
		Slug:         "test-category-1",
	}

	if !(reflect.DeepEqual(category, expectedCategory)) {
		t.Fatalf("Wrong content, was expecting %v, but got %v\n", expectedCategory, category)
	}
}

func TestRetrieveSubCategories(t *testing.T) {
	expectedSubCategories1 := []SubCategory{
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

	expectedSubCategories2 := []SubCategory{
		{
			Id:               3,
			CategoryName:     "testSubCategory3",
			Slug:             "test-sub-category-3",
			ParentCategoryId: 2,
		},
		{
			Id:               4,
			CategoryName:     "testSubCategory4",
			Slug:             "test-sub-category-4",
			ParentCategoryId: 2,
		},
	}

	expectedSubCategories := append(expectedSubCategories1, expectedSubCategories2...)

	t.Run(
		"without query params",
		func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			rows := sqlmock.NewRows([]string{"id", "category_name", "slug", "parent_category_id"}).
				AddRow(1, "testSubCategory1", "test-sub-category-1", 1).
				AddRow(2, "testSubCategory2", "test-sub-category-2", 1).
				AddRow(3, "testSubCategory3", "test-sub-category-3", 2).
				AddRow(4, "testSubCategory4", "test-sub-category-4", 2)

			mock.ExpectQuery(regexp.QuoteMeta("select * from sub_categories")).
				WillReturnRows(rows)

			var queryParams map[string][]string // r.URL.Query() returns query params in map[string][]string
			subCategories, err := retrieveSubCategories(db, queryParams)

			if !(reflect.DeepEqual(subCategories, expectedSubCategories)) {
				t.Fatalf("Wrong content, was expecting %v, but got %v", expectedSubCategories, subCategories)
			}
		},
	)
	t.Run(
		"with query params",
		func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			// In case parent_category_id is 1
			categoryId := sqlmock.NewRows([]string{"id"}).
				AddRow(1)

			mock.ExpectQuery(regexp.QuoteMeta("select id from categories where slug = $1")).
				WithArgs("test-category-1").
				WillReturnRows(categoryId)

			rows := sqlmock.NewRows([]string{"id", "category_name", "slug", "parent_category_id"}).
				AddRow(1, "testSubCategory1", "test-sub-category-1", 1).
				AddRow(2, "testSubCategory2", "test-sub-category-2", 1)

			mock.ExpectQuery(regexp.QuoteMeta("select * from sub_categories where parent_category_id = $1")).
				WithArgs(1).
				WillReturnRows(rows)

			queryParams := map[string][]string{
				"category-name": {"test-category-1"},
			}

			subCategories, err := retrieveSubCategories(db, queryParams)

			if !(reflect.DeepEqual(subCategories, expectedSubCategories1)) {
				t.Fatalf("Wrong content, was expecting %v, but got %v", expectedSubCategories, subCategories)
			}

			// In case parent_category_id is 2
			categoryId = sqlmock.NewRows([]string{"id"}).
				AddRow(2)

			mock.ExpectQuery(regexp.QuoteMeta("select id from categories where slug = $1")).
				WithArgs("test-category-2").
				WillReturnRows(categoryId)

			rows = sqlmock.NewRows([]string{"id", "category_name", "slug", "parent_category_id"}).
				AddRow(3, "testSubCategory3", "test-sub-category-3", 2).
				AddRow(4, "testSubCategory4", "test-sub-category-4", 2)

			mock.ExpectQuery(regexp.QuoteMeta("select * from sub_categories where parent_category_id = $1")).
				WithArgs(2).
				WillReturnRows(rows)

			queryParams = map[string][]string{
				"category-name": {"test-category-2"},
			}

			subCategories, err = retrieveSubCategories(db, queryParams)

			if !(reflect.DeepEqual(subCategories, expectedSubCategories2)) {
				t.Fatalf("Wrong content, was expecting %v, but got %v", expectedSubCategories, subCategories)
			}
		},
	)
}

func TestRetrieveSubCategory(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	row := sqlmock.NewRows([]string{"id", "category_name", "slug", "parent_category_id"}).
		AddRow(1, "testSubCategory1", "test-sub-category-1", 1)

	mock.ExpectQuery(regexp.QuoteMeta("select * from sub_categories where slug = $1")).
		WithArgs("test-sub-category-1").
		WillReturnRows(row)

	subCategory, err := retrieveSubCategory(db, "test-sub-category-1")
	if err != nil {
		t.Fatal(err)
	}

	expectedSubCategory := SubCategory{
		Id:               1,
		CategoryName:     "testSubCategory1",
		Slug:             "test-sub-category-1",
		ParentCategoryId: 1,
	}

	if !(reflect.DeepEqual(subCategory, expectedSubCategory)) {
		t.Fatalf("Wrong content, was expecting %v, but got %v\n", expectedSubCategory, subCategory)
	}
}

func TestRetrievePosts(t *testing.T) {
	postCreatedAt, _ := time.Parse("2006-01-02 15:04:05.999999-07", "2006-01-02 15:04:05.999999-07")
	postUpdatedAt, _ := time.Parse("2006-01-02 15:04:05.999999-07", "2006-01-02 15:04:05.999999-07")
	expectedPosts1 := []Post{
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
	}

	expectedPosts2 := []Post{
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

	var expectedPosts []Post
	expectedPosts = append(expectedPosts, expectedPosts1...)
	expectedPosts = append(expectedPosts, expectedPosts2...)

	t.Run(
		"without query params",
		func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			rows := sqlmock.NewRows([]string{"id", "category_id", "sub_category_id", "title", "slug", "eye_catching_img", "content", "meta_description", "is_public", "created_at", "updated_at"}).
				AddRow(1, 1, 1, "testPost1", "test-post-1", "test_post_1.png", "This is 1st post", "This is 1st post", false, postCreatedAt, postUpdatedAt).
				AddRow(2, 1, 2, "testPost2", "test-post-2", "test_post_2.png", "This is 2nd post", "This is 2nd post", false, postCreatedAt, postUpdatedAt)

			mock.ExpectQuery(regexp.QuoteMeta("select * from posts")).
				WillReturnRows(rows)

			var queryParams map[string][]string
			posts, err := retrievePosts(db, queryParams)

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

			categoryId := sqlmock.NewRows([]string{"id"}).
				AddRow(1)

			mock.ExpectQuery(regexp.QuoteMeta("select id from categories where slug = $1")).
				WithArgs("test-category-1").
				WillReturnRows(categoryId)

			rows := sqlmock.NewRows([]string{"id", "category_id", "sub_category_id", "title", "slug", "eye_catching_img", "content", "meta_description", "is_public", "created_at", "updated_at"}).
				AddRow(1, 1, 1, "testPost1", "test-post-1", "test_post_1.png", "This is 1st post", "This is 1st post", false, postCreatedAt, postUpdatedAt).
				AddRow(2, 1, 2, "testPost2", "test-post-2", "test_post_2.png", "This is 2nd post", "This is 2nd post", false, postCreatedAt, postUpdatedAt)

			mock.ExpectQuery(regexp.QuoteMeta("select * from posts where category_id = $1")).
				WithArgs(1).
				WillReturnRows(rows)

			queryParams := map[string][]string{
				"category-name": {"test-category-1"},
			}

			posts, err := retrievePosts(db, queryParams)
			if err != nil {
				panic(err)
			}

			if !(reflect.DeepEqual(posts, expectedPosts)) {
				t.Fatalf("Wrong content, was expecting %v, but got %v\n", expectedPosts, posts)
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

			subCategoryId := sqlmock.NewRows([]string{"id"}).
				AddRow(2)

			mock.ExpectQuery(regexp.QuoteMeta("select id from sub_categories where slug = $1")).
				WithArgs("test-sub-category-2").
				WillReturnRows(subCategoryId)

			row := sqlmock.NewRows([]string{"id", "category_id", "sub_category_id", "title", "slug", "eye_catching_img", "content", "meta_description", "is_public", "created_at", "updated_at"}).
				AddRow(2, 1, 2, "testPost2", "test-post-2", "test_post_2.png", "This is 2nd post", "This is 2nd post", false, postCreatedAt, postUpdatedAt)

			mock.ExpectQuery(regexp.QuoteMeta("select * from posts where sub_category_id = $1")).
				WithArgs(2).
				WillReturnRows(row)

			queryParams := map[string][]string{
				"sub-category-name": {"test-sub-category-2"},
			}
			posts, err := retrievePosts(db, queryParams)

			if !(reflect.DeepEqual(posts, expectedPosts2)) {
				t.Fatalf("Wrong content, was expecting %v, but got %v\n", expectedPosts2, posts)
			}
		},
	)
}

func TestRetrievePost(t *testing.T) {
	postCreatedAt, _ := time.Parse("2006-01-02 15:04:05.999999-07", "2006-01-02 15:04:05.999999-07")
	postUpdatedAt, _ := time.Parse("2006-01-02 15:04:05.999999-07", "2006-01-02 15:04:05.999999-07")
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	row := sqlmock.NewRows([]string{"id", "category_id", "sub_category_id", "title", "slug", "eye_catching_img", "content", "meta_description", "is_public", "created_at", "updated_at"}).
		AddRow(1, 1, 1, "testPost1", "test-post-1", "test_post_1.png", "This is 1st post", "This is 1st post", false, postCreatedAt, postUpdatedAt)

	mock.ExpectQuery(regexp.QuoteMeta("select * from posts where slug = $1")).
		WithArgs("test-post-1").
		WillReturnRows(row)

	post, err := retrievePost(db, "test-post-1")
	if err != nil {
		panic(err)
	}

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

func TestCategoryCreate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery(regexp.QuoteMeta("insert into categories (category_name, slug) values ($1, $2) returning id")).
		WithArgs("testCategory1", "test-category-1").
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	category := Category{
		CategoryName: "testCategory1",
		Slug:         "test-category-1",
	}

	if err := category.create(db); err != nil {
		t.Fatal(err)
	}
}

func TestCategoryUpdate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectExec(regexp.QuoteMeta("update categories set category_name = $2, slug = $3 where id = $1")).
		WithArgs(1, "testCategory2", "test-category-2").
		WillReturnResult(sqlmock.NewResult(1, 2))

	category := Category{
		Id:           1,
		CategoryName: "testCategory2",
		Slug:         "test-category-2",
	}

	if err := category.update(db); err != nil {
		t.Fatal(err)
	}
}

func TestCategoryDelete(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectExec(regexp.QuoteMeta("delete from categories where id = $1")).
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(1, 3))

	category := Category{
		Id:           1,
		CategoryName: "testCategory2",
		Slug:         "test-category-2",
	}

	if err := category.delete(db); err != nil {
		t.Fatal(err)
	}
}

func TestSubCategoryMethod(t *testing.T) {
	t.Run(
		"create",
		func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			mock.ExpectQuery(regexp.QuoteMeta("insert into sub_categories (category_name, slug, parent_category_id) values ($1, $2, $3) returning id")).
				WithArgs("testSubCategory1", "test-sub-category-1", 1).
				WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

			subCategory := SubCategory{
				CategoryName:     "testSubCategory1",
				Slug:             "test-sub-category-1",
				ParentCategoryId: 1,
			}

			if err := subCategory.create(db); err != nil {
				t.Fatal(err)
			}
		},
	)
	t.Run(
		"update",
		func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			mock.ExpectExec(regexp.QuoteMeta("update sub_categories set category_name = $2, slug = $3, parent_category_id = $4 where id = $1")).
				WithArgs(1, "testSubCategory1", "test-sub-category-1", 1).
				WillReturnResult(sqlmock.NewResult(1, 3))

			subCategory := SubCategory{
				Id:               1,
				CategoryName:     "testSubCategory1",
				Slug:             "test-sub-category-1",
				ParentCategoryId: 1,
			}

			if err := subCategory.update(db); err != nil {
				t.Fatal(err)
			}
		},
	)
	t.Run(
		"delete",
		func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			mock.ExpectExec(regexp.QuoteMeta("delete from sub_categories where id = $1")).
				WithArgs(1).
				WillReturnResult(sqlmock.NewResult(1, 4))

			subCategory := SubCategory{
				Id:               1,
				CategoryName:     "testSubCategory1",
				Slug:             "test-sub-category-1",
				ParentCategoryId: 1,
			}

			if err := subCategory.delete(db); err != nil {
				t.Fatal(err)
			}
		},
	)
}

func TestPost(t *testing.T) {
	postCreatedAt, _ := time.Parse("2006-01-02 15:04:05.999999-07", "2006-01-02 15:04:05.999999-07")
	postUpdatedAt, _ := time.Parse("2006-01-02 15:04:05.999999-07", "2006-01-02 15:04:05.999999-07")

	t.Run(
		"create",
		func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			mock.ExpectQuery(regexp.QuoteMeta("insert into posts (category_id, sub_category_id, title, slug, eye_catching_img, content, meta_description, is_public) values ($1, $2, $3, $4, $5, $6, $7, $8) returning id, created_at, updated_at")).
				WithArgs(1, 1, "testPost1", "test-post-1", "test_post_1.png", "This is 1st post", "This is 1st post", false).
				WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at"}).AddRow(1, postCreatedAt, postUpdatedAt))

			post := Post{
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

			if err := post.create(db); err != nil {
				t.Fatal(err)
			}
		},
	)
	t.Run(
		"update",
		func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			mock.ExpectExec(regexp.QuoteMeta("update posts set category_id = $2, sub_category_id = $3, title = $4, slug = $5, eye_catching_img = $6, content = $7, meta_description = $8, is_public = $9 where id = $1")).
				WithArgs(1, 1, 1, "testPost1", "test-post-1", "test_post_1.png", "This is 1st post", "This is 1st post", false).
				WillReturnResult(sqlmock.NewResult(1, 10))

			post := Post{
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

			if err := post.update(db); err != nil {
				t.Fatal(err)
			}
		},
	)
	t.Run(
		"delete",
		func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			mock.ExpectExec(regexp.QuoteMeta("delete from posts where id = $1")).
				WithArgs(1).
				WillReturnResult(sqlmock.NewResult(1, 9))

			post := Post{
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

			if err := post.delete(db); err != nil {
				t.Fatal(err)
			}
		},
	)
}
