package router

import (
	"myresto/internals/handler"
	"myresto/internals/repository"
	"myresto/internals/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RouteHandler(db *gorm.DB) *gin.Engine {

	// dependency injection
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	r := gin.Default()

	api := r.Group("/api")

	{
		user := api.Group("/users")

		{
			user.POST("/signup", userHandler.SignUp)
		}
	}

	return r
}
