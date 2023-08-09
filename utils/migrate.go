package utils

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/k2sebeom/go-echo-server-template/config"
)


func NewMigration(filePath string, c config.DBConfig) (*migrate.Migrate, error) {
	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		"care25test",
		"testpw",
		"localhost",
		5432,
		"testdb",
	)
	fileUrl := fmt.Sprintf("file://%s", filePath)
	return migrate.New(fileUrl, dbUrl)
}