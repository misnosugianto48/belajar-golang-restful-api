package category_service

import (
	"belajar-golang-restful-api/helper"
	"belajar-golang-restful-api/model/entity"
	"belajar-golang-restful-api/model/web/category_web"
	"belajar-golang-restful-api/repository/category_repo"
	"context"
	"database/sql"
	"time"
)

type CategoryServiceImpl struct {
	categoryRepository category_repo.CategoryRepository
	DB                 *sql.DB
}

func (service *CategoryServiceImpl) Create(ctx context.Context, r category_web.CategoryRequest) category_web.CategoryResponse {
	//	start transactions database
	tx, err := service.DB.Begin()
	helper.PanicIfError("Failed starts transactions: ", err)

	defer helper.CommitOrRollback(tx)

	//	set requestBody
	category := entity.Category{
		Name:      r.Name,
		CreatedAt: r.CreatedAt,
		UpdatedAt: r.UpdatedAt,
	}

	//	get repository action
	category = service.categoryRepository.Create(ctx, tx, category)

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) Update(ctx context.Context, r category_web.CategoryUpdateRequest) category_web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError("Failed starts transactions: ", err)
	defer helper.CommitOrRollback(tx)

	// check category if exists
	category, err := service.categoryRepository.FindById(ctx, tx, r.CategoryId)
	helper.PanicIfError("category not found", err)

	updatedNow := time.Now().Format(time.RFC3339)

	//	set requestBody
	category.Name = r.Name
	category.UpdatedAt = updatedNow

	//	get repository action
	category = service.categoryRepository.Update(ctx, tx, category)

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) Delete(ctx context.Context, categoryId string) {
	tx, err := service.DB.Begin()
	helper.PanicIfError("Failed starts transactions: ", err)
	defer helper.CommitOrRollback(tx)

	// check category if exists
	category, err := service.categoryRepository.FindById(ctx, tx, categoryId)
	helper.PanicIfError("category not found", err)

	service.categoryRepository.Delete(ctx, tx, category)
}

func (service *CategoryServiceImpl) FIndById(ctx context.Context, categoryId string) category_web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError("Failed starts transactions: ", err)
	defer helper.CommitOrRollback(tx)

	category, err := service.categoryRepository.FindById(ctx, tx, categoryId)
	helper.PanicIfError("category not found", err)

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) FindAll(ctx context.Context) []category_web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError("Failed starts transactions: ", err)
	defer helper.CommitOrRollback(tx)

	categories := service.categoryRepository.FindAll(ctx, tx)

	return helper.ToCategoryResponses(categories)
}
