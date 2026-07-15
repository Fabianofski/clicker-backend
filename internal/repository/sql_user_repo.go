package repository

import (
	"context"
	"database/sql"

	"f4b1.dev/clicker-backend/internal/models"
)

type SQLUserRepository struct {
	db *sql.DB
}

func NewSQLUserRepository(db *sql.DB) UserRepository {
	return &SQLUserRepository{db: db}
}

func (r *SQLUserRepository) GetUserById(ctx context.Context, id string) (*models.User, error) {
	panic("unimplemented")
}
func (r *SQLUserRepository) CreateUser(ctx context.Context, id string) (*models.User, error) {
	panic("unimplemented")
}
func (r *SQLUserRepository) DeleteUser(ctx context.Context, id string) (*models.User, error) {
	panic("unimplemented")
}
