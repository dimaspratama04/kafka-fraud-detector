package main

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

func fraudDetector(tx Transaction) bool {
	return tx.Amount > 1000000
}

func main() {
	config.LoadEnv()
	broker := config.GetKafkaBroker()
	topicName := config.GetKafkaTopic()

	kafkaConfig := &kafka.ConfigMap{
		"bootstrap.servers": broker,
		"group.id":          "transactions-golang",
		"auto.offset.reset": "earliest",
	}

	c, err := kafka.NewConsumer(kafkaConfig)

	if err != nil {
		fmt.Println("Oops, something went wrong...")
		panic(err)
	}

	err = c.Subscribe(string(topicName), nil)

	if err != nil {
		fmt.Println("Ooops, can`t subscribe topics...")
		panic(err)
	}

	fmt.Println("✅ Fraud detection service started...")

	for {
		msg, err := c.ReadMessage(-1)

		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}

		var tx Transaction
		if err := json.Unmarshal(msg.Value, &tx); err != nil {
			fmt.Print("Invalid JSON: %v\n", err)
			continue
		}

		if fraudDetector(tx) {
			fmt.Printf("🚨 Fraud detected! AccountID: %s | Amount: %.2f | Time: %s\n", tx.AccountID, tx.Amount, tx.Timestamp)
		} else {
			fmt.Printf("✅ Normal transaction from AccountID: %s | Amount: %.2f | Time: %s\n", tx.AccountID, tx.Amount, tx.Timestamp)
		}
	}

	c.Close()

}
