package service

import "GolangwithFrame/src/infrastructure/repository"

type Service struct {
	Repository repository.Database
}

func New(repo repository.Database) *Service {
	return &Service{
		Repository: repo,
	}
}
