package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load(".env.example")
	if err != nil {
		log.Fatal("❌ Error loading .env.example file")
	}
}

func GetKafkaTopic() string {
	return os.Getenv("KAFKA_TOPIC")
}

func GetKafkaBroker() string {
	return os.Getenv("KAFKA_BROKER")
}
