package categoryrepository

import (
	"belajar-golang-restful-api/model/entity"
	"context"
	"database/sql"
)

// TODO: Create contract to struct category
type CategoryRepository interface {
	Create(ctx context.Context, tx *sql.Tx, category entity.Category) entity.Category
	Update(ctx context.Context, tx *sql.Tx, category entity.Category) entity.Category
	Delete(ctx context.Context, tx *sql.Tx, category entity.Category)
	FindById(ctx context.Context, tx *sql.Tx, categoryId string) (entity.Category, error)
	FindAll(ctx context.Context, tx *sql.Tx) []entity.Category
}
