package models

type Users interface {
	GetAll() []User
	FindByEmail(email string) (*User, error)
	FindById(int64) (*User, error)
}

type User struct {
	ID        int64        `json:"id,omitempty"`
	Name      string       `json:"name" validate:"required,min=3,max=50"`
	Email     string       `json:"email,omitempty" validate:"required,email"`
	Password  string       `json:"password,omitempty" validate:"required,min=6"`
	Role      string       `json:"role,omitempty"`
	Review    []UserReview `json:"review"`
	CreatedAt string       `json:"created_at,omitempty"`
}

type UserReview struct {
	MovieID        int64  `json:"movie_id"`
	UserID         int64  `json:"user_id"`
	Review         string `json:"review"`
	ReviewDatetime string `json:"review_datetime"`
}
