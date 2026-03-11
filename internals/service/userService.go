package service

import (
	"context"
	"fmt"
	"log"
	"myresto/internals/dto"
	"myresto/internals/repository"
	"myresto/pkg/cfg"
	"myresto/pkg/jwt"
	"myresto/pkg/smtp"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	SignUp(ctx context.Context, req dto.SignupRequest) (*dto.UserResponse, error)
	VerifyEmail(ctx context.Context, token string) error
	SetPassword(ctx context.Context, req dto.SetPasswordRequest) error
	Login(ctx context.Context, req dto.LoginRequest) (*dto.TokenResponse, error)
	RefreshToken(ctx context.Context, req dto.RefreshTokenRequest) (*dto.RefreshTokenResponse, error)
}

type UserServiceImpl struct {
	userRepo repository.UserRepository
	smtp     *smtp.SMTPService
	cfg      *cfg.Config
}

// Constructor for user service
func NewUserService(userRepo repository.UserRepository, smtp *smtp.SMTPService, cfg *cfg.Config) UserService {
	return &UserServiceImpl{
		userRepo: userRepo,
		smtp:     smtp,
		cfg:      cfg,
	}
}

func (s *UserServiceImpl) SignUp(ctx context.Context, req dto.SignupRequest) (*dto.UserResponse, error) {

	user, err := s.userRepo.SignUp(ctx, &req)
	if err != nil {
		return nil, fmt.Errorf("couldn't signup due to : %w", err)
	}

	// generate verification token
	token := uuid.New().String()

	expiry := time.Now().Add(
		time.Duration(s.cfg.VerificationTokenExpiryHours) * time.Hour,
	)

	// store token in db
	if err := s.userRepo.SaveVerificationToken(ctx, user.ID, token, expiry); err != nil {
		return nil, fmt.Errorf("failed to store verification token : %w", err)
	}

	// send verificaton email
	if err := s.smtp.SendVerificationEmail(
		user.Email,
		user.RestaurantName,
		token,
	); err != nil {
		return nil, fmt.Errorf("failed to send verification email : %w", err)
	}

	log.Println("Verfication Email Send Successfully..")

	signUpResp := &dto.UserResponse{
		ID:             user.ID,
		RestaurantName: user.RestaurantName,
		Username:       user.Username,
		Email:          user.Email,
		EmailVerified:  user.EmailVerified,
	}

	return signUpResp, nil

}

func (s *UserServiceImpl) VerifyEmail(ctx context.Context, token string) error {

	vtoken, err := s.userRepo.FindVerificationToken(ctx, token)
	if err != nil {
		return fmt.Errorf("invalid verification token")
	}

	if time.Now().After(vtoken.ExpiresAt) {
		return fmt.Errorf("verification token expired")
	}

	if err := s.userRepo.VerifyUserEmail(ctx, vtoken.UserID); err != nil {
		return err
	}

	if err := s.userRepo.DeleteVerificationToken(ctx, token); err != nil {
		return err
	}

	return nil
}

func (s *UserServiceImpl) SetPassword(ctx context.Context, req dto.SetPasswordRequest) error {

	vtoken, err := s.userRepo.FindVerificationToken(ctx, req.Token)
	if err != nil {
		return fmt.Errorf("invalid token")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	if err := s.userRepo.UpdatePassword(ctx, vtoken.UserID, string(hash)); err != nil {
		return err
	}

	return nil
}

func (s *UserServiceImpl) Login(ctx context.Context, req dto.LoginRequest) (*dto.TokenResponse, error) {

	user, err := s.userRepo.FindByEmail(ctx, req.Email)
	if err != nil {
		return nil, fmt.Errorf("invalid credentials")
	}

	// check password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, fmt.Errorf("invalid credentials")
	}

	// ensure email verified
	if !user.EmailVerified {
		return nil, fmt.Errorf("email not verified")
	}

	// generate access token
	accessToken, _, err := jwt.GenerateAccessToken(
		s.cfg,
		user.ID,
		user.Email,
	)
	if err != nil {
		return nil, err
	}

	// generate refresh token
	refreshToken, refreshExpiry, err := jwt.GenerateRefreshToken(
		s.cfg,
		user.ID,
		user.Email,
	)
	if err != nil {
		return nil, err
	}

	// store refresh token in session table
	err = s.userRepo.CreateSession(ctx, user.ID, refreshToken, refreshExpiry)
	if err != nil {
		return nil, err
	}

	return &dto.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *UserServiceImpl) RefreshToken(ctx context.Context, req dto.RefreshTokenRequest) (*dto.RefreshTokenResponse, error) {

	// Validate refresh JWT
	claims, err := jwt.ValidateRefreshToken(req.RefreshToken, s.cfg)
	if err != nil {
		return nil, fmt.Errorf("invalid refresh token")
	}

	// Check session in DB
	session, err := s.userRepo.FindSessionByToken(ctx, req.RefreshToken)
	if err != nil {
		return nil, fmt.Errorf("session not found")
	}

	// Check expiry
	if time.Now().After(session.ExpiresAt) {
		return nil, fmt.Errorf("refresh token expired")
	}

	// Generate new access token
	accessToken, expiresAt, err :=
		jwt.GenerateAccessToken(s.cfg, claims.UserID, claims.Email)
	if err != nil {
		return nil, err
	}

	return &dto.RefreshTokenResponse{
		AccessToken: accessToken,
		ExpiresAt:   expiresAt,
	}, nil
}
