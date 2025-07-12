package producer

import (
	"schedule_gateway/global"
	"schedule_gateway/pkg/loggers"

	"github.com/streadway/amqp"
	"go.uber.org/zap"
)

type IProducer interface{}

type Producer struct {
	logger  *loggers.LoggerZap
	conn    *amqp.Connection
	channel *amqp.Channel
}

func NewProducer() *Producer {
	logger := global.Logger
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		logger.ErrorString("Failed to connect to RabbitMQ", zap.Error(err))
		return nil
	}

	channel, err := conn.Channel()
	if err != nil {
		logger.ErrorString("Failed to open a channel", zap.Error(err))
	}

	return &Producer{
		logger:  logger,
		conn:    conn,
		channel: channel,
	}
}
