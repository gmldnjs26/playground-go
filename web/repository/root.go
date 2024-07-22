package repository

import (
	"sync"
)

type Repository struct {
	UserRepository *UserRepository
}

var (
	repositoryInit     sync.Once
	repositoryInstance *Repository
)

func NewRepository() *Repository {
	repositoryInit.Do(func() {
		repositoryInstance = &Repository{
			UserRepository: newUserRepository(),
		}
	})
	return repositoryInstance
}
