package models

type User struct {
	ID        int64       `json:"id"`
	Email     string      `json:"email"`
	Password  string      `json:"password"`
	Movies    []UserMovie `json:"movies"`
	Role      string      `json:"role"`
	CreatedAt string      `json:"created_at"`
}

type UserMovie struct {
	MovieID        int64 `json:"movie_id"`
	MovieTheaterID int64 `json:"movie_theater_id"`
}
