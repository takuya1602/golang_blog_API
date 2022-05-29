package postgresql

import (
	"backend/app/domain/entity"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func assertSubCategories(t *testing.T, ret []entity.SubCategory, subCategories []entity.SubCategory) {
	for i, r := range ret {
		assert.Equal(t, r.Id, subCategories[i].Id)
		assert.Equal(t, r.Name, subCategories[i].Name)
		assert.Equal(t, r.Slug, subCategories[i].Slug)
		assert.Equal(t, r.ParentCategoryId, subCategories[i].ParentCategoryId)
		assert.Equal(t, r.ParentCategoryName, subCategories[i].ParentCategoryName)
		assert.Equal(t, r.ParentCategorySlug, subCategories[i].ParentCategorySlug)
	}
}

func TestSubCategoryRepository_GetSubCategories(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	subCategories := []entity.SubCategory{
		{
			Id:                 1,
			Name:               "testSubCategory1",
			Slug:               "test-sub-category-1",
			ParentCategoryId:   1,
			ParentCategoryName: "testCategory1",
			ParentCategorySlug: "test-category-1",
		},
		{
			Id:                 2,
			Name:               "testSubCategory2",
			Slug:               "test-sub-category-2",
			ParentCategoryId:   1,
			ParentCategoryName: "testCategory1",
			ParentCategorySlug: "test-category-1",
		},
	}

	fields := []string{
		"id",
		"name",
		"slug",
		"parent_category_id",
		"parent_category_name",
		"parent_category_slug",
	}

	rows := sqlmock.NewRows(fields).
		AddRow(
			subCategories[0].Id,
			subCategories[0].Name,
			subCategories[0].Slug,
			subCategories[0].ParentCategoryId,
			subCategories[0].ParentCategoryName,
			subCategories[0].ParentCategorySlug,
		).
		AddRow(
			subCategories[1].Id,
			subCategories[1].Name,
			subCategories[1].Slug,
			subCategories[1].ParentCategoryId,
			subCategories[1].ParentCategoryName,
			subCategories[1].ParentCategorySlug,
		)

	t.Run(
		"with query params: category-name",
		func(t *testing.T) {
			mock.ExpectQuery(regexp.QuoteMeta(`
				select 
				sub_categories.id as id, sub_categories.name, sub_categories.slug,
				categories.id as parent_category_id, categories.name as parent_category_name, categories.slug as parent_category_slug 
				from sub_categories 
				inner join categories 
				on sub_categories.parent_category_id = categories.id
				where categories.slug = $1
			`)).WithArgs(subCategories[0].ParentCategorySlug).WillReturnRows(rows)

			r := NewSubcategoryRepository(db)

			queryParams := map[string][]string{
				"category-name": {
					subCategories[0].ParentCategorySlug,
				},
			}

			ret, err := r.GetSubCategories(queryParams)

			assert.NoError(t, err)
			assertSubCategories(t, ret, subCategories)
		},
	)

	t.Run(
		"without query params",
		func(t *testing.T) {
			mock.ExpectQuery(regexp.QuoteMeta(`
				select 
				sub_categories.id as id, sub_categories.name, sub_categories.slug, 
				categories.id as parent_category_id, categories.name as parent_category_name, categories.slug as parent_category_slug 
				from sub_categories 
				inner join categories
				on sub_categories.parent_category_id = categories.id
			`)).WillReturnRows(rows)

			r := NewSubcategoryRepository(db)

			var queryParams map[string][]string

			ret, err := r.GetSubCategories(queryParams)

			assert.NoError(t, err)
			assertSubCategories(t, ret, subCategories)
		},
	)
}

func TestSubCategoryRepository_CRUD(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	subCategory := entity.SubCategory{
		Id:                 1,
		Name:               "testSubCategory1",
		Slug:               "test-sub-category-1",
		ParentCategoryId:   1,
		ParentCategoryName: "testCategory1",
		ParentCategorySlug: "test-category-1",
	}

	fields := []string{
		"id",
		"name",
		"slug",
		"parent_category_id",
		"parent_category_name",
		"parent_category_slug",
	}

	rows := sqlmock.NewRows(fields).
		AddRow(
			subCategory.Id,
			subCategory.Name,
			subCategory.Slug,
			subCategory.ParentCategoryId,
			subCategory.ParentCategoryName,
			subCategory.ParentCategorySlug,
		)

	t.Run(
		"GetSubCategoryBySlug",
		func(t *testing.T) {
			mock.ExpectQuery(regexp.QuoteMeta(`
				select
				sub_categories.id as id, sub_categories.name, sub_categories.slug,
				categories.name as parent_category_name, categories.slug as parent_category_slug 
				from sub_categories
				inner join categories
				on sub_categories.parent_category_id = categories.id
				where sub_categories.slug = $1
			`)).WithArgs(subCategory.ParentCategorySlug).WillReturnRows(rows)

			r := NewSubcategoryRepository(db)

			ret, err := r.GetSubCategoryBySlug(subCategory.ParentCategorySlug)

			assert.NoError(t, err)
			assert.Equal(t, ret.Id, subCategory.Id)
			assert.Equal(t, ret.Name, subCategory.Name)
			assert.Equal(t, ret.Slug, subCategory.Slug)
			assert.Equal(t, ret.ParentCategoryId, subCategory.ParentCategoryId)
			assert.Equal(t, ret.ParentCategoryName, subCategory.ParentCategoryName)
			assert.Equal(t, ret.ParentCategorySlug, subCategory.ParentCategorySlug)
		},
	)

	t.Run(
		"Create",
		func(t *testing.T) {
			mock.ExpectExec(regexp.QuoteMeta("insert into sub_categories (name, slug, parent_category_id) values ($1, $2, $3)")).
				WithArgs(subCategory.Name, subCategory.Slug, subCategory.ParentCategoryId).
				WillReturnResult(sqlmock.NewResult(1, 4))

			r := NewSubcategoryRepository(db)

			err := r.Create(subCategory)

			assert.NoError(t, err)
		},
	)

	t.Run(
		"Update",
		func(t *testing.T) {
			mock.ExpectExec(regexp.QuoteMeta("update sub_categories set name = $2, slug = $3, parent_category_id = $4 where id = $1")).
				WithArgs(subCategory.Id, subCategory.Name, subCategory.Slug, subCategory.ParentCategoryId).
				WillReturnResult(sqlmock.NewResult(1, 3))

			r := NewSubcategoryRepository(db)

			err := r.Update(subCategory)

			assert.NoError(t, err)
		},
	)

	t.Run(
		"Delete",
		func(t *testing.T) {
			mock.ExpectExec(regexp.QuoteMeta("delete from sub_categories where id = $1")).
				WithArgs(subCategory.Id).
				WillReturnResult(sqlmock.NewResult(1, 6))

			r := NewSubcategoryRepository(db)

			err := r.Delete(subCategory)

			assert.NoError(t, err)
		},
	)
}
