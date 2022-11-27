package postgres_test

import (
	"testing"

	"github.com/ibrat-muslim/blog-app/storage/repo"
	"github.com/stretchr/testify/require"
)

func createLike(t *testing.T) *repo.Like {
	like, err := strg.Like().Create(&repo.Like{
		PostID: 1,
		UserID: 1,
		Status: true,
	})

	require.NoError(t, err)
	require.NotEmpty(t, like)

	return like
}

func deleteLike(id int64, t *testing.T) {
	err := strg.Like().Delete(id)
	require.NoError(t, err)
}

func TestCreateLike(t *testing.T) {
	l := createLike(t)
	deleteLike(l.ID, t)
}

func TestGetLike(t *testing.T) {
	l := createLike(t)

	like, err := strg.Like().Get(l.ID)
	require.NoError(t, err)
	require.NotEmpty(t, like)

	deleteLike(like.ID, t)
}

func TestGetAllLikes(t *testing.T) {
	l := createLike(t)

	likes, err := strg.Like().GetAll(&repo.GetLikesParams{
		Limit: 10,
		Page: 1,
	})

	require.NoError(t, err)
	require.GreaterOrEqual(t, len(likes.Likes), 1)
	require.GreaterOrEqual(t, int(likes.Count), 1)

	deleteLike(l.ID, t)
}

func TestUpdateLike(t *testing.T) {
	l := createLike(t)

	l.PostID = 2
	l.UserID = 2
	l.Status = false

	err := strg.Like().Update(l)
	require.NoError(t, err)

	deleteLike(l.ID, t)
}

func TestDeleteLike(t *testing.T) {
	l := createLike(t)
	deleteLike(l.ID, t)
}