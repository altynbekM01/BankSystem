package repository

import "GolangwithFrame/src/domain/model"

type CreditRepository interface {
	CreateCredit(credit model.Credit)
	FindAllCredits() []model.Credit
	GetCredit(id int) (model.Credit, error)
}

func (db *Database) CreateCredit(credit model.Credit) {
	db.Connection.Create(&credit)
}

func (db *Database) FindAllCredits() []model.Credit {
	var credits []model.Credit
	db.Connection.Set("gorm:auto_preload", true).Order("id").Find(&credits)
	return credits
}

func (db *Database) FindCreditsById(userId int) ([]model.Credit, error) {
	var credits []model.Credit
	user := model.User{Login: uint(userId)}
	db.Connection.Set("gorm:auto_preload", true).Where("user_id=?", userId).Order("id").Find(&credits)
	if db.Connection.Where("login = ?", userId).First(&user).Error != nil {
		return credits, db.Connection.Where("login = ?", userId).First(&user).Error
	}
	return credits, nil
}
