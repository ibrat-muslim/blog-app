package postgres_test

import (
	"testing"

	"github.com/bxcodec/faker/v4"
	"github.com/ibrat-muslim/blog-app/storage/repo"
	"github.com/stretchr/testify/require"
)

func createCategory(t *testing.T) *repo.Category {
	category, err := strg.Category().Create(&repo.Category{
		Title: faker.Sentence(),
	})

	require.NoError(t, err)
	require.NotEmpty(t, category)

	return category
}

// func deleteCategory(id int64, t *testing.T) {
// 	err := strg.Category().Delete(id)
// 	require.NoError(t, err)
// }

func TestCreateCategory(t *testing.T) {
	createCategory(t)
}

func TestGetCategory(t *testing.T) {
	c := createCategory(t)

	category, err := strg.Category().Get(c.ID)
	require.NoError(t, err)
	require.NotEmpty(t, category)
}
