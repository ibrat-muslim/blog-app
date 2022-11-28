package postgres

import (
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

func (l *likeRepo) Get(userID, postID int64) (*repo.Like, error) {
	query := `
		SELECT
			id,
			post_id,
			user_id,
			status
		FROM likes
		WHERE user_id = $1 AND post_id = $2
	`

	var result repo.Like

	err := l.db.Get(&result, query, userID, postID)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (l *likeRepo) GetLikesDislikesCount(postID int64) (*repo.LikesDislikesCountsResult, error) {
	var result repo.LikesDislikesCountsResult

	query := `
		SELECT
			COUNT(1) FILTER (WHERE status=true) as likes_count,
			COUNT(1) FILTER (WHERE status=false) as dislikes_count 
		FROM likes
		WHERE post_id = $1
		`

	err := l.db.Get(&result, query, postID)

	if err != nil {
		return nil, err
	}

	return &result, nil
}
