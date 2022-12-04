package repo

import "time"

const (
	UserTypeSuperAdmin = "superadmin"
	UserTypeUser       = "user"
)

type User struct {
	ID              int64     `db:"id"`
	FirstName       string    `db:"first_name"`
	LastName        string    `db:"last_name"`
	PhoneNumber     *string   `db:"phone_number"`
	Email           string    `db:"email"`
	Gender          *string   `db:"gender"`
	Password        string    `db:"password"`
	Username        *string   `db:"username"`
	ProfileImageUrl *string   `db:"profile_image_url"`
	Type            string    `db:"type"`
	CreatedAt       time.Time `db:"created_at"`
}

type GetUsersParams struct {
	Page   int32
	Limit  int32
	Search string
}

type GetUsersResult struct {
	Users []*User
	Count int32
}

type UpdatePassword struct {
	UserID   int64
	Password string
}

type UserStorageI interface {
	Create(user *User) (*User, error)
	Get(id int64) (*User, error)
	GetByEmail(email string) (*User, error)
	GetAll(params *GetUsersParams) (*GetUsersResult, error)
	Update(user *User) error
	Delete(id int64) error
	UpdatePassword(req *UpdatePassword) error
}
