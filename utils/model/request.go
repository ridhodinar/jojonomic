package model

type TopupRequest struct {
	Gram  float64 `json:"gram"`
	Harga float64 `json:"harga"`
	Norek string  `json:"norek"`
}

type InputHargaRequest struct {
	AdminID      string  `json:"admin_id"`
	HargaTopup   float64 `json:"harga_topup"`
	HargaBuyback float64 `json:"harga_buyback"`
}

type CheckSaldoRequest struct {
	Norek string `json:"norek"`
}

type CheckMutasiRequest struct {
	Norek     string `json:"norek"`
	StartDate int64  `json:"start_date"`
	EndDate   int64  `json:"end_date"`
}
