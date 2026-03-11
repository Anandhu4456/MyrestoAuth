package dto

import (
	"time"

	"github.com/google/uuid"
)

type SignupRequest struct {
	RestaurantName string `json:"restaurant_name" binding:"required,min=2,max=255"`
	Username       string `json:"username"        binding:"required,min=3,max=100"`
	Email          string `json:"email"           binding:"required,email"`
}

type SetPasswordRequest struct {
	Token    string `json:"token"    binding:"required"`
	Password string `json:"password" binding:"required,min=8"`
}

type LoginRequest struct {
	Email    string `json:"email"    binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshTokenResponse struct {
	AccessToken string    `json:"access_token"`
	ExpiresAt   time.Time `json:"expires_at"`
}

type VerifyEmailResponse struct {
	Token          string `json:"token"`
	RestaurantName string `json:"restaurant_name"`
	RedirectURL    string `json:"redirect_url"`
}

type UserResponse struct {
	ID             uuid.UUID `json:"id"`
	RestaurantName string    `json:"restaurant_name"`
	Username       string    `json:"username"`
	Email          string    `json:"email"`
	EmailVerified  bool      `json:"email_verified"`
}
