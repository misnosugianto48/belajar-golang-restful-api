package category_service

import (
	"belajar-golang-restful-api/model/web/category_web"
	"belajar-golang-restful-api/repository/category_repo"
	"context"
	"database/sql"
)

type CategoryServiceImpl struct {
	categoryRepository category_repo.CategoryRepository
	DB                 *sql.DB
}

func (service *CategoryServiceImpl) Create(ctx context.Context, r category_web.CategoryRequest) {

}
