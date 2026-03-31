package store

import (
	"context"

	models "github.com/puremike/fintechy-microservices/shared/models"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) (*models.User, error)
	GetUserById(ctx context.Context, id string) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	// StoreRefreshToken(ctx context.Context, userID, refreshToken string, expires_at time.Time) error
	// UpdateUser(ctx context.Context, user *models.User, id string) error
	// ValidateRefreshToken(ctx context.Context, refreshToken string) (string, error)
	// ChangePassword(ctx context.Context, pass, id string) error
	// GetUsers(ctx context.Context) (*[]models.User, error)
	// DeleteUser(ctx context.Context, id string) error
}
