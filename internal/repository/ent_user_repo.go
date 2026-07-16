package repository

import (
	"context"
	"f4b1.dev/clicker-backend/ent"
	"f4b1.dev/clicker-backend/internal/models"
)

type EntUserRepository struct {
	db *ent.Client
}

func NewEntUserRepository(db *ent.Client) UserRepository {
	return &EntUserRepository{db: db}
}

func (r *EntUserRepository) GetUserById(ctx context.Context, id string) (*models.User, error) {
	panic("unimplemented")
}
func (r *EntUserRepository) CreateUser(ctx context.Context, id string) (*models.User, error) {
	panic("unimplemented")
}
func (r *EntUserRepository) DeleteUser(ctx context.Context, id string) (*models.User, error) {
	panic("unimplemented")
}
