package service

import (
	"web/repository"
	"web/types"
)

type UserService struct {
	userRepository *repository.UserRepository
}

func newUserService(userRepository *repository.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (s *UserService) Create(user *types.User) error {
	return s.userRepository.Create(user)
}

func (s *UserService) Update(name string, updateAge int64) error {
	return s.userRepository.Update(name, updateAge)
}

func (s *UserService) Delete(name string) error {
	return s.userRepository.Delete(name)
}

func (s *UserService) Get() []*types.User {
	return s.userRepository.Get()
}
