package vacancy

import (
	"context"
	"upwork/internal/entity"
	"upwork/internal/service/vacancy"
)

type UseCase struct {
	vacancy Vacancy
}

func VacancyUseCase(vacancy Vacancy) UseCase {
	return UseCase{vacancy}
}

func (c UseCase) Create(ctx context.Context, data vacancy.Create) (entity.Vacancy, error) {
	return c.vacancy.CreateVacancy(ctx, data)
}

func (c UseCase) GetAllUsersForHR(ctx context.Context) ([]entity.Vacancy, int, error) {
	return c.vacancy.GetAllUsersForHR(ctx)
}

func (c UseCase) GetAllUsersForAdmin(ctx context.Context) ([]entity.Vacancy, int, error) {
	return c.vacancy.GetAllUsersForAdmin(ctx)
}

func (c UseCase) SetRating(ctx context.Context, data vacancy.SetRating, id int) (entity.Vacancy, error) {
	return c.vacancy.SetRating(ctx, data, id)
}

func (c UseCase) SuccessUser(ctx context.Context, id int) (string, error) {
	return c.vacancy.SuccessUser(ctx, id)
}

func (c UseCase) RemoveUser(ctx context.Context, id int) (string, error) {
	return c.vacancy.RemoveUser(ctx, id)
}
