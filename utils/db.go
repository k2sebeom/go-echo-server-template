package utils

import (
	"fmt"

	"github.com/k2sebeom/go-echo-server-template/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func NewDBInstance(config config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		config.DB.Host,
		config.DB.UserName,
		config.DB.Password,
		config.DB.DatabaseName,
		config.DB.Port,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

    if err != nil {
		return nil, err
    }
	return db, nil
}