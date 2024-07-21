package service

import (
	"sync"
	"web/repository"
)

var (
	serviceInit     sync.Once
	serviceInstance *Service
)

type Service struct {
	repository *repository.Repository
	User       *User
}

func NewService(repository *repository.Repository) *Service {
	serviceInit.Do(func() {
		serviceInstance = &Service{
			repository: repository,
			User:       newUserService(repository.User),
		}
	})
	return serviceInstance
}
