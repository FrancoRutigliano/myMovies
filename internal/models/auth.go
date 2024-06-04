package models

type UserAuth interface {
	EmailExist(string) (*User, bool)
	CreateUser(*User) error
	FindByEmail(string) (*User, error)
	UpdateUserPassword(*User) error
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

type ChangePassword struct {
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required,min=6"`
	NewPassword string `json:"new_password" validate:"required,min=6"`
}
