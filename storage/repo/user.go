package repo

import "time"

type User struct {
	ID              int64     `db:"id"`
	FirstName       string    `db:"first_name"`
	LastName        string    `db:"last_name"`
	PhoneNumber     *string   `db:"phone_number"`
	Email           string    `db:"email"`
	Gender          *string   `db:"gender"`
	Password        string    `db:"password"`
	Username        string    `db:"username"`
	ProfileImageUrl *string   `db:"profile_image_url"`
	Type            string    `db:"type"`
	CreatedAt       time.Time `db:"created_at"`
}

type UserStorageI interface {
	Create(user *User) (*User, error)
	Get(id int64) (*User, error)
}