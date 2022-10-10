package utils

import (
	"context"
	"log"
	"os"
	"strings"

	"github.com/segmentio/kafka-go"
)

func GetKafkaWriter(kafkaURL, topic string) *kafka.Writer {
	l := log.New(os.Stdout, "kafka writer: ", 0)
	return &kafka.Writer{
		Addr:                   kafka.TCP(kafkaURL),
		Topic:                  topic,
		Balancer:               &kafka.LeastBytes{},
		Logger:                 l,
		AllowAutoTopicCreation: true,
	}
}

func GetKafkaReader(kafkaURL, topic, groupID string) *kafka.Reader {
	brokers := strings.Split(kafkaURL, ",")
	l := log.New(os.Stdout, "kafka reader: ", 0)
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brokers,
		GroupID:  groupID,
		Topic:    topic,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
		Logger:   l,
	})
}

func ConnectToKafka(kafkaURL, topic string) *kafka.Conn {
	conn, err := kafka.DialLeader(context.Background(), "tcp", kafkaURL, topic, 0)
	if err != nil {
		log.Fatal("cannot connect to kafka")
	}

	return conn
}
