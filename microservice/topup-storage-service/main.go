package main

import (
	"context"
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
	reader := utils.GetKafkaReader(utils.Config.Kafka.URL, utils.Config.Kafka.TopicTopup, "")

	defer reader.Close()

	fmt.Println("start consuming ... !!")
	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatalln(err)
		}

		req := model.TblTransaksi{}
		err = json.Unmarshal(m.Value, &req)
		if err != nil {
			fmt.Println("error unmarshal:", err)
			continue
		}

		if err := saveTopup(utils.DB, &req); err != nil {
			fmt.Println("error save topup")
			continue
		}

		fmt.Printf("message at topic:%v partition:%v offset:%v	%s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
	}
}

func saveTopup(db *gorm.DB, req *model.TblTransaksi) error {
	conn := db.Begin()

	if err := conn.Create(req).Error; err != nil {
		conn.Rollback()
		return err
	}

	if err := conn.Model(&model.TblRekening{}).Where("norek = ?", req.Norek).Update("gold_balance", req.GoldBalance).Error; err != nil {
		conn.Rollback()
		return err
	}

	return conn.Commit().Error
}
