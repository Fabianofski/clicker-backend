package repository

import (
	"context"
	"errors"
	"log/slog"

	"f4b1.dev/clicker-backend/ent"
	"f4b1.dev/clicker-backend/ent/user"
	"github.com/google/uuid"
)

var ErrUserNotFound = errors.New("user not found")

type EntUserRepository struct {
	db *ent.Client
}

func NewEntUserRepository(db *ent.Client) UserRepository {
	return &EntUserRepository{db: db}
}

func (r *EntUserRepository) GetUserById(ctx context.Context, id uuid.UUID) (*ent.User, error) {
	u, err := r.db.User.Get(ctx, id)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, ErrUserNotFound
		}

		slog.Error("query by id failed", "error", err, "id", id)
		return nil, err
	}
	return u, nil
}

func (r *EntUserRepository) GetUserWithEmail(ctx context.Context, email string) (*ent.User, error) {
	u, err := r.db.User.Query().Where(user.Email(email)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, ErrUserNotFound
		}

		slog.Error("query user by email failed", "error", err, "email", email)
		return nil, err
	}
	return u, nil
}

func (r *EntUserRepository) GetUserWithUsername(ctx context.Context, username string) (*ent.User, error) {
	u, err := r.db.User.Query().Where(user.Name(username)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, ErrUserNotFound
		}

		slog.Error("query user by username failed", "error", err, "username", username)
		return nil, err
	}
	return u, nil

}

func (r *EntUserRepository) CreateUser(ctx context.Context, email string, username string, pwHash string) error {
	_, err := r.db.User.Create().
		SetEmail(email).
		SetName(username).
		SetPwHash(pwHash).
		Save(ctx)

	if err != nil {
		slog.Error("creating user failed", "error", err, "username", username)
		return err
	}
	return nil
}

func (r *EntUserRepository) DeleteUser(ctx context.Context, id string) (*ent.User, error) {
	panic("unimplemented")
}
