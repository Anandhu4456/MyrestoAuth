package service

import "myresto/internals/repository"

type UserService interface {
}

type UserServiceImpl struct {
	userRepo repository.UserRepository
}

// Constructor for user service
func NewUserService(userRepo repository.UserRepository) UserService {
	return &UserServiceImpl{
		userRepo: userRepo,
	}
}