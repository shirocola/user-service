package persistence

import (
    "errors"
    "github.com/shirocola/user-service/internal/domain"
)

type userRepository struct {
    users map[string]*domain.User
}

func NewUserRepository() domain.UserRepository {
    return &userRepository{
        users: make(map[string]*domain.User),
    }
}

func (r *userRepository) GetByID(id string) (*domain.User, error) {
    if user, ok := r.users[id]; ok {
        return user, nil
    }
    return nil, errors.New("user not found")
}

func (r *userRepository) Create(user *domain.User) error {
    if _, exists := r.users[user.ID]; exists {
        return errors.New("user already exists")
    }
    r.users[user.ID] = user
    return nil
}
