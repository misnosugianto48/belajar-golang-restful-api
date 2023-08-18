package categoryservice

import (
	categoryweb "belajar-golang-restful-api/model/web/category_web"
	"context"
)

type CategoryService interface {
	Create(ctx context.Context, r categoryweb.CategoryRequest) categoryweb.CategoryResponse
	Update(ctx context.Context, r categoryweb.CategoryUpdateRequest) categoryweb.CategoryResponse
	Delete(ctx context.Context, categoryId string)
	FindById(ctx context.Context, categoryId string) categoryweb.CategoryResponse
	FindAll(ctx context.Context) []categoryweb.CategoryResponse
}
