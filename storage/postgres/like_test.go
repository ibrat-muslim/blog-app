package postgres_test

import (
	"testing"

	"github.com/ibrat-muslim/blog-app/storage/repo"
	"github.com/stretchr/testify/require"
)

func createLike(t *testing.T) {
	post := createPost(t)
	user := createUser(t)

	err := strg.Like().CreateOrUpdate(&repo.Like{
		PostID: post.ID,
		UserID: user.ID,
		Status: true,
	})

	require.NoError(t, err)
}

func TestCreateLike(t *testing.T) {
	createLike(t)
}

func TestGetLike(t *testing.T) {
	createLike(t)
	post := createPost(t)
	user := createUser(t)

	like, err := strg.Like().Get(post.ID, user.ID)
	require.NoError(t, err)
	require.NotEmpty(t, like)
}

func TestGetLikesDislikesCount(t *testing.T) {  //TODO
	createLike(t)
	post := createPost(t)

	result, err := strg.Like().GetLikesDislikesCount(post.ID)

	require.NoError(t, err)
	require.GreaterOrEqual(t, int(result.LikesCount), 0)
	require.GreaterOrEqual(t, int(result.DislikesCount), 0)
}