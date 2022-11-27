package models

import "time"

type Comment struct {
	ID          int64      `json:"id"`
	PostID      int64      `json:"post_id"`
	UserID      int64      `json:"user_id"`
	Description string     `json:"description"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}

type CreateCommentRequest struct {
	PostID      int64  `json:"post_id"`
	UserID      int64  `json:"user_id"`
	Description string `json:"description"`
}

type GetCommentsResult struct {
	Comments []*Comment `json:"comments"`
	Count    int32      `json:"count"`
}
