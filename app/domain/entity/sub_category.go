package entity

type SubCategory struct {
	Id                 int
	Name               string
	Slug               string
	ParentCategoryId   int
	ParentCategoryName string
	ParentCategorySlug string
}

func NewSubCategory(id int, name string, slug string, parentCategoryId int, parentCategoryName string, parentCategorySlug string) (subCategory SubCategory) {
	subCategory = SubCategory{
		Id:                 id,
		Name:               name,
		Slug:               slug,
		ParentCategoryId:   parentCategoryId,
		ParentCategoryName: parentCategoryName,
		ParentCategorySlug: parentCategorySlug,
	}
	return
}
