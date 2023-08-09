package services

import (
	"github.com/k2sebeom/go-echo-server-template/models/schema"
	"github.com/k2sebeom/go-echo-server-template/repositories"
)

type UserService struct {
	userRepo repositories.IUserRepository
}

func NewUserService(
	userRepo repositories.IUserRepository,
) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) GetUser(userId uint) (*schema.User, error) {
	return s.userRepo.GetUserById(userId)
}
