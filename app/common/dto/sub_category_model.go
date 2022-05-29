package dto

type SubCategoryModel struct {
	Id                 int    `json:"id"`
	Name               string `json:"name"`
	Slug               string `json:"slug"`
	ParentCategoryName string `json:"parent_category"`
}

func NewSubCategoryModel(id int, name string, slug string, parentCategory string) (subCategoryModel SubCategoryModel) {
	subCategoryModel = SubCategoryModel{
		Id:                 id,
		Name:               name,
		Slug:               slug,
		ParentCategoryName: parentCategory,
	}
	return
}
