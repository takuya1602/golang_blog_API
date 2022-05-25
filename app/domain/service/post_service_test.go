package service

import (
	"backend/app/common/dto"
	"backend/app/domain/entity"
	mocks "backend/mocks/repository"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPostService_GetAll(t *testing.T) {
	postCreatedAt, _ := time.Parse("2006-01-02 15:04:05.999999-07", "2006-01-02 15:04:05.999999-07")
	postUpdatedAt, _ := time.Parse("2006-01-02 15:04:05.999999-07", "2006-01-02 15:04:05.999999-07")

	posts := []entity.Post{
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

	r := new(mocks.IPostRepository)

	r.On("GetAll").Return(posts, nil)

	s := NewPostService(r)

	ret, err := s.GetAll()

	assert.NoError(t, err)
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
	r.AssertExpectations(t)
}

func TestPostService_GetWithCategoryQuery(t *testing.T) {
	postCreatedAt, _ := time.Parse("2006-01-02 15:04:05.999999-07", "2006-01-02 15:04:05.999999-07")
	postUpdatedAt, _ := time.Parse("2006-01-02 15:04:05.999999-07", "2006-01-02 15:04:05.999999-07")

	posts := []entity.Post{
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
	categoryName := "testCategory1"

	r := new(mocks.IPostRepository)

	r.On("GetFilterCategory", categoryName).Return(posts, nil)

	s := NewPostService(r)

	ret, err := s.GetWithCategoryQuery(categoryName)

	assert.NoError(t, err)
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
	r.AssertExpectations(t)
}

func TestPostService_GetWithSubCategoryQuery(t *testing.T) {
	postCreatedAt, _ := time.Parse("2006-01-02 15:04:05.999999-07", "2006-01-02 15:04:05.999999-07")
	postUpdatedAt, _ := time.Parse("2006-01-02 15:04:05.999999-07", "2006-01-02 15:04:05.999999-07")

	posts := []entity.Post{
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

	subCategoryName := "test-sub-category-1"

	r := new(mocks.IPostRepository)

	r.On("GetFilterSubCategory", subCategoryName).Return(posts, nil)

	s := NewPostService(r)

	ret, err := s.GetWithSubCategoryQuery(subCategoryName)

	assert.NoError(t, err)
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
	r.AssertExpectations(t)
}

func TestPostService_GetBySlug(t *testing.T) {
	postCreatedAt, _ := time.Parse("2006-01-02 15:04:05.999999-07", "2006-01-02 15:04:05.999999-07")
	postUpdatedAt, _ := time.Parse("2006-01-02 15:04:05.999999-07", "2006-01-02 15:04:05.999999-07")

	post := entity.Post{
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

	slug := "test-post-1"

	r := new(mocks.IPostRepository)

	r.On("GetBySlug", slug).Return(post, nil)

	s := NewPostService(r)

	ret, err := s.GetBySlug(slug)

	assert.NoError(t, err)
	assert.Equal(t, ret.Id, post.Id)
	assert.Equal(t, ret.CategoryId, post.CategoryId)
	assert.Equal(t, ret.SubCategoryId, post.SubCategoryId)
	assert.Equal(t, ret.Title, post.Title)
	assert.Equal(t, ret.Slug, post.Slug)
	assert.Equal(t, ret.EyeCatchingImg, post.EyeCatchingImg)
	assert.Equal(t, ret.Content, post.Content)
	assert.Equal(t, ret.MetaDescription, post.MetaDescription)
	assert.Equal(t, ret.IsPublic, post.IsPublic)
	assert.Equal(t, ret.CreatedAt, post.CreatedAt)
	assert.Equal(t, ret.UpdatedAt, post.UpdatedAt)
	r.AssertExpectations(t)
}

func TestPostService_Create(t *testing.T) {
	postCreatedAt, _ := time.Parse("2006-01-02 15:04:05.999999-07", "2006-01-02 15:04:05.999999-07")
	postUpdatedAt, _ := time.Parse("2006-01-02 15:04:05.999999-07", "2006-01-02 15:04:05.999999-07")

	postDto := dto.PostModel{
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
	post := entity.Post{
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

	r := new(mocks.IPostRepository)

	r.On("Create", post).Return(nil)

	s := NewPostService(r)

	err := s.Create(postDto)

	assert.NoError(t, err)
	r.AssertExpectations(t)
}

func TestPostService_Update(t *testing.T) {
	postCreatedAt, _ := time.Parse("2006-01-02 15:04:05.999999-07", "2006-01-02 15:04:05.999999-07")
	postUpdatedAt, _ := time.Parse("2006-01-02 15:04:05.999999-07", "2006-01-02 15:04:05.999999-07")

	postDto := dto.PostModel{
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
	post := entity.Post{
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

	r := new(mocks.IPostRepository)

	r.On("Update", post).Return(nil)

	s := NewPostService(r)

	err := s.Update(postDto)

	assert.NoError(t, err)
	r.AssertExpectations(t)
}

func TestPostService_Delete(t *testing.T) {
	postCreatedAt, _ := time.Parse("2006-01-02 15:04:05.999999-07", "2006-01-02 15:04:05.999999-07")
	postUpdatedAt, _ := time.Parse("2006-01-02 15:04:05.999999-07", "2006-01-02 15:04:05.999999-07")

	postDto := dto.PostModel{
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
	post := entity.Post{
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

	r := new(mocks.IPostRepository)

	r.On("Delete", post).Return(nil)

	s := NewPostService(r)

	err := s.Delete(postDto)

	assert.NoError(t, err)
	r.AssertExpectations(t)
}
