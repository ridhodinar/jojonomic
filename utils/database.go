package utils

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func InitializeDatabase() {
	fmt.Println("Start Connecting to DB...")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable ",
		Config.Database.Host,
		Config.Database.Username,
		Config.Database.Password,
		Config.Database.Name,
		Config.Database.Port,
	)

	cfg := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	}

	db, err := gorm.Open(postgres.Open(dsn), cfg)
	if err != nil {
		fmt.Println("Failed to open connection")
		panic(err)
	}

	fmt.Println("Success to connect DB")

	DB = db
}
