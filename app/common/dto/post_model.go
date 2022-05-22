package dto

import "time"

type PostModel struct {
	Id              int       `json:"id"`
	CategoryId      int       `json:"category_id"`
	SubCategoryId   int       `json:"sub_category_id"`
	Title           string    `json:"title"`
	Slug            string    `json:"slug"`
	EyeCatchingImg  string    `json:"eye_catching_img"`
	Content         string    `json:"content"`
	MetaDescription string    `json:"meta_description"`
	IsPublic        bool      `json:"is_public"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

func NewPostModel(id int, categoryId int, subCategoryId int, title string, slug string, eyeCatchingImg string, content string, metaDescription string, isPublic bool, createdAt time.Time, updatedAt time.Time) (postModel PostModel) {
	postModel = PostModel{
		Id:              id,
		CategoryId:      categoryId,
		SubCategoryId:   subCategoryId,
		Title:           title,
		Slug:            slug,
		EyeCatchingImg:  eyeCatchingImg,
		Content:         content,
		MetaDescription: metaDescription,
		IsPublic:        isPublic,
		CreatedAt:       createdAt,
		UpdatedAt:       updatedAt,
	}
	return
}
