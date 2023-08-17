package category

type CategoryRequest struct {
	Name string
}

type CategoryUpdateRequest struct {
	Id   int
	Name string
}
