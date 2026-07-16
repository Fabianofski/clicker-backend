package repository

import (
	"context"

	"f4b1.dev/clicker-backend/ent"
	"github.com/google/uuid"
)

type UserRepository interface {
	GetUserById(ctx context.Context, id uuid.UUID) (*ent.User, error)
	GetUserWithEmail(ctx context.Context, email string) (*ent.User, error)
	GetUserWithUsername(ctx context.Context, username string) (*ent.User, error)
	CreateUser(ctx context.Context, email string, username string, pwHash string) error
	DeleteUser(ctx context.Context, id string) (*ent.User, error)
}
