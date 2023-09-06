package vacancy

import (
	"context"
	"upwork/internal/entity"
)

type Service struct {
	repo Repository
}

func ServiceVacancy(repo Repository) Service {
	return Service{repo: repo}
}

func (s Service) CreateVacancy(ctx context.Context, data Create) (entity.Vacancy, error) {
	return s.repo.CreateVacancy(ctx, data)
}

func (s Service) GetAllUsersForHR(ctx context.Context) ([]entity.Vacancy, int, error) {
	return s.repo.GetAllUsersForHR(ctx)
}

func (s Service) GetAllUsersForAdmin(ctx context.Context) ([]entity.Vacancy, int, error) {
	return s.repo.GetAllUsersForAdmin(ctx)
}

func (s Service) SetRating(ctx context.Context, data SetRating, id int) (entity.Vacancy, error) {
	return s.repo.SetRating(ctx, data, id)
}

func (s Service) SuccessUser(ctx context.Context, id int) (string, error) {
	return s.repo.SuccessUser(ctx, id)
}

func (s Service) RemoveUser(ctx context.Context, id int) (string, error) {
	return s.repo.RemoveUser(ctx, id)
}
