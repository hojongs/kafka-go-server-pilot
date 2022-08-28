package main

import (
	"context"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

var topic = "my-topic"
var partition = 0
var kafkaUrl = "localhost:9092"

func main() {
	println("hello world")
	produceMessage()
}

func produceMessage() {
	conn, err := kafka.DialLeader(context.Background(), "tcp", kafkaUrl, topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	_, err = conn.WriteMessages(
		kafka.Message{Value: []byte("one!")},
		kafka.Message{Value: []byte("two!")},
		kafka.Message{Value: []byte("three!")},
	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	if err := conn.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
}
