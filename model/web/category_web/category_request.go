package categoryweb

type CategoryRequest struct {
	Name      string `validate:"required,min=1,max=30" json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type CategoryUpdateRequest struct {
	CategoryId string `validate:"required" json:"category_id"`
	Name       string `validate:"required,min=1,max=30" json:"name"`
	UpdatedAt  string `json:"updated_at"`
}
