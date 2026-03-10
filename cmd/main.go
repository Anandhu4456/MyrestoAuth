package main

import (
	"log"
	"myresto/internals/db"
	"myresto/pkg/cfg"
)

func main() {
	conf, err := cfg.LoadConfig()
	if err != nil {
		log.Fatalf( "config loading error : %v", err)
	}
	gdb, sqlDB, err := db.NewPsqlDB(conf)
	if err != nil {
		log.Fatalf("psql connection error : %v", err)
	}

	defer sqlDB.Close()

	if err := db.AutoMigrateModels(gdb); err != nil {
		log.Fatalf("migration failed : %v", err)
	}
}