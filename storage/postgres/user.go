package postgres

import (
	"github.com/ibrat-muslim/blog-app/storage/repo"
	"github.com/jmoiron/sqlx"
)

type userRepo struct {
	db *sqlx.DB
}

func NewUser(db *sqlx.DB) repo.UserStorageI {
	return &userRepo{
		db: db,
	}
}

func (ur *userRepo) Create(user *repo.User) (*repo.User, error) {
	query := `
		INSERT INTO users (
			first_name,
			last_name,
			phone_number,
			email,
			gender,
			password,
			username,
			profile_image_url,
			type
		) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id, created_at
	`

	row := ur.db.QueryRow(
		query,
		user.FirstName,
		user.LastName,
		user.PhoneNumber,
		user.Email,
		user.Gender,
		user.Password,
		user.Username,
		user.ProfileImageUrl,
		user.Type,
	)

	err := row.Scan(
		&user.ID,
		&user.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (ur *userRepo) Get(id int64) (*repo.User, error) {
	query := `
		SELECT
			id,
			first_name,
			last_name,
			phone_number,
			email,
			gender,
			password,
			username,
			profile_image_url,
			type,
			created_at
		FROM users
		WHERE id = $1
	`

	var result repo.User

	err := ur.db.Get(&result, query, id)

	if err != nil {
		return nil, err
	}

	return &result, nil
}
