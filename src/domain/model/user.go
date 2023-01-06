package model

type User struct {
	Login    uint   `gorm:"type:bigint;primary_key;unique;NOT NULL" json:"login"`
	Password string `json:"password"`
	Card     uint   `gorm:"type:bigint;" json:"card"`
	Balance  uint   `gorm:"type:bigint" json:"balance"`
	Bonus    uint   `gorm:"type:bigint;default:0" json:"bonus"`
}
