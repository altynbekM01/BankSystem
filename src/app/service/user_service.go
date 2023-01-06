package service

import (
	"GolangwithFrame/src/domain/model"
)

type UserService interface {
	CreateUser(model.User) model.User
	FindAllUsers() []model.User
	UpdateUser(user model.User) error
	DeleteUser(user model.User) error
	GetUser(login string) (model.User, error)
}

func (service *Service) CreateUser(user model.User) model.User {
	service.Repository.CreateUser(user)
	return user
}

func (service *Service) FindAllUsers() []model.User {
	return service.Repository.FindAllUsers()
}

func (service *Service) UpdateUser(user model.User) error {
	return service.Repository.UpdateUser(user)
}

func (service *Service) DeleteUser(user model.User) error {
	return service.Repository.DeleteUser(user)

}

func (service *Service) GetUser(login uint) (model.User, error) {
	return service.Repository.GetUser(login)
}

func (service *Service) GetUser1(login uint) (model.User, error) {
	return service.Repository.GetUser1(login)
}
