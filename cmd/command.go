package cmd

import (
	"log"

	"github.com/omimic12/shared-lib/config"
	"github.com/omimic12/shared-lib/migrations"
	"gorm.io/gorm"
)

func Commands(db *gorm.DB) bool {
	cfg, err := config.GetConfig()
	if err != nil {
		return false
	}

	if cfg.Command.Migrate {
		if err := migrations.Migrate(db); err != nil {
			log.Fatalf("error migration: %v", err)
		}
		log.Println("migration completed successfully")
	}

	if cfg.Command.Seed {
		if err := migrations.Seeder(db); err != nil {
			log.Fatalf("error migration seeder: %v", err)
		}
		log.Println("seeder completed successfully")
	}

	return true
}
