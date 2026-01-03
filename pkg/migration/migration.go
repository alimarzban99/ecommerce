package migration

import (
	"fmt"

	"github.com/alimarzban99/ecommerce/config"
)

func GetDatabaseURL() string {
	cfg := config.Cfg.Database
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Name,
		cfg.SSLMode,
	)
}

func GetMigrationPath() string {
	return "database/migrations"
}
