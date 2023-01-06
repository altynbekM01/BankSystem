package model

type Cart struct {
	ProductID uint    `gorm:"NOT NULL" json:"product_id"`
	Product   Product `gorm:"foreignkey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;references:Id" json:"-"`
	UserLogin uint    `gorm:"NOT NULL" json:"user_login"`
	User      User    `gorm:"foreignkey:UserLogin;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;references:Login" json:"-"`
	Quantity  uint    `gorm:"type:integer;default:0" json:"quantity"`
}
