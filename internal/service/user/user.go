package user

import (
	"context"
)

type Service struct {
	repo Repository
}

func ServiceUser(repo Repository) Service {
	return Service{repo: repo}
}

func (s Service) CreateUser(ctx context.Context, data Create) (string, error) {
	return s.repo.CreateUser(ctx, data)
}
