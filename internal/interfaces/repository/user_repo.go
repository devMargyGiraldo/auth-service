package repository

import (
	"auth_service/internal/domain"
	"errors"
)

type UserRepository interface {
	Save(user *domain.User) error
	FindByUsername(username string) (*domain.User, error)
}

// Mock implementation for now
type InMemoryUserRepo struct {
	users map[string]*domain.User
}

func NewInMemoryUserRepo() *InMemoryUserRepo {
	return &InMemoryUserRepo{
		users: make(map[string]*domain.User),
	}
}

func (repo *InMemoryUserRepo) Save(user *domain.User) error {
	repo.users[user.Username] = user
	return nil
}

func (repo *InMemoryUserRepo) FindByUsername(username string) (*domain.User, error) {
	user, exists := repo.users[username]
	if !exists {
		return nil, errors.New("user not found")
	}
	return user, nil
}
