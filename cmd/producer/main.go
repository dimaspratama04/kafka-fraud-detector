package main

import (
	"time"

	"kafka-fraud-detections/internal/producer"
)

func main() {
	tx := producer.Transaction{
		AccountID: "acc101",
		Amount:    3000000,
		Timestamp: time.Now(),
	}

	if err := producer.SendTransaction(tx); err != nil {
		panic(err)
	}
}
