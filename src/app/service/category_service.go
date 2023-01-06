package service

import (
	"GolangwithFrame/src/domain/model"
)

type CategoryService interface {
	CreateCategory(model.Category) model.Category
	FindAllCategory() []model.Category
	UpdateCategory(category model.Category) error
	DeleteCategory(category model.Category) error
	GetCategory(id int) (model.Category, error)
	FindProductByCategory(category_id int) ([]model.Product, error)
}

//type Service struct {
//	Repository repository.ProductRepository
//}
//func New(repo repository.Repository) Service {
//	return &Service{
//		Repository: repo,
//	}
//}

func (service *Service) CreateCategory(category model.Category) model.Category {
	service.Repository.CreateCategory(category)
	return category
}

func (service *Service) FindAllCategory() []model.Category {
	return service.Repository.FindAllCategory()
}

func (service *Service) UpdateCategory(category model.Category) error {
	return service.Repository.UpdateCategory(category)
}

func (service *Service) DeleteCategory(category model.Category) error {
	return service.Repository.DeleteCategory(category)

}

func (service *Service) GetCategory(id int) (model.Category, error) {
	return service.Repository.GetCategory(id)
}
func (service *Service) FindProductsByCategory(category_id int) ([]model.Product, error) {
	return service.Repository.FindProductsByCategory(category_id)
}
