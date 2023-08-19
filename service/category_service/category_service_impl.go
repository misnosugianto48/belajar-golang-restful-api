package categoryservice

import (
	"belajar-golang-restful-api/helper"
	"belajar-golang-restful-api/model/entity"
	categoryweb "belajar-golang-restful-api/model/web/category_web"
	categoryrepository "belajar-golang-restful-api/repository/category_repo"
	"context"
	"database/sql"
	"time"

	"github.com/go-playground/validator"
)

type CategoryServiceImpl struct {
	CategoryRepository categoryrepository.CategoryRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewCategoryService(categoryRepository categoryrepository.CategoryRepository, db *sql.DB, validate *validator.Validate) CategoryService {
	return &CategoryServiceImpl{
		CategoryRepository: categoryRepository,
		DB:                 db,
		Validate:           validate,
	}
}

func (service *CategoryServiceImpl) Create(ctx context.Context, r categoryweb.CategoryRequest) categoryweb.CategoryResponse {
	//	validate request body
	err := service.Validate.Struct(r)
	helper.PanicIfError("Failed to validation request", err)

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
	category = service.CategoryRepository.Create(ctx, tx, category)

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) Update(ctx context.Context, r categoryweb.CategoryUpdateRequest) categoryweb.CategoryResponse {
	//	validate request body
	err := service.Validate.Struct(r)
	helper.PanicIfError("Failed to validation request", err)

	tx, err := service.DB.Begin()
	helper.PanicIfError("Failed starts transactions: ", err)
	defer helper.CommitOrRollback(tx)

	// check category if exists
	category, err := service.CategoryRepository.FindById(ctx, tx, r.CategoryId)
	helper.PanicIfError("category not found", err)

	updatedNow := time.Now().Format(time.RFC3339)

	//	set requestBody
	category.Name = r.Name
	category.UpdatedAt = updatedNow

	//	get repository action
	category = service.CategoryRepository.Update(ctx, tx, category)

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) Delete(ctx context.Context, categoryId string) {
	tx, err := service.DB.Begin()
	helper.PanicIfError("Failed starts transactions: ", err)
	defer helper.CommitOrRollback(tx)

	// check category if exists
	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	helper.PanicIfError("category not found: ", err)

	service.CategoryRepository.Delete(ctx, tx, category)
}

func (service *CategoryServiceImpl) FindById(ctx context.Context, categoryId string) categoryweb.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError("Failed starts transactions: ", err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	helper.PanicIfError("category not found: ", err)

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) FindAll(ctx context.Context) []categoryweb.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError("Failed starts transactions: ", err)
	defer helper.CommitOrRollback(tx)

	categories := service.CategoryRepository.FindAll(ctx, tx)

	return helper.ToCategoryResponses(categories)
}
