package postgres

import (
	"database/sql"
	"fmt"

	"github.com/ibrat-muslim/blog-app/storage/repo"
	"github.com/jmoiron/sqlx"
)

type likeRepo struct {
	db *sqlx.DB
}

func NewLike(db *sqlx.DB) repo.LikeStorageI {
	return &likeRepo{
		db: db,
	}
}

func (l *likeRepo) Create(like *repo.Like) (*repo.Like, error) {
	query := `
		INSERT INTO likes (
			post_id,
			user_id,
			status
		) VALUES($1, $2, $3)
		RETURNING id
	`

	row := l.db.QueryRow(
		query,
		like.PostID,
		like.UserID,
		like.Status,
	)

	err := row.Scan(
		&like.ID,
	)

	if err != nil {
		return nil, err
	}

	return like, nil
}

func (l *likeRepo) Get(id int64) (*repo.Like, error) {
	query := `
		SELECT
			id,
			post_id,
			user_id,
			status
		FROM likes
		WHERE id = $1
	`

	var result repo.Like

	err := l.db.Get(&result, query, id)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (l *likeRepo) GetAll(params *repo.GetLikesParams) (*repo.GetLikesResult, error) {
	result := repo.GetLikesResult{
		Likes: make([]*repo.Like, 0),
		Count: 0,
	}

	offset := (params.Page - 1) * params.Limit

	limit := fmt.Sprintf(" LIMIT %d OFFSET %d ", params.Limit, offset)

	query := `
		SELECT
			id,
			post_id,
			user_id,
			status
		FROM likes
		ORDER BY id DESC
		` + limit
	
	err := l.db.Select(&result.Likes, query)

	if err != nil {
		return nil, err
	}

	queryCount := `SELECT count(1) FROM likes` //TODO

	err = l.db.Get(&result.Count, queryCount)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (l *likeRepo) Update(like *repo.Like) error {
	query := `
		UPDATE likes SET
			post_id = $1,
			user_id = $2,
			status = $3
		WHERE id  = $4
	`

	result, err := l.db.Exec(
		query,
		like.PostID,
		like.UserID,
		like.Status,
		like.ID,
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

func (l *likeRepo) Delete(id int64) error {
	query := `DELETE FROM likes WHERE id = $1`

	result, err := l.db.Exec(query, id)

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