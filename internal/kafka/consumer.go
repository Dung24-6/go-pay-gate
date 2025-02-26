package kafka

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaConsumer struct {
	consumer *kafka.Consumer
}

func NewKafkaConsumer(broker, groupID, topic string) (*KafkaConsumer, error) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": broker,
		"group.id":          groupID,
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		return nil, err
	}

	err = c.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		c.Close()
		return nil, err
	}

	return &KafkaConsumer{consumer: c}, nil
}

func (kc *KafkaConsumer) ConsumeMessages(handler func(string, string)) error {
	run := true
	for run {
		msg, err := kc.consumer.ReadMessage(-1)
		if err == nil {
			handler(string(msg.Key), string(msg.Value))
		} else {
			// Có thể thêm logic xử lý lỗi ở đây
			fmt.Printf("Consumer error: %v\n", err)
			if err.(kafka.Error).IsFatal() {
				return err
			}
		}
	}
	return nil
}

func (kc *KafkaConsumer) Close() {
	kc.consumer.Close()
}
