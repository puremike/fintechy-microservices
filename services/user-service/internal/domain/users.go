package domain

import "time"

type User struct {
	ID        string `json:"id" db:"id"`
	Email     string `json:"email" db:"email"`
	FirstName string `json:"first_name" db:"first_name"`
	LastName  string `json:"last_name" db:"last_name"`
	Password  string `json:"password,omitempty" db:"password"`
	Role      string `json:"role" db:"role"`
	// Status for account lifecycle (active, suspended, pending)
	Status        string    `json:"status" db:"status"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" db:"updated_at"`
	Currency      string    `json:"currency" db:"currency"`
	CorrelationID string    `json:"correlation_id" db:"correlation_id"`
}
