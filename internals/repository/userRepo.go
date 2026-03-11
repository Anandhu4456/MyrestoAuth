package repository

import (
	"context"
	"myresto/internals/domain"
	"myresto/internals/dto"

	"gorm.io/gorm"
)

type UserRepository interface {
	SignUp(ctx context.Context, req *dto.SignupRequest) (*domain.User, error)
}

type UserRepoImpl struct {
	db *gorm.DB
}

// Constructor for user repo
func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepoImpl{
		db: db,
	}
}

func (r *UserRepoImpl) SignUp(ctx context.Context, req *dto.SignupRequest) (*domain.User, error) {
	user := &domain.User{
		RestaurantName: req.RestaurantName,
		Username:       req.Username,
		Email:          req.Email,
	}

	if err := r.db.WithContext(ctx).Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
