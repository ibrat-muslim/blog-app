package repo

type Like struct {
	ID     int64 `db:"id"`
	PostID int64 `db:"post_id"`
	UserID int64 `db:"user_id"`
	Status bool  `db:"status"`
}

type GetLikesParams struct {
	Limit int32
	Page  int32
}

type GetLikesResult struct {
	Likes []*Like
	Count int32
}

type LikeStorageI interface {
	Create(like *Like) (*Like, error)
	Get(id int64) (*Like, error)
	GetAll(params *GetLikesParams) (*GetLikesResult, error)
	Update(like *Like) error
	Delete(id int64) error
}
