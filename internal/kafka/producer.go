package kafka

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaProducer struct {
	producer *kafka.Producer
	topic    string
}

func NewKafkaProducer(broker, topic string) (*KafkaProducer, error) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": broker})
	if err != nil {
		return nil, err
	}

	return &KafkaProducer{producer: p, topic: topic}, nil
}

func (kp *KafkaProducer) ProduceMessage(key, value string) error {
	err := kp.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &kp.topic, Partition: kafka.PartitionAny},
		Key:            []byte(key),
		Value:          []byte(value),
	}, nil)

	if err != nil {
		return err
	}
	return nil
}

func (kp *KafkaProducer) Close() {
	kp.producer.Close()
}
