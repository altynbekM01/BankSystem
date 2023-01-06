package model

type TransferModel struct {
	Receiver int `json:"receiver"`
	Amount   int `json:"amount"`
}
