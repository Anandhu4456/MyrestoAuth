package handler

import (
	"myresto/internals/dto"
	"myresto/internals/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	SignUp(ctx *gin.Context)
}

type UserHandlerImpl struct {
	userService service.UserService
}

// Constructor for user handler
func NewUserHandler(userService service.UserService) UserHandler {
	return &UserHandlerImpl{
		userService: userService,
	}
}

func (h *UserHandlerImpl) SignUp(ctx *gin.Context) {

	var req dto.SignupRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid requst body",
		})
		return
	}

	response, err := h.userService.SignUp(ctx.Request.Context(), req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, response)
}
