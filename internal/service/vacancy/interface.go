package vacancy

import (
	"context"
	"upwork/internal/entity"
)

type Repository interface {
	CreateVacancy(ctx context.Context, data Create) (entity.Vacancy, error)
	GetAllUsersForHR(ctx context.Context) ([]entity.Vacancy, int, error)
	GetAllUsersForAdmin(ctx context.Context) ([]entity.Vacancy, int, error)
	SetRating(ctx context.Context, data SetRating, id int) (entity.Vacancy, error)
	SuccessUser(ctx context.Context, id int) (string, error)
	RemoveUser(ctx context.Context, id int) (string, error)
}
