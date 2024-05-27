package models

type UserAuth interface {
	EmailExist(string) error
	CreateUser(*User) error
	FindByEmail(string) (*User, error)
}

type UserRegister struct {
	Name     string `json:"name" validate:"required,min=3,max=50"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type UserLogin struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}
