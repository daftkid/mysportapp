package service

import (
	"github.com/daftkid/mysportapp/domain/avatar"
)

type AvatarService interface {
	GetAllAvatars() ([]avatar.Avatar, error)
	GetAvatarById(string) (*avatar.Avatar, error)
}

type DefaultAvatarService struct {
	repo avatar.AvatarRepository
}

func (s DefaultAvatarService) GetAllAvatars() ([]avatar.Avatar, error) {
	return s.repo.FindAll()
}

func (s DefaultAvatarService) GetAvatarById(id string) (*avatar.Avatar, error) {
	return s.repo.GetById(id)
}

func NewDefaultAvatarService(repository avatar.AvatarRepository) DefaultAvatarService {
	return DefaultAvatarService{repository}
}
