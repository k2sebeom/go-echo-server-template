package integration_test

import (
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/k2sebeom/go-echo-server-template/config"
	"github.com/k2sebeom/go-echo-server-template/server"
	"github.com/k2sebeom/go-echo-server-template/utils"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type
(
	IntegrationSuite struct {
		suite.Suite
		Server *server.Server
		M *migrate.Migrate
		DB *gorm.DB
	}

	IntegrationSettings struct {
		EnvPath string
		MigrationPath string
	}
)

func (suite *IntegrationSuite) Setup(setting IntegrationSettings) {
	config.InitializeConfig(setting.EnvPath)
	s, err := server.NewServer()
	if err != nil {
		log.Fatal(err)
	}
	suite.Server = s
	m, err := utils.NewMigration(
		setting.MigrationPath,
		config.GlobalConfig.DB,
	)
	if err != nil {
		log.Fatal(err)
	}
	suite.M = m
	db, err := utils.NewDBInstance(*config.GlobalConfig)
	if err != nil {
		log.Fatal(err)
	}
	suite.DB = db
}
