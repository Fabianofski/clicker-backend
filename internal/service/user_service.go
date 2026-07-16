package service

import (
	"context"
	"errors"
	"time"

	"f4b1.dev/clicker-backend/ent"
	"f4b1.dev/clicker-backend/internal/middleware"
	"f4b1.dev/clicker-backend/internal/repository"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var ErrInvalidCredentials = errors.New("Invalid email or password.")

type UserService struct {
	userRepo  repository.UserRepository
	jwtSecret string
}

func NewUserService(userRepo repository.UserRepository) *UserService {
	return &UserService{userRepo, "secret"}
}

func (u *UserService) SignUp(ctx context.Context, email string, username string, password string) (*ent.User, error) {
	// user, err := u.userRepo.CreateUser(ctx, "1")
	panic("unimplemented")
}

func (u *UserService) Login(ctx context.Context, email string, password string) (string, error) {
	user, err := u.userRepo.GetUserWithEmail(ctx, email)
	if err != nil {
		return "", ErrInvalidCredentials
	}

	err = bcrypt.CompareHashAndPassword([]byte(password), []byte(user.PwHash))
	if err != nil {
		return "", ErrInvalidCredentials
	}

	claims := &middleware.CustomClaims{
		UserID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(u.jwtSecret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (u *UserService) DeleteAccount(ctx context.Context, email string, password string) (*ent.User, error) {
	panic("unimplemented")
}
