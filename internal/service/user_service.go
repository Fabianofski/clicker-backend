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

var ErrInvalidCredentials = errors.New("invalid email or password")
var ErrEmailAlreadyInUse = errors.New("email adress is already in use")
var ErrUsernameAlreadyInUse = errors.New("username is already in use")
var ErrAccountCreationError = errors.New("account couldnt be created")

type UserService struct {
	userRepo  repository.UserRepository
	jwtSecret string
}

func NewUserService(userRepo repository.UserRepository, jwtSecret string) *UserService {
	return &UserService{userRepo, jwtSecret}
}

func (u *UserService) SignUp(ctx context.Context, email string, username string, password string) error {
	user, err := u.userRepo.GetUserWithEmail(ctx, email)
	if err != nil || user != nil {
		if user != nil {
			return ErrEmailAlreadyInUse
		}
		if !errors.Is(repository.ErrUserNotFound, err) {
			return err
		}
	}

	user, err = u.userRepo.GetUserWithUsername(ctx, username)
	if err != nil || user != nil {
		if user != nil {
			return ErrUsernameAlreadyInUse
		}
		if !errors.Is(repository.ErrUserNotFound, err) {
			return err
		}
	}

	pwHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	err = u.userRepo.CreateUser(ctx, email, username, string(pwHash))
	if err != nil {
		return ErrAccountCreationError
	}

	return nil
}

func (u *UserService) Login(ctx context.Context, email string, password string) (string, error) {
	user, err := u.userRepo.GetUserWithEmail(ctx, email)
	if err != nil {
		return "", ErrInvalidCredentials
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PwHash), []byte(password))
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
