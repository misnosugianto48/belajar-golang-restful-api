package category_web

type CategoryRequest struct {
	Name      string
	CreatedAt string
	UpdatedAt string
}

type CategoryUpdateRequest struct {
	CategoryId string
	Name       string
	UpdatedAt  string
}
