package models

type Movies interface {
	FindById(int64) (*Movie, error)
	FindAll() ([]Movie, error)
	CreateMovie(*Movie) error
	UpdateMovie(*Movie) error
	DeleteMovie(int64) error
}

type Movie struct {
	ID        int64    `json:"id,omitempty"`
	CreatedAt string   `json:"created_at"`
	Title     string   `json:"title" validate:"required,min=3,max=50"`
	Year      int32    `json:"year,omitempty" validate:"required,min=1900,max=2100"`
	Runtime   int32    `json:"runtime,omitempty" validate:"required,min=50"`
	Review    []Review `json:"reviews,omitempty"`
	Genres    []string `json:"genres,omitempty" validate:"required,min=1,max=3"`
}
