package config

import (
	"os"
)


type DBConfig struct {
	UserName string
	Password string
	DatabaseName string
	Host string
	Port string
}

func NewDBConfig() DBConfig {
	return DBConfig{
		UserName: os.Getenv("DB_USER_NAME"),
		Password: os.Getenv("DB_PASSWORD"),
		DatabaseName: os.Getenv("DB_NAME"),
		Host: os.Getenv("DB_HOST"),
		Port: os.Getenv("DB_PORT"),
	}
}