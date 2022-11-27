package postgres

import (
	"database/sql"
	"fmt"

	"github.com/ibrat-muslim/blog-app/storage/repo"
	"github.com/jmoiron/sqlx"
)

type postRepo struct {
	db *sqlx.DB
}

func NewPost(db *sqlx.DB) repo.PostStorageI {
	return &postRepo{
		db: db,
	}
}

func (pr *postRepo) Create(post *repo.Post) (*repo.Post, error) {
	query := `
		INSERT INTO posts (
			title,
			description,
			image_url,
			user_id,
			category_id
		) VALUES($1, $2, $3, $4, $5)
		RETURNING id, created_at
	`

	row := pr.db.QueryRow(
		query,
		post.Title,
		post.Description,
		post.ImageUrl,
		post.UserID,
		post.CategoryID,
	)

	err := row.Scan(
		&post.ID,
		&post.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return post, nil
}

func (pr *postRepo) Get(id int64) (*repo.Post, error) {
	queryView := `UPDATE posts SET views_count = views_count + 1 WHERE id = $1`
	
	_, err := pr.db.Exec(queryView, id)
	if err != nil {
		return nil, err
	}

	query := `
		SELECT
			id,
			title,
			description,
			image_url,
			user_id,
			category_id,
			created_at,
			updated_at,
			views_count
		FROM posts
		WHERE id = $1
	`

	var result repo.Post

	err = pr.db.Get(&result, query, id)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (pr *postRepo) GetAll(params *repo.GetPostsParams) (*repo.GetPostsResult, error) {
	result := repo.GetPostsResult{
		Posts: make([]*repo.Post, 0),
		Count: 0,
	}

	offset := (params.Page - 1) * params.Limit

	limit := fmt.Sprintf(" LIMIT %d OFFSET %d ", params.Limit, offset)

	filter := ""

	if params.Title != "" {
		str := "%" + params.Title + "%"
		filter += fmt.Sprintf(`
				WHERE title ILIKE '%s'`, str,
		)	
	}

	query := `
		SELECT
			id,
			title,
			description,
			image_url,
			user_id,
			category_id,
			created_at,
			updated_at,
			views_count
		FROM posts
		` + filter + `
		ORDER BY created_at DESC
		` + limit
	
	err := pr.db.Select(&result.Posts, query)

	if err != nil {
		return nil, err
	}

	queryCount := `SELECT count(1) FROM posts ` + filter

	err = pr.db.Get(&result.Count, queryCount)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (pr *postRepo) Update(post * repo.Post) error {
	query := `
		UPDATE posts SET
			title = $1,
			description = $2,
			image_url = $3,
			user_id = $4,
			category_id = $5,
			updated_at = $6
		WHERE id = $7
	`

	result, err := pr.db.Exec(
		query,
		post.Title,
		post.Description,
		post.ImageUrl,
		post.UserID,
		post.CategoryID,
		post.UpdatedAt,
		post.ID,
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

func (pr *postRepo) Delete(id int64) error {
	query := `DELETE FROM posts WHERE id = $1`

	result, err := pr.db.Exec(query, id)

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