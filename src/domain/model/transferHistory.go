package model

type TransferHistory struct {
	Sender   int `gorm:"type:bigint" json:"sender"`
	Receiver int `gorm:"type:bigint" json:"receiver"`
	Amount   int `gorm:"type:bigint" json:"amount"`
}
