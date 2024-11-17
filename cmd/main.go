package main

import (
	"encoding/json"
	"log"
	"notification-service/config"
	"notification-service/consumers/email"
	"notification-service/helpers"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("./../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	consumer, err := kafka.NewConsumer(config.KafkaConfig())
	if err != nil {
		panic(err)
	}

	defer consumer.Close()

	err = consumer.Subscribe("signup_user", nil)
	if err != nil {
		panic(err)
	}

	for {
		message, err := consumer.ReadMessage(-1)
		if err != nil {
			panic(err)
		}

		// Decode JSON message
		var messageBody helpers.MessageBody
		if err := json.Unmarshal(message.Value, &messageBody); err != nil {
			panic(err)
		}

		switch messageBody.MediaType {
		case "email":
			email.SendNotification(&messageBody)

		default:
			log.Printf("Unknown event type: %s", messageBody.MediaType)
		}
	}
}
