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

func TestPostHandler_Get(t *testing.T) {
	t.Run(
		"with query param: category-name",
		func(t *testing.T) {
			postCreatedAt, _ := time.Parse("2006-01-02 15:04:05.999999-07", "2006-01-02 15:04:05.999999-07")
			postUpdatedAt, _ := time.Parse("2006-01-02 15:04:05.999999-07", "2006-01-02 15:04:05.999999-07")
			postDtos := []dto.PostModel{
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
			categorySlug := "test-category-1"

			s := new(mocks.IPostService)

			s.On("GetWithCategoryQuery", categorySlug).Return(postDtos, nil)

			h := NewPostHandler(s)

			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/posts?category-name=test-category-1", nil)

			err := h.Get(w, r)

			assert.NoError(t, err)
			s.AssertExpectations(t)
		},
	)
	t.Run(
		"with query param: sub-category-name",
		func(t *testing.T) {
			postCreatedAt, _ := time.Parse("2006-01-02 15:04:05.999999-07", "2006-01-02 15:04:05.999999-07")
			postUpdatedAt, _ := time.Parse("2006-01-02 15:04:05.999999-07", "2006-01-02 15:04:05.999999-07")
			postDtos := []dto.PostModel{
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
			subCategorySlug := "test-sub-category-1"

			s := new(mocks.IPostService)

			s.On("GetWithSubCategoryQuery", subCategorySlug).Return(postDtos, nil)

			h := NewPostHandler(s)

			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/posts?sub-category-name=test-sub-category-1", nil)

			err := h.Get(w, r)

			assert.NoError(t, err)
			s.AssertExpectations(t)
		},
	)
	t.Run(
		"without query param",
		func(t *testing.T) {
			postCreatedAt, _ := time.Parse("2006-01-02 15:04:05.999999-07", "2006-01-02 15:04:05.999999-07")
			postUpdatedAt, _ := time.Parse("2006-01-02 15:04:05.999999-07", "2006-01-02 15:04:05.999999-07")
			postDtos := []dto.PostModel{
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

			s := new(mocks.IPostService)

			s.On("GetAll").Return(postDtos, nil)

			h := NewPostHandler(s)

			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/posts/", nil)

			err := h.Get(w, r)

			assert.NoError(t, err)
			s.AssertExpectations(t)
		},
	)
}

func TestPostHandler_GetBySlug(t *testing.T) {
	postDto := dto.PostModel{
		Id:              1,
		CategoryId:      1,
		SubCategoryId:   1,
		Title:           "testPost1",
		Slug:            "test-post-1",
		EyeCatchingImg:  "test_post_1.png",
		Content:         "This is 1st post",
		MetaDescription: "This is 1st post",
		IsPublic:        true,
	}

	s := new(mocks.IPostService)

	s.On("GetBySlug", postDto.Slug).Return(postDto, nil)

	h := NewPostHandler(s)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/posts/test-post-1/", nil)

	err := h.GetBySlug(w, r, postDto.Slug)

	assert.NoError(t, err)
	s.AssertExpectations(t)
}

func TestPostHandler_Create(t *testing.T) {
	postDto := dto.PostModel{
		Id:              1,
		CategoryId:      1,
		SubCategoryId:   1,
		Title:           "testPost1",
		Slug:            "test-post-1",
		EyeCatchingImg:  "test_post_1.png",
		Content:         "This is 1st post",
		MetaDescription: "This is 1st post",
		IsPublic:        true,
	}
	json := strings.NewReader(`{
		"id": 1,
		"category_id": 1,
		"sub_category_id": 1,
		"title": "testPost1",
		"slug": "test-post-1",
		"eye_catching_img": "test_post_1.png",
		"content": "This is 1st post",
		"meta_description": "This is 1st post",
		"is_public": true
	}`)

	s := new(mocks.IPostService)

	s.On("Create", postDto).Return(nil)

	h := NewPostHandler(s)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/posts/", json)

	err := h.Create(w, r)

	assert.NoError(t, err)
	s.AssertExpectations(t)
}

func TestPostHandler_Update(t *testing.T) {
	postDto := dto.PostModel{
		Id:              1,
		CategoryId:      1,
		SubCategoryId:   1,
		Title:           "testPost1",
		Slug:            "test-post-1",
		EyeCatchingImg:  "test_post_1.png",
		Content:         "This is 1st post",
		MetaDescription: "This is 1st post",
		IsPublic:        true,
	}
	json := strings.NewReader(`{
		"id": 1,
		"category_id": 1,
		"sub_category_id": 1,
		"title": "testPost1",
		"slug": "test-post-1",
		"eye_catching_img": "test_post_1.png",
		"content": "This is 1st post",
		"meta_description": "This is 1st post",
		"is_public": true
	}`)

	s := new(mocks.IPostService)

	s.On("GetBySlug", postDto.Slug).Return(postDto, nil)
	s.On("Update", postDto).Return(nil)

	h := NewPostHandler(s)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("PUT", "/posts/test-post-1/", json)

	err := h.Update(w, r)

	assert.NoError(t, err)
	s.AssertExpectations(t)
}

func TestPostHandler_Delete(t *testing.T) {
	postDto := dto.PostModel{
		Id:              1,
		CategoryId:      1,
		SubCategoryId:   1,
		Title:           "testPost1",
		Slug:            "test-post-1",
		EyeCatchingImg:  "test_post_1.png",
		Content:         "This is 1st post",
		MetaDescription: "This is 1st post",
		IsPublic:        true,
	}

	s := new(mocks.IPostService)

	s.On("GetBySlug", postDto.Slug).Return(postDto, nil)
	s.On("Delete", postDto).Return(nil)

	h := NewPostHandler(s)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("DELETE", "/posts/test-post-1/", nil)

	err := h.Delete(w, r)

	assert.NoError(t, err)
	s.AssertExpectations(t)
}
