package repository

import (
	"sync"
)

type Repository struct {
	User *UserRepository
}

var (
	repositoryInit     sync.Once
	repositoryInstance *Repository
)

func NewRepository() *Repository {
	repositoryInit.Do(func() {
		repositoryInstance = &Repository{
			User: newUserRepository(),
		}
	})
	return repositoryInstance
}
