package user

import (
	"context"
	"server/util"
	"strconv"
	"time"

	// "github.com/golang-jwt/jwt/v4"
)

type service struct {
	Repository
	timeout time.Duration
}

func NewService(repository Repository) Service {
	return &service{
		repository,
		time.Duration(2) * time.Second,
	}
}

func (s *service) CreateUser(c context.Context, req *CreateUserReq) (*CreateUserRes, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	// TODO: HASH PASSWORD
	hashedpassword, err := util.hashPassword(req.password)
	if err != nil {
		return nil, err
	}

	u := &User{
		username: req.username,
		email:    req.email,
		password: hashedpassword,
	}

	r, err := s.Repository.CreateUser(ctx, u)
	if err != nil {
		return nil, err
	}

	res := &CreateUserRes{
		user_id:  strconv.Itoa(int(r.user_id)),
		username: r.username,
		email:    r.email,
	}

	return res, nil
}
