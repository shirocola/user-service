package usecase

import "github.com/shirocola/user-service/internal/domain"

type UserUsecase interface {
    GetUserByID(id string) (*domain.User, error)
    CreateUser(user *domain.User) error
}

type userUsecase struct {
    userRepo domain.UserRepository
}

func NewUserUsecase(repo domain.UserRepository) UserUsecase {
    return &userUsecase{
        userRepo: repo,
    }
}

func (u *userUsecase) GetUserByID(id string) (*domain.User, error) {
    return u.userRepo.GetByID(id)
}

func (u *userUsecase) CreateUser(user *domain.User) error {
    return u.userRepo.Create(user)
}
