package repository

import (
	"context"
	"errors"
	"log/slog"

	"f4b1.dev/clicker-backend/ent"
	"f4b1.dev/clicker-backend/ent/user"
)

var ErrUserNotFound = errors.New("user not found")

type EntUserRepository struct {
	db *ent.Client
}

func NewEntUserRepository(db *ent.Client) UserRepository {
	return &EntUserRepository{db: db}
}

func (r *EntUserRepository) GetUserById(ctx context.Context, id string) (*ent.User, error) {
	panic("unimplemented")
}

func (r *EntUserRepository) GetUserWithEmail(ctx context.Context, email string) (*ent.User, error) {
	u, err := r.db.User.Query().Where(user.Email(email)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, ErrUserNotFound
		}

		slog.Error("Query failed", "error", err, "email", email)
		return nil, err
	}
	return u, nil
}

func (r *EntUserRepository) CreateUser(ctx context.Context, id string) (*ent.User, error) {
	panic("unimplemented")
}

func (r *EntUserRepository) DeleteUser(ctx context.Context, id string) (*ent.User, error) {
	panic("unimplemented")
}
