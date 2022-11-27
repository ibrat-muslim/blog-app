package repo

import "time"

type Post struct {
	ID          int64      `db:"id"`
	Title       string     `db:"title"`
	Description string     `db:"description"`
	ImageUrl    *string    `db:"image_url"`
	UserID      int64      `db:"user_id"`
	CategoryID  int64      `db:"category_id"`
	CreatedAt   time.Time  `db:"created_at"`
	UpdatedAt   *time.Time `db:"updated_at"`
	ViewsCount  int32      `db:"views_count"`
}

type PostStorageI interface {
	Create(post *Post) (*Post, error)
	Get(id int64) (*Post, error)
}
