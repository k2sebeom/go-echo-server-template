package config

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/joho/godotenv"
)

type GoEnv string

const (
	LocalEnv GoEnv = "local"
	ProdEnv GoEnv = "prod"
	DevEn GoEnv = "dev"
)

type Config struct {
	Env GoEnv

	Port string
	DB DBConfig
	Auth AuthConfig
}

var GlobalConfig *Config = nil;

func NewConfig(envDir string) *Config {
	env := os.Getenv("GOENV")
	if env == "" {
		env = string(LocalEnv)
	}

	envFileName := path.Join(envDir, fmt.Sprintf(".env.%s", env))
	err := godotenv.Load(envFileName)
	if err != nil {
		log.Printf("Error Loading variables from %s", envFileName)
	} else {
		log.Printf("Loaded envs from %s", envFileName)
	}
	return &Config{
		Env: GoEnv(env),
		Port: os.Getenv("PORT"),
		DB: NewDBConfig(),
		Auth: NewAuthConfig(),
	}
}

func InitializeConfig(envDir string) {
	GlobalConfig = NewConfig(envDir)
}
