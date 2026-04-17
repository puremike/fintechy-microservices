package contracts

import "time"

const (
	// The "Addresses" (Routing Keys)
	EventUserCreated = "user.created"
	EventUserDeleted = "user.deleted"
	EventUserUpdated = "user.updated"
)

// type AmqpMessage struct {
// 	Data []byte `json:"data"`
// }

type UserCreatedEvent struct {
	UserID    string    `json:"user_id"`
	Email     string    `json:"email"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Currency  string    `json:"currency"`
	CreatedAt time.Time `json:"created_at"`

	// CorrelationID: Pass this from the API Gateway to ALL services.
	// It allows you to search logs across 4 services for one single request.
	CorrelationID string `json:"correlation_id"`
}
