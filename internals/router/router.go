package router

import (
	"myresto/internals/handler"
	"myresto/internals/repository"
	"myresto/internals/service"
	"myresto/pkg/cfg"
	"myresto/pkg/smtp"

	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

func RouteHandler(db *gorm.DB, smtpC smtp.SMTPConfig, cfg *cfg.Config) *gin.Engine {

	// dependency injection
	smtpS := smtp.NewSMTPService(smtpC)
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo, smtpS, cfg)
	userHandler := handler.NewUserHandler(userService)

	r := gin.Default()

	api := r.Group("/api")

	{
		user := api.Group("/users")

		{
			user.POST("/signup", userHandler.SignUp)
			user.GET("/verify-email", userHandler.VerifyEmail)
			user.POST("/set-password", userHandler.SetPassword)
			user.POST("/login", userHandler.Login)
			user.POST("/refresh", userHandler.RefreshToken)
		}
	}

	return r
}
