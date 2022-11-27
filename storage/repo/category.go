package repo

import "time"

type Category struct {
	ID        int64     `db:"id"`
	Title     string    `db:"title"`
	CreatedAt time.Time `db:"created_at"`
}

type GetCategoriesParams struct {
	Page  int32
	Limit int32
	Title string
}

type GetCategoriesResult struct {
	Categories []*Category
	Count      int32
}

type CategoryStorageI interface {
	Create(category *Category) (*Category, error)
	Get(id int64) (*Category, error)
	GetAll(params *GetCategoriesParams) (*GetCategoriesResult, error)
	Update(category *Category) error
	Delete(id int64) error
}
