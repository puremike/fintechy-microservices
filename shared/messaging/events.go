package messaging

const (
	UserExchange = "user.events"

	// The "Mailbox" for the Transaction Service
	// It will receive ALL user-related info it needs
	TransactionUserQueue = "transaction.service.user.sync"
)
