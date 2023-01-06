package repository

import (
	model "GolangwithFrame/src/domain/model"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type UserRepository interface {
	CreateUser(user model.User)
	UpdateUser(user model.User) error
	DeleteUser(user model.User) error
	FindAllUsers() []model.User
	GetUser(login string) (model.User, error)
}

func (db *Database) GetUser(login uint) (model.User, error) {
	user := model.User{}
	db.Connection.Where("login = ?", login).First(&user)
	//fmt.Println(user)
	//fmt.Println(err)
	//fmt.Println(err)
	//if err != nil {
	//	fmt.Println(user)
	//	return user, err
	//
	//} else {
	//	//fmt.Println(user)
	//	return user, nil
	//}
	return user, nil
}

func (db *Database) GetUser1(login uint) (model.User, error) {
	user := model.User{}
	db.Connection.Where("login = ?", 1).First(&user)
	//fmt.Println(user)
	//fmt.Println(err)
	//fmt.Println(err)
	//if err != nil {
	//	fmt.Println(user)
	//	return user, err
	//
	//} else {
	//	//fmt.Println(user)
	//	return user, nil
	//}
	return user, nil
}

func (db *Database) CreateUser(user model.User) {
	db.Connection.Create(&user)
}
func (db *Database) UpdateUser(user model.User) error {
	currentUser := model.User{}

	err := db.Connection.Where("id = ?", user.Login).First(&currentUser).Error

	if err != nil {
		return db.Connection.Where("id = ?", user.Login).First(&currentUser).Error
	}
	db.Connection.Save(&user)
	return nil

}
func (db *Database) DeleteUser(user model.User) error {
	err := db.Connection.Where("id = ?", user.Login).First(&user).Error
	db.Connection.Delete(&user)
	return err
}
func (db *Database) FindAllUsers() []model.User {
	var users []model.User
	db.Connection.Set("gorm:auto_preload", true).Order("id").Find(&users)
	return users
}
