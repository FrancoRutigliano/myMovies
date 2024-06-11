package models

type Review struct {
	MovieID        int64  `json:"movie_id"`
	UserID         int64  `json:"user_id"`
	Message        string `json:"message"`
	ReviewDatetime string `json:"review_datetime"`
}
