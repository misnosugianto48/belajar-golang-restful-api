package categoryrepository

import (
	"belajar-golang-restful-api/helper"
	"belajar-golang-restful-api/model/entity"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/google/uuid"
)

type CategoryRepositoryImpl struct {
}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func generateCategoryID() string {
	newUUID := uuid.New()
	return fmt.Sprintf("category-%s", newUUID.String()[:12])
}

func (repository *CategoryRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, category entity.Category) entity.Category {
	query := "INSERT INTO category (category_id, name, created_at, updated_at) VALUES (?, ?, ?, ?)"

	// Create New UUID
	categoryId := generateCategoryID()

	// get the time and string conversion to format RFC3339
	createdNow := time.Now().Format(time.RFC3339)
	updatedNow := time.Now().Format(time.RFC3339)

	result, err := tx.ExecContext(ctx, query, categoryId, category.Name, createdNow, updatedNow)
	helper.PanicIfError("Failed to execute SQL query : ", err)

	rowsAffected, err := result.RowsAffected()
	helper.PanicIfError("Failed to get RowsAffected:", err)

	if rowsAffected != 1 {
		log.Info("Expected 1 row affected, but got:", rowsAffected)
	} else {
		log.Info("Category created successfully")
	}

	category.CategoryId = categoryId
	category.CreatedAt = createdNow
	category.UpdatedAt = updatedNow
	return category
}

func (repository *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category entity.Category) entity.Category {
	query := "UPDATE `category` SET `name` = ?, `updated_at` = ? WHERE `category_id` = ?"

	//	set time updated
	updatedNow := time.Now().Format(time.RFC3339)

	_, err := tx.ExecContext(ctx, query, category.Name, updatedNow, category.CategoryId)
	helper.PanicIfError("Failed to execute SQL query : ", err)

	category.UpdatedAt = updatedNow
	return category
}

func (repository *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category entity.Category) {
	query := "DELETE FROM category WHERE `category_id` = ?"
	_, err := tx.ExecContext(ctx, query, category.CategoryId)
	helper.PanicIfError("Failed to execute SQL query : ", err)

}

func (repository *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId string) (entity.Category, error) {
	query := "SELECT * FROM `category` WHERE `category_id` =  ?"
	rows, err := tx.QueryContext(ctx, query, categoryId)
	helper.PanicIfError("Failed to execute SQL query : ", err)
	//	verify rows close to handle invalid connection
	defer rows.Close()

	category := entity.Category{}
	if rows.Next() {
		err := rows.Scan(&category.CategoryId, &category.Name, &category.CreatedAt, &category.UpdatedAt)
		helper.PanicIfError("Failed to read columns from the database: ", err)
		return category, nil
	} else {
		return category, errors.New("category not found")
	}
}

func (repository *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []entity.Category {
	query := "SELECT * FROM `category`"
	rows, err := tx.QueryContext(ctx, query)
	helper.PanicIfError("Failed to execute SQL query : ", err)
	defer rows.Close()

	var categories []entity.Category
	for rows.Next() {
		category := entity.Category{}
		err := rows.Scan(&category.CategoryId, &category.Name, &category.CreatedAt, &category.UpdatedAt)
		helper.PanicIfError("Failed to read columns from the database: ", err)
		categories = append(categories, category)
	}
	return categories
}
