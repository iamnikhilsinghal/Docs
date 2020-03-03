package models

type User struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Gender   string `json:"gender" validate:"required"`
	Dob      string `json:"dob" validate:"required"`
	City     string `json:"city" validate:"required"`
	Status   string
}
