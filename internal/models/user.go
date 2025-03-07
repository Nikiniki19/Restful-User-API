package models

import "gorm.io/gorm"

type Users struct {
	UserName string `json:"user_name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"min=8,max=15"`
	Role     string `json:"role"`
}

type SignupUsers struct {
	gorm.Model
	UserName string `json:"username"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"-"`
	Role     string `gorm:"default:user"`
}

type SignupUsersResp struct {
	ID      int    `json:"id"`
	Message string `json:"message"`
}

type Login struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"min=8,max=15"`
}

type LoginResp struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

type LoginDB struct {
	gorm.Model
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"-"`
	Role     string `json:"role"`
}

type FetchAllUsers struct {
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

type FetchUserByID struct {
	ID string `json:"id"`
}

type UpdateUserByID struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
}
