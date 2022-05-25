package postgresql

import (
	"backend/app/domain/entity"
	"reflect"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestSubCategoryRepositoryGetAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "name", "slug", "parent_category_id"}).
		AddRow(1, "testSubCategory1", "test-sub-category-1", 1).
		AddRow(2, "testSubCategory2", "test-sub-category-2", 2)

	mock.ExpectQuery(regexp.QuoteMeta("select id, name, slug, parent_category_id from sub_categories")).
		WillReturnRows(rows)

	r := NewSubcategoryRepository(db)

	subCategories, err := r.GetAll()

	expectedSubCategories := []entity.SubCategory{
		{
			Id:               1,
			Name:             "testSubCategory1",
			Slug:             "test-sub-category-1",
			ParentCategoryId: 1,
		},
		{
			Id:               2,
			Name:             "testSubCategory2",
			Slug:             "test-sub-category-2",
			ParentCategoryId: 2,
		},
	}

	if !(reflect.DeepEqual(subCategories, expectedSubCategories)) {
		t.Fatalf("Wrong content, was expecting %v, but got %v\n", expectedSubCategories, subCategories)
	}
}

func TestSubcategoryRepositoryGetFilterParentCategory(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery(regexp.QuoteMeta("select id from categories where slug = $1")).
		WithArgs("test-category-1").
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	rows := sqlmock.NewRows([]string{"id", "name", "slug", "parent_category_id"}).
		AddRow(1, "testSubCategory1", "test-sub-category-1", 1).
		AddRow(2, "testSubCategory2", "test-sub-category-2", 1)

	mock.ExpectQuery(regexp.QuoteMeta("select * from sub_categories where parent_category_id = $1")).
		WithArgs(1).
		WillReturnRows(rows)

	r := NewSubcategoryRepository(db)

	subCategories, err := r.GetFilterParentCategory("test-category-1")
	if err != nil {
		t.Fatal(err)
	}

	expectedSubCategories := []entity.SubCategory{
		{
			Id:               1,
			Name:             "testSubCategory1",
			Slug:             "test-sub-category-1",
			ParentCategoryId: 1,
		},
		{
			Id:               2,
			Name:             "testSubCategory2",
			Slug:             "test-sub-category-2",
			ParentCategoryId: 1,
		},
	}

	if !(reflect.DeepEqual(subCategories, expectedSubCategories)) {
		t.Fatalf("Wrong content, was expecting %v, but got %v\n", expectedSubCategories, subCategories)
	}
}

func TestSubCategoryRepositoryGetBySlug(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "name", "slug", "parent_category_id"}).
		AddRow(1, "testSubCategory1", "test-sub-category-1", 1)

	mock.ExpectQuery(regexp.QuoteMeta("select id, name, slug, parent_category_id from sub_categories where slug = $1")).
		WithArgs("test-sub-category-1").
		WillReturnRows(rows)

	r := NewSubcategoryRepository(db)

	subCategory, err := r.GetBySlug("test-sub-category-1")

	expectedSubCategory := entity.SubCategory{
		Id:               1,
		Name:             "testSubCategory1",
		Slug:             "test-sub-category-1",
		ParentCategoryId: 1,
	}

	if !(reflect.DeepEqual(subCategory, expectedSubCategory)) {
		t.Fatalf("Wrong content, was expecting %v, but got %v\n", expectedSubCategory, subCategory)
	}
}

func TestSubCategoryRepositoryCreate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectExec(regexp.QuoteMeta("insert into sub_categories (name, slug, parent_category_id) values ($1, $2, $3)")).
		WithArgs("testSubCategory1", "test-sub-category-1", 1).
		WillReturnResult(sqlmock.NewResult(1, 3))

	r := NewSubcategoryRepository(db)

	subCategory := entity.SubCategory{
		Name:             "testSubCategory1",
		Slug:             "test-sub-category-1",
		ParentCategoryId: 1,
	}

	if err := r.Create(subCategory); err != nil {
		t.Fatal(err)
	}
}

func TestSubCategoryRepositoryUpdate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectExec(regexp.QuoteMeta("update sub_categories set name = $2, slug = $3, parent_category_id = $4 where id = $1")).
		WithArgs(1, "testSubCategory1", "test-sub-category-1", 1).
		WillReturnResult(sqlmock.NewResult(1, 3))

	r := NewSubcategoryRepository(db)

	subCategory := entity.SubCategory{
		Id:               1,
		Name:             "testSubCategory1",
		Slug:             "test-sub-category-1",
		ParentCategoryId: 1,
	}

	if err := r.Update(subCategory); err != nil {
		t.Fatal(err)
	}
}

func TestSubCategoryRepositoryDelete(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectExec(regexp.QuoteMeta("delete from sub_categories where id = $1")).
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(1, 4))

	r := NewSubcategoryRepository(db)

	subCategory := entity.SubCategory{
		Id:               1,
		Name:             "testSubCategory1",
		Slug:             "test-sub-category-1",
		ParentCategoryId: 1,
	}

	if err := r.Delete(subCategory); err != nil {
		t.Fatal(err)
	}
}
