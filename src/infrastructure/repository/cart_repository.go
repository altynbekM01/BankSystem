package repository

import (
	model "GolangwithFrame/src/domain/model"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type CartRepository interface {
	CreateCart(cart model.Cart)
	DeleteCategory(login uint) error
	FindAllCarts() []model.Cart
	GetUserCart(login uint) (model.Cart, error)
}

func (db *Database) CreateCart(cart model.Cart) {
	db.Connection.Create(&cart)

}

func (db *Database) DeleteCartByLogin(login uint) error {
	err := db.Connection.Where("user_login LIKE ?", login).Delete(&model.Cart{}).Error
	return err
}
func (db *Database) FindAllCarts() []model.Cart {
	var carts []model.Cart
	db.Connection.Set("gorm:auto_preload", true).Order("user_login").Find(&carts)
	return carts
}
func (db *Database) GetUserCart(login uint) ([]model.Cart, error) {
	var carts []model.Cart
	cart := model.Cart{UserLogin: login}
	db.Connection.Set("gorm:auto_preload", true).Where("user_login=?", login).Find(&carts)
	if db.Connection.Where("user_login = ?", login).First(&cart).Error != nil {
		return carts, db.Connection.Where("user_login = ?", login).First(&cart).Error
	}
	return carts, nil

}
