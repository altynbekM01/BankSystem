package service

import "GolangwithFrame/src/domain/model"

type CreditService interface {
	CreateCredit(model.Credit) model.Credit
	FindAllCredits() []model.Credit
	FindCreditsById(userId uint) ([]model.Credit, error)
}

func (service *Service) CreateCredit(credit model.Credit) model.Credit {
	service.Repository.CreateCredit(credit)
	return credit
}

func (service *Service) FindAllCredits() []model.Credit {
	return service.Repository.FindAllCredits()
}

func (service *Service) FindCreditsById(userId uint) ([]model.Credit, error) {
	return service.Repository.FindCreditsById(int(userId))
}
