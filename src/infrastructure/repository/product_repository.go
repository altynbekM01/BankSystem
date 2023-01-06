package repository

import (
	model "GolangwithFrame/src/domain/model"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//
//const (
//	host     = "localhost"
//	port     = 5432
//	user     = "postgres"
//	password = "admin12345"
//	dbname   = "postgres"
//)

type ProductRepository interface {
	CreateProduct(product model.Product)
	UpdateProduct(product model.Product) error
	DeleteProduct(product model.Product) error
	FindAllProducts() []model.Product
	GetProduct(id int) (model.Product, error)
	//Get(product model.Product) model.Product
}

//func NewProductRepository() ProductRepository {
//
//	db := NewDB()
//	return &Database{
//		Connection: db,
//	}
//
//}

func (db *Database) GetProduct(id int) (model.Product, error) {
	product := model.Product{}
	err := db.Connection.Where("id = ?", id).First(&product).Error
	if err != nil {
		return product, err
	} else {
		return product, nil
	}
	//db.Connection.First()
	//db.Connection.First(&prod, "id=?", product.Id)
	////db.Connection.Where("name = ?", "jinzhu").First(&product)
	//return db.Connection.First(&prod, "id=?", product.Id), nil
}

func (db *Database) CreateProduct(product model.Product) {
	db.Connection.Create(&product)
}
func (db *Database) UpdateProduct(product model.Product) error {
	currentProduct := model.Product{}

	err := db.Connection.Where("id = ?", product.Id).First(&currentProduct).Error
	//fmt.Println(err)
	//db.Connection.Save(&product)
	//db.Connection.Where("id = ?", product.Id).First(&product).Error
	if err != nil {
		return db.Connection.Where("id = ?", product.Id).First(&currentProduct).Error
	}
	db.Connection.Save(&product)
	return nil
	//return nil

}
func (db *Database) DeleteProduct(product model.Product) error {
	err := db.Connection.Where("id = ?", product.Id).First(&product).Error
	db.Connection.Delete(&product)
	return err
}
func (db *Database) FindAllProducts() []model.Product {
	var products []model.Product
	db.Connection.Set("gorm:auto_preload", true).Order("id").Find(&products)
	return products
}

//
//func (db *Database) Get(id uint64) {
//	var product model.Product
//	return
//}
