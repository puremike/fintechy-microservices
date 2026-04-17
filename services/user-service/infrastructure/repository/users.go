package repository

import (
	"context"
	"database/sql"

	"github.com/puremike/fintechy-microservices/services/user-service/internal/domain"
)

type UserStore struct {
	db *sql.DB
}

func (u *UserStore) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {

	ctx, cancel := context.WithTimeout(ctx, QueryBackgroundTimeout)
	defer cancel()

	query := `INSERT INTO users (email, first_name, last_name, password, role, status, currency, correlation_id)
			  VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
			  RETURNING id, created_at, updated_at`

	tx, err := u.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	if err = tx.QueryRowContext(ctx, query, user.Email, user.FirstName, user.LastName, user.Password, user.Role, user.Status, user.Currency, user.CorrelationID).Scan(&user.ID, &user.Email, &user.FirstName, &user.LastName, &user.Password, &user.Role, &user.Status, &user.CreatedAt, &user.UpdatedAt, &user.Currency, &user.CorrelationID); err != nil {
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return user, nil
}
