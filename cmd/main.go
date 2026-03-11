package main

import (
	"log"
	"myresto/internals/db"
	"myresto/internals/router"
	"myresto/pkg/cfg"
	"myresto/pkg/smtp"
)

func main() {
	conf, err := cfg.LoadConfig()
	if err != nil {
		log.Fatalf("config loading error : %v", err)
	}
	gdb, sqlDB, err := db.NewPsqlDB(conf)
	if err != nil {
		log.Fatalf("psql connection error : %v", err)
	}

	defer sqlDB.Close()

	if err := db.AutoMigrateModels(gdb); err != nil {
		log.Fatalf("migration failed : %v", err)
	}

	smtpConfig := smtp.SMTPConfig{
		Host:        conf.SMTPHost,
		Port:        conf.SMTPPort,
		SenderEmail: conf.SMTPEmail,
		AppPassword: conf.SMTPAppPassword,
		FromName:    conf.EmailFromName,
		BaseURL:     conf.BaseURL,
	}

	engine := router.RouteHandler(gdb, smtpConfig, conf)

	if err := engine.Run(":" + conf.PORT); err != nil {
		log.Fatalf("server failed due to : %v", err)
	}

	log.Println("Myresto Authentication Server running on Port : " + conf.PORT)
}
