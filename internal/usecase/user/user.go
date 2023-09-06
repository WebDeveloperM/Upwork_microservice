package user

import (
	"context"
	"upwork/internal/service/user"
)

type UseCase struct {
	user User
}

func UserUseCase(user User) UseCase {
	return UseCase{user}
}

func (s UseCase) Create(ctx context.Context, data user.Create) (string, error) {
	return s.user.CreateUser(ctx, data)
}
