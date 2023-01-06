package service

import (
	"GolangwithFrame/src/domain/model"
)

type ProductService interface {
	CreateProduct(model.Product) model.Product
	FindAllProducts() []model.Product
	UpdateProduct(product model.Product) error
	DeleteProduct(product model.Product) error
	GetProduct(id int) (model.Product, error)
}

//type Service struct {
//	Repository repository.ProductRepository
//}
//func New(repo repository.Repository) Service {
//	return &Service{
//		Repository: repo,
//	}
//}

func (service *Service) CreateProduct(product model.Product) model.Product {
	service.Repository.CreateProduct(product)
	return product
}

func (service *Service) FindAllProducts() []model.Product {
	return service.Repository.FindAllProducts()
}

func (service *Service) UpdateProduct(product model.Product) error {
	return service.Repository.UpdateProduct(product)
}

func (service *Service) DeleteProduct(product model.Product) error {
	return service.Repository.DeleteProduct(product)

}

func (service *Service) GetProduct(id int) (model.Product, error) {
	return service.Repository.GetProduct(id)
}
