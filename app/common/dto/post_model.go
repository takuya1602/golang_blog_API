package dto

import "time"

type PostModel struct {
	Id              int       `json:"id"`
	Title           string    `json:"title"`
	Slug            string    `json:"slug"`
	EyeCatchingImg  string    `json:"eye_catching_img"`
	Content         string    `json:"content"`
	MetaDescription string    `json:"meta_description"`
	IsPublic        bool      `json:"is_public"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	CategoryId      int       `json:"category_id"`
	CategoryName    string    `json:"category_name"`
	CategorySlug    string    `json:"category_slug"`
	SubCategoryId   int       `json:"sub_category_id"`
	SubCategoryName string    `json:"sub_category_name"`
	SubCategorySlug string    `json:"sub_category_slug"`
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
