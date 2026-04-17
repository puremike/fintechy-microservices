package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/puremike/fintechy-microservices/services/user-service/internal/domain"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *domain.User) (*domain.User, error)
	// GetUserById(ctx context.Context, id string) (*domain.User, error)
	// GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
	// StoreRefreshToken(ctx context.Context, userID, refreshToken string, expires_at time.Time) error
	// UpdateUser(ctx context.Context, user *domain.User, id string) error
	// ValidateRefreshToken(ctx context.Context, refreshToken string) (string, error)
	// ChangePassword(ctx context.Context, pass, id string) error
	// GetUsers(ctx context.Context) (*[]domain.User, error)
	// DeleteUser(ctx context.Context, id string) error
}

type Storage struct {
	Users UserRepository
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{
		Users: &UserStore{db},
	}
}

var (
	QueryBackgroundTimeout = 5 * time.Second
)
