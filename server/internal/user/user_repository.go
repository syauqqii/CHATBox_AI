package user

import (
	"context"
	"database/sql"
)

type DBTX interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

type repository struct {
	db DBTX
}

func NewRepository(db DBTX) Repository {
	return &repository{db: db}
}

func (r *repository) CreateUser(ctx context.Context, user *User) (*User, error) {
	query := "INSERT INTO users(username, email, password) VALUES (?, ?, ?)"
	result, err := r.db.ExecContext(ctx, query, user.Username, user.Email, user.Password)

	lastInsertID, err := result.LastInsertId()

	if err != nil {
		return &User{}, err
	}

	user.ID = int(lastInsertID)
	return user, nil
}

func (r *repository) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	u := User{}

	query := "SELECT id, email, username, password FROM users WHERE email = ?"
	err := r.db.QueryRowContext(ctx, query, email).Scan(&u.ID, &u.Email, &u.Username, &u.Password)

	if err != nil {
		return &User{}, nil
	}

	return &u, nil
}