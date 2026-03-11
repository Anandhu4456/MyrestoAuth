package handler

import (
	"myresto/internals/dto"
	"myresto/internals/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	SignUp(ctx *gin.Context)
	VerifyEmail(ctx *gin.Context)
	SetPassword(ctx *gin.Context)
	Login(ctx *gin.Context)
	RefreshToken(ctx *gin.Context)
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

func (h *UserHandlerImpl) VerifyEmail(ctx *gin.Context) {

	token := ctx.Query("token")

	if token == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "token required",
		})
		return
	}

	err := h.userService.VerifyEmail(ctx.Request.Context(), token)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "email verified successfully",
	})
}

func (h *UserHandlerImpl) SetPassword(ctx *gin.Context) {

	var req dto.SetPasswordRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid body"})
		return
	}

	err := h.userService.SetPassword(ctx.Request.Context(), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "password set successfully",
	})
}

func (h *UserHandlerImpl) Login(ctx *gin.Context) {

	var req dto.LoginRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	resp, err := h.userService.Login(ctx.Request.Context(), req)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *UserHandlerImpl) RefreshToken(ctx *gin.Context) {

	var req dto.RefreshTokenRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request",
		})
		return
	}

	resp, err := h.userService.RefreshToken(ctx.Request.Context(), req)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}
