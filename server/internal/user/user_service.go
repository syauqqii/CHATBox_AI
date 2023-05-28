package user

import (
	"context"
	"strconv"
	"time"
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
		UserID: strconv.Itoa(int(r.UserID)),
		username: r.username,

	}
}
