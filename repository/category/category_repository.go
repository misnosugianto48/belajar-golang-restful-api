package category

import (
	"belajar-golang-restful-api/model/entity"
	"context"
	"database/sql"
)

// TODO: Create contract to category
type CategoryRepository interface {
	Create(ctx context.Context, tx sql.Tx, category entity.Category) entity.Category
	Update(ctx context.Context, tx sql.Tx, category entity.Category) entity.Category
	Delete(ctx context.Context, tx sql.Tx, category entity.Category)
	FindById(ctx context.Context, tx sql.Tx, categoryId int) entity.Category
	FindAll(ctx context.Context, tx sql.Tx) []entity.Category
}
