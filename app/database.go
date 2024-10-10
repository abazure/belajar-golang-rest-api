package app

import (
	"belajar-golang-rest-api/helper"
	"database/sql"
	_ "github.com/jackc/pgx/stdlib"
	"time"
)

func NewDB() *sql.DB {
	db, err := sql.Open("pgx", "postgres://postgres:@localhost:5432/belajar_golang_rest-api")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(10)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)
	return db
}
