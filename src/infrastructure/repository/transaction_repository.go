package repository

import "GolangwithFrame/src/domain/model"

import (
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

type Transaction interface {
	Transfer(product model.Product) error
}

func (db *Database) Transfer(sender int, receiver int, amount uint) int {
	user_sender := model.User{}
	user_receiver := model.User{}

	//db.Model(&User{}).Where("active = ?", true).Update("name", "hello")
	//
	//db.Connection.Model(user)

	db.Connection.Where("login = ?", sender).First(&user_sender)
	history := model.TransferHistory{Sender: int(user_sender.Card), Receiver: receiver, Amount: int(amount)}

	err1 := db.Connection.Where("card = ?", receiver).First(&user_receiver).Error
	if user_sender.Balance < amount {
		return 3
	} else if err1 != nil {
		total := amount + 1000
		db.Connection.Model(&user_sender).Update("balance", user_sender.Balance-total)
		db.Connection.Create(&history)
		return 2
	} else {
		db.Connection.Model(&user_sender).Update("balance", user_sender.Balance-amount)
		db.Connection.Model(&user_receiver).Update("balance", user_receiver.Balance+amount)
		db.Connection.Create(&history)
		return 1
	}
}

func (db *Database) SeeHistorySender(login uint) ([]model.TransferHistory, error) {
	var transactions []model.TransferHistory
	user_sender := model.User{}
	db.Connection.Where("login = ?", login).First(&user_sender)
	sender_card := user_sender.Card

	db.Connection.Set("gorm:auto_preload", true).Where("sender=?", sender_card).Find(&transactions)
	if db.Connection.Where("sender = ?", sender_card).First(&model.TransferHistory{}).Error != nil {
		return transactions, db.Connection.Where("sender = ?", sender_card).First(&model.TransferHistory{}).Error
	}
	return transactions, nil

}

func (db *Database) SeeHistoryReceiver(login uint) ([]model.TransferHistory, error) {
	var transactions []model.TransferHistory
	user_sender := model.User{}
	db.Connection.Where("login = ?", login).First(&user_sender)
	sender_card := user_sender.Card

	db.Connection.Set("gorm:auto_preload", true).Where("receiver=?", sender_card).Find(&transactions)
	if db.Connection.Where("receiver = ?", sender_card).First(&model.TransferHistory{}).Error != nil {
		return transactions, db.Connection.Where("receiver = ?", sender_card).First(&model.TransferHistory{}).Error
	}
	return transactions, nil

}
