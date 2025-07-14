package producer

import (
	"encoding/json"
	"fmt"
	"kafka-fraud-detections/config"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Transaction struct {
	AccountID string    `json:"account_id"`
	Amount    float64   `json:"amount"`
	Timestamp time.Time `json:"timestamp"`
}

func SendTransaction(tx Transaction) error {
	config.LoadEnv()

	broker := config.GetKafkaBroker()
	topicName := config.GetKafkaTopic()

	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": broker})
	if err != nil {
		return fmt.Errorf("failed to create producer: %w", err)
	}
	defer p.Close()

	data, err := json.Marshal(tx)
	if err != nil {
		return fmt.Errorf("failed to marshal transaction: %w", err)
	}

	err = p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topicName, Partition: kafka.PartitionAny},
		Value:          data,
	}, nil)

	if err != nil {
		return fmt.Errorf("failed to produce message: %w", err)
	}

	p.Flush(1000)
	fmt.Println("✅ Produced transaction:", string(data))
	return nil
}
