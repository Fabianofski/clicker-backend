package repository

import (
	"context"

	"f4b1.dev/clicker-backend/internal/models"
)

type UserRepository interface {
	GetUserById(ctx context.Context, id string) (*models.User, error)
	CreateUser(ctx context.Context, id string) (*models.User, error)
	DeleteUser(ctx context.Context, id string) (*models.User, error)
}
