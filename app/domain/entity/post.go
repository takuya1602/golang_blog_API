package entity

import "time"

type Post struct {
	Id              int
	Title           string
	Slug            string
	EyeCatchingImg  string
	Content         string
	MetaDescription string
	IsPublic        bool
	CreatedAt       time.Time
	UpdatedAt       time.Time
	CategoryId      int
	CategoryName    string
	CategorySlug    string
	SubCategoryId   int
	SubCategoryName string
	SubCategorySlug string
}

func NewPost(id int, categoryId int, subCategoryId int, title string, slug string, eyeCatchingImg string, content string, metaDescription string, isPublic bool, createdAt time.Time, updatedAt time.Time) (post Post) {
	post = Post{
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
