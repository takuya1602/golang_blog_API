package entity

type Category struct {
	Id   int
	Name string
	Slug string
}

func NewCategory(id int, name string, slug string) (category Category) {
	category = Category{
		Id:   id,
		Name: name,
		Slug: slug,
	}
	return
}
