package dto

type SubCategoryModel struct {
	Id                 int    `json:"id"`
	Name               string `json:"name"`
	Slug               string `json:"slug"`
	ParentCategoryId   int    `json:"parent_category_id"`
	ParentCategoryName string `json:"parent_category_name"`
	ParentCategorySlug string `json:"parent_category_slug"`
}

func NewSubCategoryModel(id int, name string, slug string, parentCategoryId int, parentCategoryName string, parentCategorySlug string) (subCategoryModel SubCategoryModel) {
	subCategoryModel = SubCategoryModel{
		Id:                 id,
		Name:               name,
		Slug:               slug,
		ParentCategoryId:   parentCategoryId,
		ParentCategoryName: parentCategoryName,
		ParentCategorySlug: parentCategorySlug,
	}
	return
}
