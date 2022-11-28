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

func TestCreateLike(t *testing.T) {
	createLike(t)
}

func TestGetLike(t *testing.T) {
	l := createLike(t)

	like, err := strg.Like().Get(l.UserID, l.PostID)
	require.NoError(t, err)
	require.NotEmpty(t, like)
}

func TestGetLikesDislikesCount(t *testing.T) {  //TODO
	l := createLike(t)

	result, err := strg.Like().GetLikesDislikesCount(l.PostID)

	require.NoError(t, err)
	require.GreaterOrEqual(t, int(result.LikesCount), 0)
	require.GreaterOrEqual(t, int(result.DislikesCount), 0)
}