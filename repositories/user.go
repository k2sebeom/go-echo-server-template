package repositories

import (
	"errors"

	"github.com/k2sebeom/go-echo-server-template/models"
	"github.com/k2sebeom/go-echo-server-template/models/schema"

	"gorm.io/gorm"
)

type IUserRepository interface {
	GetUserById(id uint) (*schema.User, error)

	CreateUser(user *schema.User) (*schema.User, error)
	UpdateUser(user schema.User) (*schema.User, error) 

	// ID PW
	GetUserByEmail(email string) (*schema.User, error)

	// Social
	GetUserBySocialHash(hash string) (*schema.User, error)	
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepoImpl(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		db: db,
	}
}

func (repo *UserRepositoryImpl) GetUserById(id uint) (*schema.User, error) {
	user := schema.User{}
	
	if err := repo.db.Preload("Contact").Preload("Address").Preload("SocialCredential").Preload("PasswordCredential").First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New(models.ErrUserNotFound)
		}
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepositoryImpl) CreateUser(user *schema.User) (*schema.User, error) {
	if err := repo.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *UserRepositoryImpl) UpdateUser(user schema.User) (*schema.User, error) {
	if err := repo.db.Save(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepositoryImpl) GetUserByEmail(email string) (*schema.User, error) {
	user := schema.User{}
	
	if err := repo.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepositoryImpl) GetUserBySocialHash(hash string) (*schema.User, error) {
	cred := schema.SocialCredential{}
	
	if err := repo.db.Where("identifier_hash = ?", hash).First(&cred).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New(models.ErrUserNotFound)
		}
		return nil, err
	}
	user := schema.User{}

	if err := repo.db.Table("users").First(&user, cred.UserId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New(models.ErrUserNotFound)
		}
		return nil, err
	}
	return &user, nil
}
