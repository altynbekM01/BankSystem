package repository

import (
	"GolangwithFrame/src/domain/model"
	"fmt"

	"github.com/jinzhu/gorm"
)

type Database struct {
	Connection *gorm.DB
}

const (
	// Local DB
	host     = "localhost"
	port     = 5432
	user     = "postgres3"
	password = "postgres3"
	dbname   = "postgres3"
)

func NewRepository() Database {

	db := NewDB()
	return Database{
		Connection: db,
	}

}

func (db *Database) CloseDB() {
	err := db.Connection.Close()
	if err != nil {
		panic("Failed to Connect")
	}

}

func NewDB() *gorm.DB {
	//fmt.Println(os.Getenv("LOCAL_DB_USER"))
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		panic("Can't connect to database")
	}
	db.AutoMigrate(&model.User{}, &model.Product{}, &model.Category{}, &model.Cart{}, &model.TransferHistory{}, &model.Credit{})

	return db
}
