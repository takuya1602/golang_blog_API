package service

import (
	"backend/app/common/dto"
	"backend/app/domain/entity"
	"backend/app/domain/repository"
)

type IPostService interface {
	GetAll() ([]dto.PostModel, error)
	GetBySlug(string) (dto.PostModel, error)
	GetWithCategoryQuery(string) ([]dto.PostModel, error)
	GetWithSubCategoryQuery(string) ([]dto.PostModel, error)
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
	categoryName := s.IPostRepository.GetNameFromCategoryId(post.CategoryId)
	subCategoryName := s.IPostRepository.GetNameFromSubCategoryId(post.SubCategoryId)
	postDto = dto.NewPostModel(post.Id, categoryName, subCategoryName, post.Title, post.Slug, post.EyeCatchingImg, post.Content, post.MetaDescription, post.IsPublic, post.CreatedAt, post.UpdatedAt)
	return
}

func (s *PostService) convertToDtosFromEntities(posts []entity.Post) (postDtos []dto.PostModel) {
	for _, post := range posts {
		categoryName := s.IPostRepository.GetNameFromCategoryId(post.CategoryId)
		subCategoryName := s.IPostRepository.GetNameFromSubCategoryId(post.SubCategoryId)
		postDto := dto.NewPostModel(post.Id, categoryName, subCategoryName, post.Title, post.Slug, post.EyeCatchingImg, post.Content, post.MetaDescription, post.IsPublic, post.CreatedAt, post.UpdatedAt)
		postDtos = append(postDtos, postDto)
	}
	return
}

func (s *PostService) convertToEntityFromDto(postDto dto.PostModel) (post entity.Post) {
	categoryId := s.IPostRepository.GetIdFromCategoryName(postDto.CategoryName)
	subCategoryId := s.IPostRepository.GetIdFromSubCategoryName(postDto.SubCategoryName)
	post = entity.NewPost(postDto.Id, categoryId, subCategoryId, postDto.Title, postDto.Slug, postDto.EyeCatchingImg, postDto.Content, postDto.MetaDescription, postDto.IsPublic, postDto.CreatedAt, postDto.UpdatedAt)
	return
}

func (s *PostService) convertEntitiesFromDtos(postDtos []dto.PostModel) (posts []entity.Post) {
	for _, postDto := range postDtos {
		categoryId := s.IPostRepository.GetIdFromCategoryName(postDto.CategoryName)
		subCategoryId := s.IPostRepository.GetIdFromSubCategoryName(postDto.SubCategoryName)
		post := entity.NewPost(postDto.Id, categoryId, subCategoryId, postDto.Title, postDto.Slug, postDto.EyeCatchingImg, postDto.Content, postDto.MetaDescription, postDto.IsPublic, postDto.CreatedAt, postDto.UpdatedAt)
		posts = append(posts, post)
	}
	return
}

func (s *PostService) GetAll() (postDtos []dto.PostModel, err error) {
	posts, err := s.IPostRepository.GetAll()
	if err != nil {
		return
	}
	postDtos = s.convertToDtosFromEntities(posts)
	return
}

func (s *PostService) GetWithCategoryQuery(categoryName string) (postDtos []dto.PostModel, err error) {
	posts, err := s.IPostRepository.GetFilterCategory(categoryName)
	if err != nil {
		return
	}
	postDtos = s.convertToDtosFromEntities(posts)
	return
}

func (s *PostService) GetWithSubCategoryQuery(subCategoryName string) (postDtos []dto.PostModel, err error) {
	posts, err := s.IPostRepository.GetFilterSubCategory(subCategoryName)
	if err != nil {
		return
	}
	postDtos = s.convertToDtosFromEntities(posts)
	return
}

func (s *PostService) GetBySlug(slug string) (postDto dto.PostModel, err error) {
	post, err := s.IPostRepository.GetBySlug(slug)
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
