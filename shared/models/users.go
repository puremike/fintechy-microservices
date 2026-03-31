package models

import "time"

type User struct {
	ID       string `json:"id" db:"id"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password,omitempty" db:"password"`
	Role     string `json:"role" db:"role"`
	// Status for account lifecycle (active, suspended, pending_verification)
	Status string `json:"status" db:"status"`

	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type UserProfile struct {
	UserID    string `json:"user_id" db:"user_id"`
	FirstName string `json:"first_name" db:"first_name"`
	LastName  string `json:"last_name" db:"last_name"`
	Currency  string `json:"currency" db:"currency"`
}
