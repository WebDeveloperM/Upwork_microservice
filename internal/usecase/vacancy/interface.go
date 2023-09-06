package vacancy

import (
	"context"
	"upwork/internal/entity"
	"upwork/internal/service/vacancy"
)

type Vacancy interface {
	CreateVacancy(ctx context.Context, data vacancy.Create) (entity.Vacancy, error)
	GetAllUsersForHR(ctx context.Context) ([]entity.Vacancy, int, error)
	GetAllUsersForAdmin(ctx context.Context) ([]entity.Vacancy, int, error)
	SetRating(ctx context.Context, data vacancy.SetRating, id int) (entity.Vacancy, error)
	SuccessUser(ctx context.Context, id int) (string, error)
	RemoveUser(ctx context.Context, id int) (string, error)
}
