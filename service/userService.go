package service

import (
	"github.com/daftkid/mysportapp/domain/user"
)

type UserService interface {
	GetAllUsers() ([]user.User, error)
}

type DefaultUserService struct {
	repo user.UserRepository
}

func (s DefaultUserService) GetAllUsers() ([]user.User, error) {
	return s.repo.FindAll()
}

func NewDefaultUserService(repository user.UserRepository) DefaultUserService {
	return DefaultUserService{repository}
}
