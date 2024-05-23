package models

type User struct {
	ID        int64       `json:"id,omitempty"`
	Email     string      `json:"email" validate:"required,email"`
	Password  string      `json:"password" validate:"required"`
	Movies    []UserMovie `json:"movies"`
	Role      string      `json:"role" validate:"required,oneof=admin user employee"`
	CreatedAt string      `json:"created_at"`
}

type UserMovie struct {
	MovieID        int64 `json:"movie_id"`
	MovieTheaterID int64 `json:"movie_theater_id"`
}
