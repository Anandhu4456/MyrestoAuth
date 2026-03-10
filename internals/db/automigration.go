package db

import (
	"fmt"
	"myresto/internals/domain"

	"gorm.io/gorm"
)

func AutoMigrateModels(db *gorm.DB) error {
	if err := db.AutoMigrate(
		domain.User{},
		domain.Session{},
	); err != nil {

		return fmt.Errorf("migrations failed due to : %w", err)
	}

	return nil

}
