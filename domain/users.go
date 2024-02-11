package domain

import (
	"context"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int64     `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Status    string    `json:"status"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

func LoadTestUser() *User {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("test"), 8)
	return &User{Password: string(hashedPassword), Username: "test_user", Email: "test@test.com"}
}

// UserUseCase represent the user's usecases
type UserUseCase interface {
	Login(context.Context, *User) error
	Register(context.Context, *User) error
}
