package repository

import (
	"belajar-golang-rest-api/helper"
	"belajar-golang-rest-api/model/domain"
	"context"
	"database/sql"
	"errors"
)

type CategoryRepositoryImpl struct {
}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (repository *CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "INSERT INTO categories (name) VALUES ($1);"
	result, err := tx.ExecContext(ctx, SQL, category.Name)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	category.Id = int(id)
	return category
}

func (repository *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "UPDATE categories SET name = $1 WHERE id = $2;"
	_, err := tx.ExecContext(ctx, SQL, category.Name, category.Id)
	helper.PanicIfError(err)

	return category
}

func (repository *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) {
	SQL := "DELETE FROM categories WHERE id = $1;"
	_, err := tx.ExecContext(ctx, SQL, category.Id)
	helper.PanicIfError(err)

}

func (repository *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error) {
	SQL := "SELECT id, name FROM categories WHERE id = $1;"
	rows, err := tx.QueryContext(ctx, SQL, categoryId)
	helper.PanicIfError(err)

	category := domain.Category{}
	if rows.Next() {
		err2 := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err2)
		return category, nil
	} else {
		return category, errors.New("not found")
	}
}

func (repository *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	SQL := "SELECT id, name FROM categories;"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			panic(err)
		}
	}(rows)
	var categories []domain.Category
	for rows.Next() {
		category := domain.Category{}
		err2 := rows.Scan(category)
		helper.PanicIfError(err2)
		categories = append(categories, category)
	}
	return categories
}
