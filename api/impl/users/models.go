package userService

import (
	"time"
)

type UserModel struct {
	ID          int        `db:"id" json:"id"`
	Username    string     `db:"username" json:"username" validate:"required,min=3,max=20"`
	Email       string     `db:"email" json:"email" validate:"required,email"`
	Password    string     `db:"password" json:"password" validate:"required,min=6,max=30"`
	FirstName   string     `db:"first_name" json:"firstName" validate:"required,min=2,max=30"`
	LastName    string     `db:"last_name" json:"lastName" validate:"required,min=2,max=30"`
	Address     string     `db:"address" json:"address" validate:"required,min=5,max=100"`
	PhoneNumber string     `db:"phone_number" json:"phoneNumber" validate:"required,min=10,max=15"`
	CreatedAt   *time.Time `db:"created_at" json:"createdAt"`
	LastLogin   *time.Time `db:"last_login" json:"lastLogin"`
	IsActive    bool       `db:"is_active" json:"isActive"`
}

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	Token     string `json:"token"`
	ExpiredIn int64
}
