package handler

import "myresto/internals/service"

type UserHandler interface {
}

type UserHandlerImpl struct {
	userService service.UserService
}

// Constructor for user handler
func NewUserHandler(userService service.UserService) UserHandler {
	return& UserHandlerImpl{
		userService: userService,
	}
}

