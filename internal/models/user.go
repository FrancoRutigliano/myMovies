package models

type Users interface {
	GetAll() []User
	FindByEmail(email string) (*User, error)
}

type User struct {
	ID        int64       `json:"id,omitempty"`
	Name      string      `json:"name" validate:"required,min=3,max=50"`
	Email     string      `json:"email" validate:"required,email"`
	Password  string      `json:"password" validate:"required,min=6"`
	Movies    []UserMovie `json:"movies"`
	Role      string      `json:"role"`
	CreatedAt string      `json:"created_at"`
}

type UserMovie struct {
	MovieID        int64 `json:"movie_id"`
	MovieTheaterID int64 `json:"movie_theater_id"`
}
