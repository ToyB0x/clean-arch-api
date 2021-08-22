package utils

import (
	"github.com/golang-migrate/migrate/v4/database/mysql"

	"github.com/toaru/clean-arch-api/config"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/toaru/clean-arch-api/pkg/server/infra/store"
)

func MigrateUp(filePath string) error {
	m, err := getMigrateInstance("file://" + filePath)
	if err != nil {
		return err
	}
	return m.Up()
}

func MigrateDrop(filePath string) error {
	m, err := getMigrateInstance("file://" + filePath)
	if err != nil {
		return err
	}
	return m.Drop()
}

func getMigrateInstance(filePath string) (*migrate.Migrate, error) {
	con := store.NewSqlHandler(config.Configs.APP_ENV).Conn
	driver, _ := mysql.WithInstance(con, &mysql.Config{})
	return migrate.NewWithDatabaseInstance(
		filePath,
		"clean-arch-api",
		driver,
	)
}
