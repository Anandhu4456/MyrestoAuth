package db

import (
	"fmt"
	"log"
	"myresto/internals/domain"

	"gorm.io/gorm"
)

func AutoMigrateModels(db *gorm.DB) error {
	if err := db.AutoMigrate(
		domain.User{},
		domain.Session{},
		domain.EmailVerificationToken{},
	); err != nil {

		return fmt.Errorf("migrations failed due to : %w", err)
	}

	log.Println("Automigration successful...")
	return nil

}
