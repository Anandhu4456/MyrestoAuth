package repository

import (
	"context"
	"myresto/internals/domain"
	"myresto/internals/dto"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
	SignUp(ctx context.Context, req *dto.SignupRequest) (*domain.User, error)
	SaveVerificationToken(ctx context.Context, userID uuid.UUID, token string, expiryTime time.Time) error
	FindVerificationToken(ctx context.Context, token string) (*domain.EmailVerificationToken, error)
	VerifyUserEmail(ctx context.Context, userID uuid.UUID) error
	DeleteVerificationToken(ctx context.Context, token string) error
	UpdatePassword(ctx context.Context, userID uuid.UUID, password string) error
	FindByEmail(ctx context.Context, email string) (*domain.User, error)
	CreateSession(ctx context.Context, userID uuid.UUID, refreshToken string, expiry time.Time) error
	FindSessionByToken(ctx context.Context, token string) (*domain.Session, error)
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
		ID: uuid.New(),
		RestaurantName: req.RestaurantName,
		Username:       req.Username,
		Email:          req.Email,
	}

	if err := r.db.WithContext(ctx).Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepoImpl) SaveVerificationToken(ctx context.Context, userID uuid.UUID, token string, expiryTime time.Time) error {
	vtoken := &domain.EmailVerificationToken{
		ID:        uuid.New(),
		UserID:    userID,
		Token:     token,
		ExpiresAt: expiryTime,
	}
	if err := r.db.WithContext(ctx).Create(vtoken).Error; err != nil {
		return err
	}

	return nil

}

func (r *UserRepoImpl) FindVerificationToken(ctx context.Context, token string) (*domain.EmailVerificationToken, error) {
	var vtoken domain.EmailVerificationToken

	if err := r.db.WithContext(ctx).
		Where("token = ?", token).
		First(&vtoken).Error; err != nil {
		return nil, err
	}

	return &vtoken, nil
}

func (r *UserRepoImpl) VerifyUserEmail(ctx context.Context, userID uuid.UUID) error {
	return r.db.WithContext(ctx).
		Model(&domain.User{}).
		Where("id = ?", userID).
		Update("email_verified", true).Error
}

func (r *UserRepoImpl) DeleteVerificationToken(ctx context.Context, token string) error {
	return r.db.WithContext(ctx).
		Where("token = ?", token).
		Delete(&domain.EmailVerificationToken{}).Error
}

func (r *UserRepoImpl) UpdatePassword(ctx context.Context, userID uuid.UUID, password string) error {
	return r.db.WithContext(ctx).
		Model(&domain.User{}).
		Where("id = ?", userID).
		Update("password", password).Error
}

func (r *UserRepoImpl) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	var user domain.User
	if err := r.db.WithContext(ctx).Where("email= ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepoImpl) CreateSession(ctx context.Context, userID uuid.UUID, refreshToken string, expiry time.Time) error {

	session := &domain.Session{
		ID:           uuid.New(),
		UserID:       userID,
		RefreshToken: refreshToken,
		ExpiresAt:    expiry,
		CreatedAt:    time.Now(),
	}

	if err := r.db.WithContext(ctx).Create(session).Error; err != nil {
		return err
	}

	return nil
}

func (r *UserRepoImpl) FindSessionByToken(ctx context.Context, token string) (*domain.Session, error) {

	var session domain.Session

	if err := r.db.WithContext(ctx).
		Where("refresh_token = ?", token).
		First(&session).Error; err != nil {
		return nil, err
	}

	return &session, nil
}
