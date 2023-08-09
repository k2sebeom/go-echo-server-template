package services

import (
	"errors"

	"github.com/k2sebeom/go-echo-server-template/config"
	"github.com/k2sebeom/go-echo-server-template/models"
	"github.com/k2sebeom/go-echo-server-template/models/schema"
	"github.com/k2sebeom/go-echo-server-template/repositories"
	"github.com/k2sebeom/go-echo-server-template/utils"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo repositories.IUserRepository
	authRepo repositories.IAuthRepository
	c *config.Config
}

func NewAuthService(
	userRepo repositories.IUserRepository,
	authRepo repositories.IAuthRepository,
) *AuthService {
	return &AuthService{
		userRepo: userRepo,
		authRepo: authRepo,
	};
}

func (s *AuthService) SignUpNewUser(email string, password string) (*schema.User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &schema.User{
		PasswordCredential: &schema.PasswordCredential{
			Email: email,
			PasswordHash: string(hash),
		},
	}

	user, err = s.userRepo.CreateUser(user)

	if err != nil {
		return nil, err
	}

	return user, nil
}


func (s *AuthService) SignInWithPassword(email string, password string) (*schema.User, error) {
	cred, err := s.authRepo.GetCredByEmail(email)
	if err != nil {
		return nil, err
	}
	if !utils.ValidatePassword(password, cred.PasswordHash) {
		return nil, errors.New(models.ErrPasswordMismatch)
	}
	user, err := s.userRepo.GetUserById(cred.UserId)
	return user, err
}

