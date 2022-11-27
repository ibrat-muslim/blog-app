package models

type Like struct {
	ID     int64 `db:"id"`
	PostID int64 `db:"post_id"`
	UserID int64 `db:"user_id"`
	Status bool  `db:"status"`
}

type CreateLikeRequest struct {
	PostID int64 `db:"post_id"`
	UserID int64 `db:"user_id"`
	Status bool  `db:"status"`
}

type GetLikesResult struct {
	Likes []*Like `json:"likes"`
	Count int32   `json:"count"`
}
