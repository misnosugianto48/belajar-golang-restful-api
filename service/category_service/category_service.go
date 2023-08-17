package category_service

import (
	"belajar-golang-restful-api/model/web/category_web"
	"context"
)

type CategoryService interface {
	Create(ctx context.Context, r category_web.CategoryRequest) category_web.CategoryResponse
	Update(ctx context.Context, r category_web.CategoryUpdateRequest) category_web.CategoryResponse
	Delete(ctx context.Context, categoryId int)
	FindById(ctx context.Context, categoryId int) category_web.CategoryResponse
	FindAll(ctx context.Context) []category_web.CategoryResponse
}
