package service

import "GolangwithFrame/src/domain/model"

type TransactionService interface {
}

//type Service struct {
//	Repository repository.ProductRepository
//}
//func New(repo repository.Repository) Service {
//	return &Service{
//		Repository: repo,
//	}
//}

func (service *Service) Transfer(sender int, receiver int, amount int) int {
	return service.Repository.Transfer(sender, receiver, uint(amount))

}

func (service *Service) SeeHistorySender(login uint) ([]model.TransferHistory, error) {
	return service.Repository.SeeHistorySender(login)

}

func (service *Service) SeeHistoryReceiver(login uint) ([]model.TransferHistory, error) {
	return service.Repository.SeeHistoryReceiver(login)

}
