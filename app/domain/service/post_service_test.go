package service

import (
	"backend/app/common/dto"
	"backend/app/domain/entity"
	mocks "backend/mocks/repository"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func AssertPosts(t *testing.T, ret []dto.PostModel, posts []entity.Post) {
	for i, r := range ret {
		assert.Equal(t, r.Id, posts[i].Id)
		assert.Equal(t, r.CategoryId, posts[i].CategoryId)
		assert.Equal(t, r.SubCategoryId, posts[i].SubCategoryId)
		assert.Equal(t, r.Title, posts[i].Title)
		assert.Equal(t, r.Slug, posts[i].Slug)
		assert.Equal(t, r.EyeCatchingImg, posts[i].EyeCatchingImg)
		assert.Equal(t, r.Content, posts[i].Content)
		assert.Equal(t, r.MetaDescription, posts[i].MetaDescription)
		assert.Equal(t, r.IsPublic, posts[i].IsPublic)
		assert.Equal(t, r.CreatedAt, posts[i].CreatedAt)
		assert.Equal(t, r.UpdatedAt, posts[i].UpdatedAt)
	}
}

func TestPostService_GetPosts(t *testing.T) {
	postCreatedAt, _ := time.Parse("2006-01-02 15:04:05.999999-07", "2006-01-02 15:04:05.999999-07")
	postUpdatedAt, _ := time.Parse("2006-01-02 15:04:05.999999-07", "2006-01-02 15:04:05.999999-07")

	posts := []entity.Post{
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
		"with query-params: category-name",
		func(t *testing.T) {
			r := new(mocks.IPostRepository)

			queryParams := map[string][]string{
				"category-name": {
					"test-category-1",
				},
			}

			r.On("GetPosts", queryParams).Return(posts, nil)

			s := NewPostService(r)

			ret, err := s.GetPosts(queryParams)

			assert.NoError(t, err)
			AssertPosts(t, ret, posts)
			r.AssertExpectations(t)
		},
	)

	t.Run(
		"with query-params: sub-category-name",
		func(t *testing.T) {
			r := new(mocks.IPostRepository)

			queryParams := map[string][]string{
				"sub-category-name": {
					"test-sub-category-1",
				},
			}

			r.On("GetPosts", queryParams).Return(posts, nil)

			s := NewPostService(r)

			ret, err := s.GetPosts(queryParams)

			assert.NoError(t, err)
			AssertPosts(t, ret, posts)
			r.AssertExpectations(t)
		},
	)

	t.Run(
		"no query-params",
		func(t *testing.T) {
			r := new(mocks.IPostRepository)

			queryParams := map[string][]string{}

			r.On("GetPosts", queryParams).Return(posts, nil)

			s := NewPostService(r)

			ret, err := s.GetPosts(queryParams)

			assert.NoError(t, err)
			AssertPosts(t, ret, posts)
			r.AssertExpectations(t)
		},
	)
}

func TestPostService_CRUD(t *testing.T) {
	postCreatedAt, _ := time.Parse("2006-01-02 15:04:05.999999-07", "2006-01-02 15:04:05.999999-07")
	postUpdatedAt, _ := time.Parse("2006-01-02 15:04:05.999999-07", "2006-01-02 15:04:05.999999-07")

	post := entity.Post{
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
	}

	postDto := dto.PostModel{
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
	}

	t.Run(
		"GetPostBySlug",
		func(t *testing.T) {
			r := new(mocks.IPostRepository)

			r.On("GetPostBySlug", post.Slug).Return(post, nil)
			s := NewPostService(r)

			ret, err := s.GetPostBySlug(post.Slug)

			assert.NoError(t, err)
			assert.Equal(t, ret.Id, post.Id)
			assert.Equal(t, ret.Title, post.Title)
			assert.Equal(t, ret.Slug, post.Slug)
			assert.Equal(t, ret.EyeCatchingImg, post.EyeCatchingImg)
			assert.Equal(t, ret.Content, post.Content)
			assert.Equal(t, ret.MetaDescription, post.MetaDescription)
			assert.Equal(t, ret.IsPublic, post.IsPublic)
			assert.Equal(t, ret.CreatedAt, post.CreatedAt)
			assert.Equal(t, ret.UpdatedAt, post.UpdatedAt)
			assert.Equal(t, ret.CategoryId, post.CategoryId)
			assert.Equal(t, ret.CategoryName, post.CategoryName)
			assert.Equal(t, ret.CategorySlug, post.CategorySlug)
			assert.Equal(t, ret.SubCategoryId, post.SubCategoryId)
			assert.Equal(t, ret.SubCategoryName, post.SubCategoryName)
			assert.Equal(t, ret.SubCategorySlug, post.SubCategorySlug)
			r.AssertExpectations(t)
		},
	)

	t.Run(
		"Create",
		func(t *testing.T) {
			r := new(mocks.IPostRepository)

			r.On("Create", post).Return(nil)

			s := NewPostService(r)

			err := s.Create(postDto)

			assert.NoError(t, err)
			r.AssertExpectations(t)
		},
	)

	t.Run(
		"Update",
		func(t *testing.T) {
			r := new(mocks.IPostRepository)

			r.On("Update", post).Return(nil)

			s := NewPostService(r)

			err := s.Update(postDto)

			assert.NoError(t, err)
			r.AssertExpectations(t)
		},
	)

	t.Run(
		"Delete",
		func(t *testing.T) {
			r := new(mocks.IPostRepository)

			r.On("Delete", post).Return(nil)

			s := NewPostService(r)

			err := s.Delete(postDto)

			assert.NoError(t, err)
			r.AssertExpectations(t)
		},
	)
}
