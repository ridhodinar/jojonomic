package main

import (
	"encoding/json"
	"fmt"
	"jojonomic/utils"
	"jojonomic/utils/model"
	"log"

	"gorm.io/gorm"
)

func main() {
	utils.InitConfig()

	utils.InitializeDatabase()

	subscribeData()

	fmt.Println("Successful running apps")
}

func subscribeData() {
	reader := utils.ConnectToKafka(utils.Config.Kafka.URL, utils.Config.Kafka.TopicInputHarga)

	defer reader.Close()

	fmt.Println("start consuming ... !!")
	for {
		m, err := reader.ReadMessage(10e6)
		if err != nil {
			log.Fatalln(err)
		}

		req := model.TblHarga{}
		err = json.Unmarshal(m.Value, &req)
		if err != nil {
			fmt.Println("error unmarshal:", err)
			continue
		}

		if err := saveHarga(utils.DB, &req); err != nil {
			fmt.Println("error save topup")
			continue
		}

		fmt.Printf("message at topic:%v partition:%v offset:%v	%s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
	}
}

func saveHarga(db *gorm.DB, req *model.TblHarga) error {
	conn := db.Begin()

	if err := conn.Create(req).Error; err != nil {
		conn.Rollback()
		return err
	}

	return conn.Commit().Error
}
