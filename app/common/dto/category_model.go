package dto

type CategoryModel struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

func NewCategoryModel(id int, name string, slug string) (categoryModel CategoryModel) {
	categoryModel = CategoryModel{
		Id:   id,
		Name: name,
		Slug: slug,
	}
	return
}
