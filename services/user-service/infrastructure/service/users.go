package service

import (
	"context"
	"fmt"

	"github.com/puremike/fintechy-microservices/services/user-service/infrastructure/repository"
	"github.com/puremike/fintechy-microservices/services/user-service/internal/domain"
	"github.com/puremike/fintechy-microservices/shared/utils"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user *domain.User) (*domain.User, error) {

	ctx, cancel := context.WithTimeout(context.Background(), repository.QueryBackgroundTimeout)
	defer cancel()

	if user.FirstName == "" || user.LastName == "" || user.Email == "" || user.Password == "" || len(user.Password) < 6 || len(user.FirstName) < 30 || len(user.LastName) < 30 || len(user.Email) < 50 {
		return nil, fmt.Errorf("invalid user details")
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		utils.Logger().Errorw("Failed to hash user password:", "error", err)
		return nil, fmt.Errorf("failed to hash user password")
	}

	us := &domain.User{
		FirstName:     user.FirstName,
		LastName:      user.LastName,
		Email:         user.Email,
		Password:      hashedPassword,
		Role:          "user",
		Status:        "active",
		Currency:      "NGN",
		CorrelationID: user.CorrelationID,
	}

	createdUser, err := s.repo.CreateUser(ctx, us)
	if err != nil {
		utils.Logger().Errorw("Failed to create user:", "error", err)
		return nil, fmt.Errorf("failed to create user")
	}

	return createdUser, nil
}
