package vacancy

import (
	"context"
	"errors"
	"github.com/uptrace/bun"
	"log"
	"os"
	"upwork/internal/entity"
	"upwork/internal/service/vacancy"
)

type Repository struct {
	*bun.DB
}

func RepositoryVacancy(DB *bun.DB) Repository {
	return Repository{DB}
}

func (r Repository) CreateVacancy(ctx context.Context, data vacancy.Create) (entity.Vacancy, error) {
	var createVacancy entity.Vacancy

	exist, _ := r.NewSelect().Model((*entity.Vacancy)(nil)).Where("phone_number = ?", data.PhoneNumber).Exists(ctx)
	if exist {
		return entity.Vacancy{}, errors.New("This Phone number (" + data.PhoneNumber + ") is busy!!!")
	}
	createVacancy.FullName = data.FullName
	createVacancy.PhoneNumber = data.PhoneNumber
	createVacancy.Email = data.Email
	createVacancy.CV = &data.CV

	_, err := r.NewInsert().Model(&createVacancy).Exec(ctx)
	if err != nil {
		return entity.Vacancy{}, err
	}

	return createVacancy, nil
}

func (r Repository) GetAllUsersForHR(ctx context.Context) ([]entity.Vacancy, int, error) {
	var allUsers []entity.Vacancy

	query := r.NewSelect().Model(&allUsers)

	query.Where("rating = ?", 0)

	count, errCount := query.ScanAndCount(ctx)
	if errCount != nil {
		return []entity.Vacancy{}, 0, errCount
	}

	return allUsers, count, nil
}

func (r Repository) GetAllUsersForAdmin(ctx context.Context) ([]entity.Vacancy, int, error) {
	var allUsers []entity.Vacancy
	//
	query := r.NewSelect().Model(&allUsers)

	query.Order("rating DESC")

	count, errCount := query.ScanAndCount(ctx)
	if errCount != nil {
		return []entity.Vacancy{}, 0, errCount
	}

	return allUsers, count, nil
}

func (r Repository) SetRating(ctx context.Context, data vacancy.SetRating, id int) (entity.Vacancy, error) {
	var updateVacancy entity.Vacancy
	errFound := r.NewSelect().Model(&updateVacancy).Where("id=?", id).Scan(ctx)
	if errFound != nil {
		return entity.Vacancy{}, errors.New("User not found!")
	}

	if updateVacancy.Rating > 0 {
		return entity.Vacancy{}, errors.New("This user has already been rated !!!")
	}
	query := r.NewUpdate().Model(&updateVacancy)

	updateVacancy.Rating = data.Rating

	_, err := query.Where("id = ?", id).Exec(ctx)

	if err != nil {
		return entity.Vacancy{}, err
	}
	return updateVacancy, nil
}

func (r Repository) SuccessUser(ctx context.Context, id int) (string, error) {
	var user entity.Vacancy

	errFound := r.NewSelect().Model(&user).Where("id=?", id).Scan(ctx)
	if errFound != nil {
		return "", errors.New("User not found!")
	}

	if user.Success {
		return "", errors.New("This user has already been verified.")
	}

	user.Success = true

	query := r.NewUpdate().Model(&user)

	_, err := query.Where("id = ?", id).Exec(ctx)
	if err != nil {
		return "", err
	}
	return "Invited to an interview. Congratulation!!!", nil
}

func (r Repository) RemoveUser(ctx context.Context, id int) (string, error) {
	var user entity.Vacancy

	errFound := r.NewSelect().Model(&user).Where("id=?", id).Scan(ctx)
	if errFound != nil {
		return "", errors.New("User not found!")
	}

	_, err := r.NewDelete().Model(&user).Where("id = ?", id).Exec(ctx)
	if err != nil {
		return "", err
	}

	e := os.Remove(string(*user.CV))
	if e != nil {
		log.Fatal(e)
	}

	return "User deleted successfully", nil
}
