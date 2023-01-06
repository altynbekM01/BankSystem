package model

type Credit struct {
	Id     uint64 `gorm:"primary_key;auto_increment;NOT NULL" json:"id"`
	UserId uint   `gorm:"type:bigint;NOT NULL" json:"user_id"`
	Amount uint64 `gorm:"NOT NULL" json:"amount"`
}
