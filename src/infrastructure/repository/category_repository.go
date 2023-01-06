package repository

import (
	model "GolangwithFrame/src/domain/model"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type CategoryRepository interface {
	CreateCategory(category model.Category)
	UpdateCategory(category model.Category) error
	DeleteCategory(category model.Category) error
	FindAllProducts() []model.Category
	GetProduct(id int) (model.Category, error)
}

func (db *Database) GetCategory(id int) (model.Category, error) {
	category := model.Category{}
	err := db.Connection.Where("id = ?", id).First(&category).Error
	if err != nil {
		return category, err
	} else {
		return category, nil
	}

}
func (db *Database) CreateCategory(category model.Category) {
	db.Connection.Create(&category)
}
func (db *Database) UpdateCategory(category model.Category) error {
	currentCategory := model.Category{}
	err := db.Connection.Where("id = ?", category.Id).First(&currentCategory).Error
	if err != nil {
		return db.Connection.Where("id = ?", category.Id).First(&currentCategory).Error
	}
	db.Connection.Save(&category)
	return nil
}
func (db *Database) DeleteCategory(category model.Category) error {
	err := db.Connection.Where("id = ?", category.Id).First(&category).Error
	db.Connection.Delete(&category)
	return err
}
func (db *Database) FindAllCategory() []model.Category {
	var categories []model.Category
	db.Connection.Set("gorm:auto_preload", true).Order("id").Find(&categories)
	return categories
}
func (db *Database) FindProductsByCategory(category_id int) ([]model.Product, error) {
	var products []model.Product
	category := model.Category{Id: uint64(category_id)}
	db.Connection.Set("gorm:auto_preload", true).Where("category_id=?", category_id).Order("id").Find(&products)
	if db.Connection.Where("id = ?", category_id).First(&category).Error != nil {
		return products, db.Connection.Where("id = ?", category_id).First(&category).Error
	}
	return products, nil
}
