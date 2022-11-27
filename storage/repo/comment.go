package repo

import "time"

type Comment struct {
	ID          int64      `db:"id"`
	PostID      int64      `db:"post_id"`
	UserID      int64      `db:"user_id"`
	Description string     `db:"description"`
	CreatedAt   time.Time  `db:"created_at"`
	UpdatedAt   *time.Time `db:"updated_at"`
}

type GetCommentsParams struct {
	Page  int32
	Limit int32
}

type GetCommentsResult struct {
	Comments []*Comment
	Count    int32
}

type CommentStorageI interface {
	Create(comment *Comment) (*Comment, error)
	Get(id int64) (*Comment, error)
	GetAll(params *GetCommentsParams) (*GetCommentsResult, error)
	Update(comment *Comment) error
	Delete(id int64) error
}
