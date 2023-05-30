//for create models

package user

import (
	"context"
)

type User struct {
	UserID   int64  `json:"user_id" db:"user_id"`
	Username string `json:"username" db:"username"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}

type CreateUserReq struct {
	Username string `json:"Username" db:"Username"`
	Email    string `json:"Email" db:"Email"`
	Password string `json:"Password" db:"Password"`
}

type CreateUserRes struct {
	UserID   string `json:"UserID" db:"UserID"`
	Username string `json:"Username" db:"Username"`
	Email    string `json:"Email" db:"Email"`
}
type LoginUserReq struct {
	Email    string `json:"Email"`
	Password string `json:"Password"`
}

type LoginUserRes struct {
	accessToken string
	UserID          string `json:"UserID"`
	Username    string `json:"Username"`
}

type Repository interface {
	CreateUser(ctx context.Context, user *User) (*User, error)
}

