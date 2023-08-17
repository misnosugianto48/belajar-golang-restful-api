package category

import (
	"belajar-golang-restful-api/helper"
	"belajar-golang-restful-api/model/entity"
	"context"
	"database/sql"
	"errors"
)

type CategoryRepositoryImpl struct {
}

func (repository *CategoryRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, category entity.Category) entity.Category {
	query := "INSERT INTO category (name) VALUES (?)"
	result, err := tx.ExecContext(ctx, query, category.Name)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	category.Id = int(id)
	return category
}

func (repository *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category entity.Category) entity.Category {
	query := "UPDATE category SET name = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, query, category.Id)
	helper.PanicIfError(err)

	return category
}

func (repository *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category entity.Category) {
	query := "DELETE FROM category WHERE id = ?"
	_, err := tx.ExecContext(ctx, query, category.Id)
	helper.PanicIfError(err)

}

func (repository *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId int) (entity.Category, error) {
	query := "SELECT id, name FROM category WHERE id = ?"
	rows, err := tx.QueryContext(ctx, query, categoryId)
	helper.PanicIfError(err)

	category := entity.Category{}
	if rows.Next() {
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
		return category, nil
	} else {
		return category, errors.New("category not found")
	}
}

func (repository *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []entity.Category {
	query := "SELECT id, name FROM category"
	rows, err := tx.QueryContext(ctx, query)
	helper.PanicIfError(err)

	var categories []entity.Category
	for rows.Next() {
		category := entity.Category{}
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
		categories = append(categories, category)
	}
	return categories
}
