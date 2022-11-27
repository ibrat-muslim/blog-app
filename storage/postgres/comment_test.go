package postgres_test

import (
	"testing"

	"github.com/bxcodec/faker/v4"
	"github.com/ibrat-muslim/blog-app/storage/repo"
	"github.com/stretchr/testify/require"
)

func createComment(t *testing.T) *repo.Comment {
	comment, err := strg.Comment().Create(&repo.Comment{
		PostID: 1,
		UserID: 1,
		Description: faker.Sentence(),
	})

	require.NoError(t, err)
	require.NotEmpty(t, comment)

	return comment
}

func deleteComment(id int64, t *testing.T) {
	err := strg.Comment().Delete(id)
	require.NoError(t, err)
}

func TestCreateComment(t *testing.T) {
	cm := createComment(t)
	deleteComment(cm.ID, t)
}

func TestGetComment(t *testing.T) {
	cm := createComment(t)

	comment, err := strg.Comment().Get(cm.ID)
	require.NoError(t, err)
	require.NotEmpty(t, comment)

	deleteComment(comment.ID, t)
}

func TestGetAllComments(t *testing.T) {
	cm := createComment(t)

	comments, err := strg.Comment().GetAll(&repo.GetCommentsParams{
		Limit: 10,
		Page: 1,
	})

	require.NoError(t, err)
	require.GreaterOrEqual(t, len(comments.Comments), 1)
	require.GreaterOrEqual(t, int(comments.Count), 1)

	deleteComment(cm.ID, t)
}

func TestUpdateComment(t *testing.T) {
	cm := createComment(t)

	cm.PostID = 2
	cm.UserID = 2
	cm.Description = faker.Sentence()

	err := strg.Comment().Update(cm)
	require.NoError(t, err)

	deleteComment(cm.ID, t)
}

func TestDeleteComment(t *testing.T) {
	cm := createComment(t)
	deleteComment(cm.ID, t)
}