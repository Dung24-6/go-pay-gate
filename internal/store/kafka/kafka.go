package kafka

import (
	"github.com/Dung24-6/go-pay-gate/internal/config"
	"github.com/segmentio/kafka-go"
)

type KafkaClient struct {
	Writer *kafka.Writer
	Reader *kafka.Reader
}

// NewKafkaClient creates new Kafka producer and consumer
func NewKafkaClient(cfg *config.KafkaConfig) (*KafkaClient, error) {
	writer := &kafka.Writer{
		Addr:         kafka.TCP(cfg.Brokers...),
		Topic:        cfg.Topic,
		MaxAttempts:  3,
		BatchTimeout: cfg.WriteTimeout,
	}

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  cfg.Brokers,
		Topic:    cfg.Topic,
		GroupID:  cfg.ConsumerGroup,
		MaxBytes: cfg.MaxMessageBytes,
	})

	return &KafkaClient{
		Writer: writer,
		Reader: reader,
	}, nil
}
