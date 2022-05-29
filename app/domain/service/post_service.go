package service

import (
	"backend/app/common/dto"
	"backend/app/domain/entity"
	"backend/app/domain/repository"
)

type IPostService interface {
	GetPosts(map[string][]string) ([]dto.PostModel, error)
	GetPostBySlug(string) (dto.PostModel, error)
	Create(dto.PostModel) error
	Update(dto.PostModel) error
	Delete(dto.PostModel) error
}

type PostService struct {
	repository.IPostRepository
}

func NewPostService(repo repository.IPostRepository) (postService IPostService) {
	postService = &PostService{repo}
	return
}

func (s *PostService) convertToDtoFromEntity(post entity.Post) (postDto dto.PostModel) {
	postDto = dto.PostModel{
		Id:              post.Id,
		Title:           post.Title,
		Slug:            post.Slug,
		EyeCatchingImg:  post.EyeCatchingImg,
		Content:         post.Content,
		MetaDescription: post.MetaDescription,
		IsPublic:        post.IsPublic,
		CreatedAt:       post.CreatedAt,
		UpdatedAt:       post.UpdatedAt,
		CategoryId:      post.CategoryId,
		CategoryName:    post.CategoryName,
		CategorySlug:    post.CategorySlug,
		SubCategoryId:   post.SubCategoryId,
		SubCategoryName: post.SubCategoryName,
		SubCategorySlug: post.SubCategorySlug,
	}
	return
}

func (s *PostService) convertToDtosFromEntities(posts []entity.Post) (postDtos []dto.PostModel) {
	for _, post := range posts {
		postDto := s.convertToDtoFromEntity(post)
		postDtos = append(postDtos, postDto)
	}
	return
}

func (s *PostService) convertToEntityFromDto(postDto dto.PostModel) (post entity.Post) {
	post = entity.Post{
		Id:              postDto.Id,
		Title:           postDto.Title,
		Slug:            postDto.Slug,
		EyeCatchingImg:  postDto.EyeCatchingImg,
		Content:         postDto.Content,
		MetaDescription: postDto.MetaDescription,
		IsPublic:        postDto.IsPublic,
		CreatedAt:       postDto.CreatedAt,
		UpdatedAt:       postDto.UpdatedAt,
		CategoryId:      postDto.CategoryId,
		CategoryName:    postDto.CategoryName,
		CategorySlug:    postDto.CategorySlug,
		SubCategoryId:   postDto.SubCategoryId,
		SubCategoryName: postDto.SubCategoryName,
		SubCategorySlug: postDto.SubCategorySlug,
	}
	return
}

func (s *PostService) convertEntitiesFromDtos(postDtos []dto.PostModel) (posts []entity.Post) {
	for _, postDto := range postDtos {
		post := s.convertToEntityFromDto(postDto)
		posts = append(posts, post)
	}
	return
}

func (s *PostService) GetPosts(queryParams map[string][]string) (postDtos []dto.PostModel, err error) {
	posts, err := s.IPostRepository.GetPosts(queryParams)
	if err != nil {
		return
	}
	postDtos = s.convertToDtosFromEntities(posts)
	return
}

func (s *PostService) GetPostBySlug(slug string) (postDto dto.PostModel, err error) {
	post, err := s.IPostRepository.GetPostBySlug(slug)
	if err != nil {
		return
	}
	postDto = s.convertToDtoFromEntity(post)
	return
}

func (s *PostService) Create(postDto dto.PostModel) (err error) {
	post := s.convertToEntityFromDto(postDto)
	err = s.IPostRepository.Create(post)
	return
}

func (s *PostService) Update(postDto dto.PostModel) (err error) {
	post := s.convertToEntityFromDto(postDto)
	err = s.IPostRepository.Update(post)
	return
}

func (s *PostService) Delete(postDto dto.PostModel) (err error) {
	post := s.convertToEntityFromDto(postDto)
	err = s.IPostRepository.Delete(post)
	return
}
