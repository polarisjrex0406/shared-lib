package database

import (
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/omimic12/shared-lib/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(dsn string) {
	var err error
	// Open the database connection
	DB, err = gorm.Open(
		postgres.Open(dsn),
		&gorm.Config{
			TranslateError: true,
		},
	)
	if err != nil {
		log.Fatal(err)
	}
}

func ConnectionString() string {
	cfg, err := config.GetConfig()
	if err != nil {
		return ""
	}
	// Construct the connection string
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=%s",
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.DBName,
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.SSLMode,
	)
	return dsn
}
