package category_web

type CategoryRequest struct {
	Name      string `validate:"required,min=1,max=30" json:"name"`
	CreatedAt string `validate:"required" json:"created_at"`
	UpdatedAt string `validate:"required" json:"updated_at"`
}

type CategoryUpdateRequest struct {
	CategoryId string `validate:"required" json:"category_id"`
	Name       string `validate:"required,min=1,max=30" json:"name"`
	UpdatedAt  string `validate:"required" json:"updated_at"`
}
