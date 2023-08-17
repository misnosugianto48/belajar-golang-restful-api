package category

import (
	"belajar-golang-restful-api/model/web/category"
	"context"
)

type CategoryService interface {
	Create(ctx context.Context, r category.CategoryRequest) category.CategoryResponse
	Update(ctx context.Context, r category.CategoryUpdateRequest) category.CategoryResponse
	Delete(ctx context.Context, categoryId int)
	FindById(ctx context.Context, categoryId int) category.CategoryResponse
	FindAll(ctx context.Context) []category.CategoryResponse
}
