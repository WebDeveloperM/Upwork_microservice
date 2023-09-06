package user

import (
	"context"
	"upwork/internal/service/user"
)

type User interface {
	CreateUser(ctx context.Context, data user.Create) (string, error)
}
