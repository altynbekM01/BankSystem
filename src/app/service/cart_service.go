package service

import "GolangwithFrame/src/domain/model"

type CartService interface {
	CreateCart(cart model.Cart)
	DeleteCategory(login uint) error
	FindAllCarts() []model.Cart
	GetUserCart(login uint) (model.Cart, error)
}

func (service *Service) CreateCart(cart model.Cart) model.Cart {
	service.Repository.CreateCart(cart)
	return cart
}

func (service *Service) FindAllCarts() []model.Cart {
	return service.Repository.FindAllCarts()
}

func (service *Service) DeleteCart(login uint) error {
	return service.Repository.DeleteCartByLogin(login)

}

func (service *Service) GetUserCart(login uint) ([]model.Cart, error) {
	return service.Repository.GetUserCart(login)
}
