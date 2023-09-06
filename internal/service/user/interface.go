package user

import (
	"context"
)

type Repository interface {
	CreateUser(ctx context.Context, data Create) (string, error)
}
