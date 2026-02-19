package rabbitmq

import (
	"context"
	"encoding/json"

	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/ixmael/99minutos/internal/core/ports"
)

type rabbitmqqueue struct {
	conn   *amqp.Connection
	ch     *amqp.Channel
	logger ports.Logger
}

func NewRabbitMQQueue(rabbitmqURL string, logger ports.Logger) (ports.QueueService, error) {
	conn, err := amqp.Dial(rabbitmqURL)
	if err != nil {
		return nil, err
	}
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	service := rabbitmqqueue{
		conn:   conn,
		ch:     ch,
		logger: logger,
	}

	return &service, nil
}

func (q *rabbitmqqueue) RegisterQueue(queue string) error {
	shipmentQueue, err := q.ch.QueueDeclare(
		queue,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	q.logger.Info("queue registered", "queue_name", shipmentQueue.Name)

	return nil
}

func (q *rabbitmqqueue) Publish(ctx context.Context, queue string, message any) error {
	body, err := json.Marshal(message)
	if err != nil {
		return err
	}

	queueMessage := amqp.Publishing{
		ContentType: "application/json",
		Body:        body,
	}

	err = q.ch.PublishWithContext(ctx,
		"",
		queue,
		false,
		false,
		queueMessage,
	)

	return err
}

func (q *rabbitmqqueue) Consume(ctx context.Context, queue string, handler func([]byte) error) error {
	msgs, err := q.ch.Consume(
		queue,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	go func() {
		for msg := range msgs {
			err := handler(msg.Body)
			if err != nil {
				q.logger.Error("orror on processing message", "error", err)
			}
		}
	}()

	return nil
}

func (q *rabbitmqqueue) Stop() {
	q.ch.Close()
	q.conn.Close()
}
