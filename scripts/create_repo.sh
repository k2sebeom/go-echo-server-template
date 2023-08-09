REPO=$1
REPO_FILE=$(echo "$REPO" | tr '[:upper:]' '[:lower:]')
REPO_NAME=$2

touch repositories/$REPO_FILE.go
echo "package repositories

import (
	\"github.com/k2sebeom/go-echo-server-template/models/schema\"

	\"gorm.io/gorm\"
)

type I${REPO}Repository interface {
	Get${REPO}ById(id uint) (*schema.${REPO}, error)
	Create${REPO}(${REPO_NAME} *schema.${REPO}) (*schema.${REPO}, error)
	Update${REPO}(${REPO_NAME} *schema.${REPO}) (*schema.${REPO}, error)
}

type ${REPO}RepositoryImpl struct {
	db *gorm.DB
}

func New${REPO}RepositoryImpl(db *gorm.DB) *${REPO}RepositoryImpl {
	return &${REPO}RepositoryImpl{
		db: db,
	}
}

func (repo *${REPO}RepositoryImpl) Get${REPO}ById(id uint) (*schema.${REPO}, error) {
	${REPO_NAME} := new(schema.${REPO})
	if err := repo.db.First(${REPO_NAME}).Error; err != nil {
		return nil, err
	}
	return ${REPO_NAME}, nil
}

func (repo *${REPO}RepositoryImpl) Create${REPO}(${REPO_NAME} *schema.${REPO}) (*schema.${REPO}, error) {
    if err := repo.db.Create(${REPO_NAME}).Error; err != nil {
		return nil, err
	}
	return ${REPO_NAME}, nil
}

func (repo *${REPO}RepositoryImpl) Update${REPO}(${REPO_NAME} *schema.${REPO}) (*schema.${REPO}, error) {
    if err := repo.db.Save(${REPO_NAME}).Error; err != nil {
		return nil, err
	}
	return ${REPO_NAME}, nil
}
" > repositories/$REPO_FILE.go