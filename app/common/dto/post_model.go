package dto

import "time"

type PostModel struct {
	Id              int       `json:"id"`
	CategoryName    string    `json:"category"`
	SubCategoryName string    `json:"sub_category"`
	Title           string    `json:"title"`
	Slug            string    `json:"slug"`
	EyeCatchingImg  string    `json:"eye_catching_img"`
	Content         string    `json:"content"`
	MetaDescription string    `json:"meta_description"`
	IsPublic        bool      `json:"is_public"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

func NewPostModel(id int, categoryName string, subCategoryName string, title string, slug string, eyeCatchingImg string, content string, metaDescription string, isPublic bool, createdAt time.Time, updatedAt time.Time) (postModel PostModel) {
	postModel = PostModel{
		Id:              id,
		CategoryName:    categoryName,
		SubCategoryName: subCategoryName,
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
