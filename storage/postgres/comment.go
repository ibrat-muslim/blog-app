package postgres

import (
	"database/sql"
	"fmt"

	"github.com/ibrat-muslim/blog-app/storage/repo"
	"github.com/jmoiron/sqlx"
)

type commentRepo struct {
	db *sqlx.DB
}

func NewComment(db *sqlx.DB) repo.CommentStorageI {
	return &commentRepo{
		db: db,
	}
}

func (cmr *commentRepo) Create(comment *repo.Comment) (*repo.Comment, error) {
	query := `
		INSERT INTO comments (
			post_id,
			user_id,
			description
		) VALUES($1, $2, $3)
		RETURNING id, created_at
	`

	row := cmr.db.QueryRow(
		query,
		comment.PostID,
		comment.UserID,
		comment.Description,
	)

	err := row.Scan(
		&comment.ID,
		&comment.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return comment, nil
}

func (cmr *commentRepo) Get(id int64) (*repo.Comment, error) {
	query := `
		SELECT
			id,
			post_id,
			user_id,
			description,
			created_at,
			updated_at
		FROM comments
		WHERE id = $1
	`

	var result repo.Comment

	err := cmr.db.Get(&result, query, id)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (cmr *commentRepo) GetAll(params *repo.GetCommentsParams) (*repo.GetCommentsResult, error) {
	result := repo.GetCommentsResult{
		Comments: make([]*repo.Comment, 0),
		Count: 0,
	}

	offset := (params.Page - 1) * params.Limit

	limit := fmt.Sprintf(" LIMIT %d OFFSET %d ", params.Limit, offset)

	query := `
		SELECT
			id,
			post_id,
			user_id,
			description,
			created_at,
			updated_at
		FROM comments
		ORDER BY created_at DESC
		` + limit
	
	err := cmr.db.Select(&result.Comments, query)

	if err != nil {
		return nil, err
	}

	queryCount := `SELECT count(1) FROM comments` //TODO

	err = cmr.db.Get(&result.Count, queryCount)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (cmr *commentRepo) Update(comment *repo.Comment) error {
	query := `
		UPDATE comments SET
			post_id = $1,
			user_id = $2,
			description = $3,
			updated_at = $4
		WHERE id = $5
	`

	result, err := cmr.db.Exec(
		query,
		comment.PostID,
		comment.UserID,
		comment.Description,
		comment.UpdatedAt,
		comment.ID,
	)

	if err != nil {
		return err
	}

	rowsCount, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if rowsCount == 0 {
		return sql.ErrNoRows
	}	

	return nil
}

func (cmr *commentRepo) Delete(id int64) error {
	query := `DELETE FROM comments WHERE id = $1`

	result, err := cmr.db.Exec(query, id)

	if err != nil {
		return err
	}

	rowsCount, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if rowsCount == 0 {
		return sql.ErrNoRows
	}

	return nil
}