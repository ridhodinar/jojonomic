package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var Config *config

type config struct {
	Database struct {
		Host     string
		Username string
		Password string
		Name     string
		Port     string
	}
	Kafka struct {
		URL             string
		TopicTopup      string
		TopicInputHarga string
		TopicBuyback    string
	}
}

func InitConfig() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	Config = &config{
		Database: struct {
			Host     string
			Username string
			Password string
			Name     string
			Port     string
		}{
			Host:     os.Getenv("db_host"),
			Username: os.Getenv("db_username"),
			Password: os.Getenv("db_password"),
			Name:     os.Getenv("db_name"),
			Port:     os.Getenv("db_port"),
		},
		Kafka: struct {
			URL             string
			TopicTopup      string
			TopicInputHarga string
			TopicBuyback    string
		}{
			URL:             os.Getenv("kafka_url"),
			TopicTopup:      "topup",
			TopicInputHarga: "input-harga",
			TopicBuyback:    "buyback",
		},
	}
}
