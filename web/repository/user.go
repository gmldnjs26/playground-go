package repository

import (
	"errors"
	"web/types"
)

type UserRepository struct {
	userMap []*types.User
}

func newUserRepository() *UserRepository {
	return &UserRepository{
		userMap: []*types.User{},
	}
}

func (r *UserRepository) Create(u *types.User) error {
	r.userMap = append(r.userMap, u)
	return nil
}

func (r *UserRepository) Update(name string, updateAge int64) error {
	isExisted := false
	for _, user := range r.userMap {
		if user.Name == name {
			isExisted = true
			user.Age = updateAge
			break
		}
	}
	if !isExisted {
		return errors.New("user not found")
	}
	return nil
}

func (r *UserRepository) Delete(name string) error {
	isExisted := false
	for index, user := range r.userMap {
		if user.Name == name {
			isExisted = true
			r.userMap = append(r.userMap[:index], r.userMap[index+1:]...)
			break
		}
	}
	if !isExisted {
		return errors.New("user not found")
	}
	return nil
}

func (r *UserRepository) Get() []*types.User {
	return r.userMap
}
