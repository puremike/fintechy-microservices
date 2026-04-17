package messaging

import (
	"context"
	"fmt"

	"github.com/puremike/fintechy-microservices/shared/contracts"
	"github.com/puremike/fintechy-microservices/shared/utils"
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
}

func NewRabbitMQ(url, exchange, queueName string) (*RabbitMQ, error) {

	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, fmt.Errorf("error connecting to rabbitmq: %v", err)
	}

	channel, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, fmt.Errorf("error creating rabbitmq channel: %v", err)
	}

	rmq := &RabbitMQ{
		Conn:    conn,
		Channel: channel,
	}

	err = rmq.setupExchangeAndQueue(exchange, queueName)

	if err != nil {
		rmq.Close()
		return nil, fmt.Errorf("error setting up exchange and queue: %v", err)
	}

	return rmq, nil
}

func (r *RabbitMQ) Close() {
	if r.Channel != nil {
		r.Channel.Close()
	}

	if r.Conn != nil {
		r.Conn.Close()
	}
}

func (r *RabbitMQ) setupExchangeAndQueue(exchange, queueName string) error {

	err := r.Channel.ExchangeDeclare(
		exchange, // name
		"topic",  // type
		true,     // durability
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)

	if err != nil {
		return fmt.Errorf("error declaring exchange: %v", err)
	}

	err = r.queueBind(queueName, exchange, []string{contracts.EventUserCreated, contracts.EventUserDeleted, contracts.EventUserUpdated})

	if err != nil {
		return fmt.Errorf("error setting up queue and binding: %v", err)
	}

	return nil
}

func (r *RabbitMQ) queueBind(queueName, exchange string, routingKeys []string) error {
	q, err := r.Channel.QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)

	if err != nil {
		return fmt.Errorf("error declaring queue: %v", err)
	}

	for _, key := range routingKeys {
		if err := r.Channel.QueueBind(q.Name, key, exchange, false, nil); err != nil {
			utils.Logger().Errorf("error binding queue %s to exchange %s with routing key %s: %v", q.Name, exchange, key, err)

			return fmt.Errorf("error binding queue %s to exchange %s: %v", q.Name, exchange, err)
		}
	}

	return nil
}

func (r *RabbitMQ) PublishMessage(ctx context.Context, exchange, routingKey string, body []byte) error {
	utils.Logger().Infow("Publishing message to RabbitMQ", "exchange", exchange, "routingKey", routingKey)

	err := r.Channel.PublishWithContext(
		ctx,
		exchange,   // exchange
		routingKey, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType:  "text/plain",
			Body:         body,
			DeliveryMode: amqp.Persistent,
		},
	)

	if err != nil {
		return fmt.Errorf("error publishing message: %v", err)
	}

	return nil
}

type MessageHandler func(context.Context, amqp.Delivery) error

func (r *RabbitMQ) ConsumeMessages(queueName string, handler MessageHandler) error {

	// QoS - Quality of Service: "Don't give me more than 1 unacknowledged message at a time"
	// 1 - Prefetch Count
	// 0 - Prefetch Size (0 means no specific limit)
	// false - Global (false means this QoS setting applies to the current channel only)

	if err := r.Channel.Qos(1, 0, false); err != nil {
		return err
	}

	msgs, err := r.Channel.Consume(
		queueName, // queue
		"",        // consumer
		false,     // auto-ack - we want to manually ack after processing
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)

	if err != nil {
		return err
	}

	ctx := context.Background()

	go func() {
		for msg := range msgs {
			if err := handler(ctx, msg); err != nil {
				utils.Logger().Errorw("Error processing message:", "error", err, "message", string(msg.Body))

				if nackError := msg.Nack(false, false); nackError != nil {
					utils.Logger().Errorw("Error sending Nack for message:", "error", nackError, "message", string(msg.Body))
				}

				continue
			}

			if ackError := msg.Ack(false); ackError != nil {
				utils.Logger().Errorw("Error sending Ack for message:", "error", ackError, "message", string(msg.Body))
			}
		}

	}()

	return nil
}
