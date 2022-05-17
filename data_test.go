package main

import (
	"reflect"
	"regexp"
	"testing"

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
