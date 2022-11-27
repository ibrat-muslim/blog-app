package postgres

import (
	"github.com/ibrat-muslim/blog-app/storage/repo"
	"github.com/jmoiron/sqlx"
)

type categoryRepo struct {
	db *sqlx.DB
}

func NewCategory(db *sqlx.DB) repo.CategoryStorageI {
	return &categoryRepo{
		db: db,
	}
}

func (cr *categoryRepo) Create(category *repo.Category) (*repo.Category, error) {
	query := `
		INSERT INTO categories (
			title
		) VALUES($1)
		RETURNING id, created_at
	`

	row := cr.db.QueryRow(
		query,
		category.Title,
	)

	err := row.Scan(
		&category.ID,
		&category.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return category, nil
}

func (cr *categoryRepo) Get(id int64) (*repo.Category, error) {
	query := `
		SELECT
			id,
			title,
			created_at
		FROM categories
		WHERE id = $1
	`

	var result repo.Category

	err := cr.db.Get(&result, query, id)

	if err != nil {
		return nil, err
	}

	return &result, nil
}
