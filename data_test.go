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
		t.Fatalf("Wrong content, was expecting %v, but got %v", expectedCategory, categories)
	}
}
