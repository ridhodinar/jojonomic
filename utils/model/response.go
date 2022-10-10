package model

type (
	CheckMutasiResponse struct {
		Date         int64   `gorm:"column:created_at" json:"date"`
		Type         string  `gorm:"column:type" json:"type"`
		Gram         float64 `gorm:"column:gold_weight" json:"gram"`
		HargaTopup   float64 `gorm:"column:harga_topup" json:"harga_topup"`
		HargaBuyback float64 `gorm:"column:harga_buyback" json:"harga_buyback"`
		Saldo        float64 `gorm:"column:gold_balance" json:"saldo"`
	}
)
