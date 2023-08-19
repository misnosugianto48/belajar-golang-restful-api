package app

import (
	"belajar-golang-restful-api/helper"
	"database/sql"
	"time"
)

func NewDatabase() *sql.DB {
	database, err := sql.Open("mysql", "root@tcp(localhost:3306)/belajar_golang_restful_api")
	helper.PanicIfError("failed to connected database: ", err)

	database.SetMaxIdleConns(5)
	database.SetMaxOpenConns(20)
	database.SetConnMaxLifetime(60 * time.Minute)
	database.SetConnMaxIdleTime(10 * time.Minute)

	return database
}
