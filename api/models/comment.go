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

type GetCommentsParams struct {
	Limit int32 `json:"limit" binding:"required" default:"10"`
	Page  int32 `json:"page" binding:"required" default:"1"`
}

type GetCommentsResponse struct {
	Comments []*Comment `json:"comments"`
	Count    int32      `json:"count"`
}
