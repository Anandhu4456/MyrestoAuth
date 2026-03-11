package db

import (
	"database/sql"
	"fmt"
	"log"
	"myresto/pkg/cfg"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPsqlDB(cfg *cfg.Config) (*gorm.DB, *sql.DB, error) {

	//dsn := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=%s", cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBPort, cfg.DBName, cfg.DBSSLMode)

	dsn := os.Getenv("DB_URL")
	log.Println("DB_URL:", dsn)

	gdb, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, nil, fmt.Errorf("database connection failed due to : %w", err)
	}

	log.Println("[DB] : connection established")

	sqlDB, err := gdb.DB()
	if err != nil {
		return nil, nil, fmt.Errorf("getting sql db failed : %w", err)
	}

	sqlDB.SetMaxOpenConns(25) // maximum number of open connections
	sqlDB.SetMaxIdleConns(10) // maximum number of connection in the idle connection pool
	sqlDB.SetConnMaxLifetime(time.Hour)

	return gdb, sqlDB, nil
}
