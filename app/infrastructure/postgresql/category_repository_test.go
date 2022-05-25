package postgresql

import (
	"backend/app/domain/entity"
	"reflect"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestCategoryRepositoryGetAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "name", "slug"}).
		AddRow(1, "testCategory1", "test-category-1").
		AddRow(2, "testCategory2", "test-category-2")

	mock.ExpectQuery(regexp.QuoteMeta("select * from categories")).
		WillReturnRows(rows)

	r := NewCategoryRepository(db)

	categories, err := r.GetAll()
	if err != nil {
		t.Fatal(err)
	}

	expectedCategories := []entity.Category{
		{
			Id:   1,
			Name: "testCategory1",
			Slug: "test-category-1",
		},
		{
			Id:   2,
			Name: "testCategory2",
			Slug: "test-category-2",
		},
	}

	if !(reflect.DeepEqual(categories, expectedCategories)) {
		t.Fatalf("Wrong content, was expecting %v, but got %v\n", expectedCategories, categories)
	}
}

func TestCategoryRepositoryCreate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectExec(regexp.QuoteMeta("insert into categories (name, slug) values ($1, $2)")).
		WithArgs("testCategory1", "test-category-1").
		WillReturnResult(sqlmock.NewResult(1, 3))

	r := NewCategoryRepository(db)

	category := entity.Category{
		Name: "testCategory1",
		Slug: "test-category-1",
	}

	if err := r.Create(category); err != nil {
		t.Fatal(err)
	}
}

func TestCategoryRepositoryUpdate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectExec(regexp.QuoteMeta("update categories set name = $2, slug = $3 where id = $1")).
		WithArgs(1, "testCategory1", "test-category-1").
		WillReturnResult(sqlmock.NewResult(1, 2))

	r := NewCategoryRepository(db)

	category := entity.Category{
		Id:   1,
		Name: "testCategory1",
		Slug: "test-category-1",
	}

	if err := r.Update(category); err != nil {
		t.Fatal(err)
	}
}

func TestCategoryRepositoryDelete(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectExec(regexp.QuoteMeta("delete from categories where id = $1")).
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(1, 3))

	r := NewCategoryRepository(db)

	category := entity.Category{
		Id:   1,
		Name: "testCategory1",
		Slug: "test-category-1",
	}

	if err := r.Delete(category); err != nil {
		t.Fatal(err)
	}
}

func TestCategoryRepositoryGetBySlug(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "name", "slug"}).
		AddRow(1, "testCategory1", "test-category-1")

	mock.ExpectQuery(regexp.QuoteMeta("select id, name, slug from categories where slug = $1")).
		WithArgs("test-category-1").
		WillReturnRows(rows)

	r := NewCategoryRepository(db)

	category, err := r.GetBySlug("test-category-1")
	if err != nil {
		t.Fatal(err)
	}

	expectedCategory := entity.Category{
		Id:   1,
		Name: "testCategory1",
		Slug: "test-category-1",
	}

	if !(reflect.DeepEqual(category, expectedCategory)) {
		t.Fatalf("Wrong content, was expecting %v, but got %v\n", expectedCategory, category)
	}
}
