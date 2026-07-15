package service

import (
	"context"

	"f4b1.dev/clicker-backend/internal/models"
	"f4b1.dev/clicker-backend/internal/repository"
)

type UserService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *UserService {
	return &UserService{userRepo}
}

func (u *UserService) SignUp(ctx context.Context, email string, username string, password string) (*models.User, error) {
	// user, err := u.userRepo.CreateUser(ctx, "1")
	panic("unimplemented")
}

func (u *UserService) Login(ctx context.Context, email string, password string) (*models.User, error) {
	panic("unimplemented")
}

func (u *UserService) DeleteAccount(ctx context.Context, email string, password string) (*models.User, error) {
	panic("unimplemented")
}
