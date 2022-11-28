package repo

type Like struct {
	ID     int64 `db:"id"`
	PostID int64 `db:"post_id"`
	UserID int64 `db:"user_id"`
	Status bool  `db:"status"`
}

type LikesDislikesCountsResult struct {
	LikesCount    int64
	DislikesCount int64
}

type LikeStorageI interface {
	Create(like *Like) (*Like, error)
	Get(userID, postID int64) (*Like, error)
	GetLikesDislikesCount(postID int64) (*LikesDislikesCountsResult, error)
}
