package service

import (
	"context"
	"fmt"
	"myresto/internals/dto"
	"myresto/internals/repository"
)

type UserService interface {
	SignUp(ctx context.Context, req dto.SignupRequest) (*dto.UserResponse, error)
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

func (s *UserServiceImpl) SignUp(ctx context.Context, req dto.SignupRequest) (*dto.UserResponse, error) {

	user, err := s.userRepo.SignUp(ctx, &req)
	if err != nil {
		return nil, fmt.Errorf("couldn't signup due to : %w", err)
	}

	// implement email sending

	signUpResp := &dto.UserResponse{
		ID:             user.ID,
		RestaurantName: user.RestaurantName,
		Username:       user.Username,
		Email:          user.Email,
		EmailVerified:  user.EmailVerified,
	}

	return signUpResp, nil

}
