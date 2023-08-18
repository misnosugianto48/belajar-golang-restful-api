package helper

import (
	"belajar-golang-restful-api/model/entity"
	categoryweb "belajar-golang-restful-api/model/web/category_web"
)

func ToCategoryResponse(category entity.Category) categoryweb.CategoryResponse {
	return categoryweb.CategoryResponse{
		CategoryId: category.CategoryId,
		Name:       category.Name,
		CreatedAt:  category.CreatedAt,
		UpdatedAt:  category.UpdatedAt,
	}
}

func ToCategoryResponses(categories []entity.Category) []categoryweb.CategoryResponse {
	var categoryResponses []categoryweb.CategoryResponse

	for _, category := range categories {
		categoryResponses = append(categoryResponses, ToCategoryResponse(category))
	}
	return categoryResponses
}
