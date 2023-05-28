//for create models

package user

import (
	"context"
)

type User struct {
	user_id  int64  `json:"user_id" db:"user_id"`
	username string `json:"username" db:"username"`
	email    string `json:"email" db:"email"`
	password string `json:"password" db:"password"`
}

type Repository interface {
	CreateUser(ctx context.Context, user *User) (*User, error)
}
