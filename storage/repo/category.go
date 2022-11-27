package repo

import "time"

type Category struct {
	ID        int64     `db:"id"`
	Title     string    `db:"title"`
	CreatedAt time.Time `db:"created_at"`
}

type CategoryStorageI interface {
	Create(category *Category) (*Category, error)
	Get(id int64) (*Category, error)
}