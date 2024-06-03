package models

import "time"

type Movie struct {
	ID        int64     `json:"id,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	Title     string    `json:"title" validate:"required,min=3,max=50"`
	Year      int32     `json:"year,omitempty" validate:"required,min=1900,max=2100"`
	Runtime   int32     `json:"runtime,omitempty" validate:"required,min=50"`
	Genres    []string  `json:"genres,omitempty" validate:"required,min=1,max=3"`
}
