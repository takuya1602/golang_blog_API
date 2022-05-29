package handler

import (
	"backend/app/common/dto"
	mocks "backend/mocks/service"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPostHandler_GetPosts(t *testing.T) {
	postCreatedAt, _ := time.Parse("2006-01-02 15:04:05.999999-07", "2006-01-02 15:04:05.999999-07")
	postUpdatedAt, _ := time.Parse("2006-01-02 15:04:05.999999-07", "2006-01-02 15:04:05.999999-07")

	postDtos := []dto.PostModel{
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
	t.Run(
		"with query param: category-name",
		func(t *testing.T) {
			s := new(mocks.IPostService)

			queryParams := map[string][]string{
				"category-name": {
					"test-category-1",
				},
			}
			s.On("GetPosts", queryParams).Return(postDtos, nil)

			h := NewPostHandler(s)

			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/posts?category-name=test-category-1", nil)

			err := h.GetPosts(w, r)

			assert.NoError(t, err)
			s.AssertExpectations(t)
		},
	)
	t.Run(
		"with query param: sub-category-name",
		func(t *testing.T) {
			s := new(mocks.IPostService)

			queryParams := map[string][]string{
				"sub-category-name": {
					"test-sub-category-1",
				},
			}

			s.On("GetPosts", queryParams).Return(postDtos, nil)

			h := NewPostHandler(s)

			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/posts?sub-category-name=test-sub-category-1", nil)

			err := h.GetPosts(w, r)

			assert.NoError(t, err)
			s.AssertExpectations(t)
		},
	)
	t.Run(
		"without query param",
		func(t *testing.T) {
			s := new(mocks.IPostService)

			queryParams := map[string][]string{}

			s.On("GetPosts", queryParams).Return(postDtos, nil)

			h := NewPostHandler(s)

			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/posts/", nil)

			err := h.GetPosts(w, r)

			assert.NoError(t, err)
			s.AssertExpectations(t)
		},
	)
}

func TestPostHandler_CRUD(t *testing.T) {
	// Create, Update, Delete allow
	// created_at, updated_at, category_id/name/slug, sub_category_name/slug
	// empty, because "posts" table of postgreSQL database dont't have these columns.
	postDto := dto.PostModel{
		Id:              1,
		Title:           "testPost1",
		Slug:            "test-post-1",
		EyeCatchingImg:  "test_post_1.png",
		Content:         "This is 1st post",
		MetaDescription: "This is 1st post",
		IsPublic:        false,
		SubCategoryId:   1,
	}

	json := strings.NewReader(`{
		"id": 1,
		"title": "testPost1",
		"slug": "test-post-1",
		"eye_catching_img": "test_post_1.png",
		"content": "This is 1st post",
		"meta_description": "This is 1st post",
		"is_public": false,
		"sub_category_id": 1
	}`)

	t.Run(
		"GetPostBySlug",
		func(t *testing.T) {
			s := new(mocks.IPostService)

			s.On("GetPostBySlug", postDto.Slug).Return(postDto, nil)

			h := NewPostHandler(s)

			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/posts/test-post-1/", nil)

			err := h.GetPostBySlug(w, r, postDto.Slug)

			assert.NoError(t, err)
			s.AssertExpectations(t)
		},
	)

	t.Run(
		"Create",
		func(t *testing.T) {
			s := new(mocks.IPostService)

			s.On("Create", postDto).Return(nil)

			h := NewPostHandler(s)

			w := httptest.NewRecorder()
			r := httptest.NewRequest("POST", "/posts/", json)

			err := h.Create(w, r)

			assert.NoError(t, err)
			s.AssertExpectations(t)
		},
	)

	t.Run(
		"Update",
		func(t *testing.T) {
			s := new(mocks.IPostService)

			s.On("GetPostBySlug", postDto.Slug).Return(postDto, nil)
			s.On("Update", postDto).Return(nil)

			h := NewPostHandler(s)

			w := httptest.NewRecorder()
			r := httptest.NewRequest("PUT", "/posts/test-post-1/", json)

			err := h.Update(w, r)

			assert.NoError(t, err)
			s.AssertExpectations(t)
		},
	)

	t.Run(
		"Delete",
		func(t *testing.T) {
			s := new(mocks.IPostService)

			s.On("GetPostBySlug", postDto.Slug).Return(postDto, nil)
			s.On("Delete", postDto).Return(nil)

			h := NewPostHandler(s)

			w := httptest.NewRecorder()
			r := httptest.NewRequest("DELETE", "/posts/test-post-1/", nil)

			err := h.Delete(w, r)

			assert.NoError(t, err)
			s.AssertExpectations(t)
		},
	)
}
