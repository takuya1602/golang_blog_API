package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"regexp"
	"strings"
	"testing"

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
