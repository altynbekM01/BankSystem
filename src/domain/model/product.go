package model

import "time"

type Product struct {
	Id          uint64    `gorm:"primary_key;autoincrement;NOT NULL;" json:"id"`
	Name        string    `gorm:"type:varchar(56);NOT NULL" json:"name"`
	Description string    `gorm:"type:varchar(100)" json:"description"`
	CategoryID  uint      `gorm:"default:NULL" json:"category_id"`
	Category    Category  `gorm:"foreign_key:CategoryID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;references:Id" json:"-"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
