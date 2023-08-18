package helper

import (
	"belajar-golang-restful-api/model/entity"
	"belajar-golang-restful-api/model/web/category_web"
)

func ToCategoryResponse(category entity.Category) category_web.CategoryResponse {
	return category_web.CategoryResponse{
		CategoryId: category.CategoryId,
		Name:       category.Name,
		CreatedAt:  category.CreatedAt,
		UpdatedAt:  category.UpdatedAt,
	}
}

func ToCategoryResponses(categories []entity.Category) []category_web.CategoryResponse {
	var categoryResponses []category_web.CategoryResponse

	for _, category := range categories {
		categoryResponses = append(categoryResponses, ToCategoryResponse(category))
	}
	return categoryResponses
}
