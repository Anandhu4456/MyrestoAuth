package repository

import (
	"gorm.io/gorm"
)

type UserRepository interface {

}

type UserRepoImpl struct {
	db *gorm.DB
}

// Constructor for user repo
func NewUserRepository(db *gorm.DB) UserRepository{
	return &UserRepoImpl{
		db: db,
	}
}

