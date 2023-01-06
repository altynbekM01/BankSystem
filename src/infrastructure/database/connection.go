package database

//
//import (
//	"ProStoreGolang/src/domain/model"
//	"fmt"
//	"github.com/jinzhu/gorm"
//	"time"
//)
//
//var dbase *gorm.DB
//
//const (
//	host     = "localhost"
//	port     = 5432
//	user     = "postgres"
//	password = "admin12345"
//	dbname   = "postgres"
//)
//
//func Init() *gorm.DB {
//	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
//		"password=%s dbname=%s sslmode=disable",
//		host, port, user, password, dbname)
//	db, err := gorm.Open("postgres", psqlInfo)
//	if err != nil {
//		panic("Can't connect to database")
//	}
//	db.AutoMigrate(&model.Token{}, &model.User{}, &model.Product{}, &model.Category{})
//	return db
//}
//
//func GetDB() *gorm.DB {
//	if dbase == nil {
//		dbase = Init()
//		var sleep = time.Duration(1)
//		for dbase == nil {
//			sleep = sleep * 2
//			fmt.Printf("Database is unavailable wait for %d seconds", sleep)
//			time.Sleep(sleep * time.Second)
//			dbase = Init()
//		}
//
//	}
//	return dbase
//}
