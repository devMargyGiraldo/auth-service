package usecase

import (
	"auth_service/internal/domain"
	"auth_service/internal/interfaces/repository"
	"auth_service/pkg/jwt"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type AuthUseCase struct {
	UserRepository repository.UserRepository
}

func NewAuthUseCase(repo repository.UserRepository) *AuthUseCase {
	return &AuthUseCase{
		UserRepository: repo,
	}
}

func (a *AuthUseCase) Register(user *domain.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return a.UserRepository.Save(user)
}

func (a *AuthUseCase) Login(username, password string) (string, error) {
	user, err := a.UserRepository.FindByUsername(username)
	if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid credentials")
	}

	token, err := jwt.GenerateJWT(user.Username)
	if err != nil {
		return "", err
	}

	return token, nil
}
