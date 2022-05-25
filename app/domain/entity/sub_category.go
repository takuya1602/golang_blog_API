package entity

type SubCategory struct {
	Id               int
	Name             string
	Slug             string
	ParentCategoryId int
}

func NewSubCategory(id int, name string, slug string, parentCategoryId int) (subCategory SubCategory) {
	subCategory = SubCategory{
		Id:               id,
		Name:             name,
		Slug:             slug,
		ParentCategoryId: parentCategoryId,
	}
	return
}
