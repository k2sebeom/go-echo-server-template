package repositories

import (
	"errors"

	"github.com/k2sebeom/go-echo-server-template/models"
	"github.com/k2sebeom/go-echo-server-template/models/schema"

	"gorm.io/gorm"
)

type IAuthRepository interface {
	CreateNewPasswordCred(cred *schema.PasswordCredential) (*schema.PasswordCredential, error)
	GetCredByEmail(email string) (*schema.PasswordCredential, error)
	GetCredByUserId(userId uint) (*schema.PasswordCredential, error)
}

type AuthRepositoryImpl struct {
	db *gorm.DB
}

func NewAuthRepositoryImpl(db *gorm.DB) *AuthRepositoryImpl {
	return &AuthRepositoryImpl{
		db: db,
	}
}

func (repo *AuthRepositoryImpl) CreateNewPasswordCred(cred *schema.PasswordCredential) (*schema.PasswordCredential, error) {
	if err := repo.db.Create(cred).Error; err != nil {
		return nil, err
	}
	return cred, nil
}

func (repo *AuthRepositoryImpl) GetCredByEmail(email string) (*schema.PasswordCredential, error) {
	cred := schema.PasswordCredential{}
	if err := repo.db.Where("email = ?", email).First(&cred).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New(models.ErrUserNotFound)
		}
		return nil, err
	}
	return &cred, nil
}

func (repo *AuthRepositoryImpl) GetCredByUserId(userId uint) (*schema.PasswordCredential, error) {
	cred := schema.PasswordCredential{
		UserId: userId,
	}

	if err := repo.db.First(&cred).Error; err != nil {
		return nil, err
	}
	return &cred, nil
}
