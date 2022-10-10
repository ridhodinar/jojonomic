package model

import "time"

type TblHarga struct {
	ReffID       string    `gorm:"column:reff_id" json:"reff_id"`
	AdminID      string    `gorm:"column:admin_id" json:"admin_id"`
	HargaTopup   float64   `gorm:"column:harga_topup" json:"harga_topup"`
	HargaBuyback float64   `gorm:"column:harga_buyback" json:"harga_buyback"`
	CreatedAt    time.Time `gorm:"column:created_at" json:"created_at"`
}

type TblTransaksi struct {
	ReffID       string  `gorm:"column:reff_id" json:"reff_id"`
	Norek        string  `gorm:"column:norek" json:"norek"`
	Type         string  `gorm:"column:type" json:"type"`
	GoldWeight   float64 `gorm:"column:gold_weight" json:"gold_weight"`
	HargaTopup   float64 `gorm:"column:harga_topup" json:"harga_topup"`
	HargaBuyback float64 `gorm:"column:harga_buyback" json:"harga_buyback"`
	GoldBalance  float64 `gorm:"column:gold_balance" json:"gold_balance"`
	CreatedAt    int64   `gorm:"column:created_at" json:"created_at"`
}

type TblRekening struct {
	ReffID       string    `gorm:"column:reff_id" json:"reff_id"`
	Norek        string    `gorm:"column:norek" json:"norek"`
	CustomerName string    `gorm:"column:customer_name" json:"customer_name"`
	GoldBalance  float64   `gorm:"column:gold_balance" json:"gold_balance"`
	CreatedAt    time.Time `gorm:"column:created_at" json:"created_at"`
}
