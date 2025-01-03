package kafka_consumer

import (
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaConsumer struct {
	consumer *kafka.Consumer
}

func NewKafkaConsumer(broker, groupID, topic string) (*KafkaConsumer, error) {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": broker,
		"group.id":          groupID,
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		return nil, fmt.Errorf("failed to create consumer: %v", err)
	}

	err = consumer.Subscribe(topic, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to subscribe topic: %v", err)
	}

	return &KafkaConsumer{consumer: consumer}, nil
}

func (k *KafkaConsumer) ConsumerMessages() {
	for {
		msg, err := k.consumer.ReadMessage(-1)
		if err == nil {
			log.Printf("Received message: %s", string(msg.Value))
		} else {
			log.Printf("failed read message: %v", err)
		}

	}
}

func (k *KafkaConsumer) Close() {
	k.consumer.Close()
}
