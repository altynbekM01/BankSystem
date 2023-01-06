package model

type Category struct {
	Id   uint64 `gorm:"primary_key;auto_increment;NOT NULL" json:"id"`
	Name string `gorm:"type:varchar(50);NOT NULL" json:"name"`
}
