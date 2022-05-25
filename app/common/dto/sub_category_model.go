package dto

type SubCategoryModel struct {
	Id               int    `json:"id"`
	Name             string `json:"name"`
	Slug             string `json:"slug"`
	ParentCategoryId int    `json:"parent_category_id"`
}

func NewSubCategoryModel(id int, name string, slug string, parentCategoryId int) (subCategoryModel SubCategoryModel) {
	subCategoryModel = SubCategoryModel{
		Id:               id,
		Name:             name,
		Slug:             slug,
		ParentCategoryId: parentCategoryId,
	}
	return
}
