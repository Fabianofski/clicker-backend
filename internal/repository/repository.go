package repository

import (
	"context"

	"f4b1.dev/clicker-backend/ent"
)

type UserRepository interface {
	GetUserById(ctx context.Context, id string) (*ent.User, error)
	GetUserWithEmail(ctx context.Context, email string) (*ent.User, error)
	CreateUser(ctx context.Context, id string) (*ent.User, error)
	DeleteUser(ctx context.Context, id string) (*ent.User, error)
}
