package jwt

import (
	"errors"
	"myresto/pkg/cfg"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Claims struct {
	UserID uuid.UUID `json:"user_id"`
	Email  string    `json:"email"`
	Type   string    `json:"type"` // access token or refresh token
	jwt.RegisteredClaims
}

func GenerateAccessToken(c *cfg.Config, userID uuid.UUID, email string) (string, time.Time, error) {

	expiresAt := time.Now().Add(time.Duration(c.JWTAccessExpiryMinute) * time.Minute)

	claims := &Claims{
		UserID: userID,
		Email:  email,
		Type:   "access",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "myrestotoday",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(c.JWTAccessSecret))
	if err != nil {
		return "", time.Time{}, err
	}

	return signedToken, expiresAt, nil
}

func GenerateRefreshToken(c *cfg.Config, userID uuid.UUID, email string) (string, time.Time, error) {

	expiresAt := time.Now().Add(time.Duration(c.JWTRefreshExpiryDays) * 24 * time.Hour)

	claims := &Claims{
		UserID: userID,
		Email:  email,
		Type:   "refresh",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "myrestotoday",
		},
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedRefreshToken, err := refreshToken.SignedString([]byte(c.JWTRefreshSecret))
	if err != nil {
		return "", time.Time{}, err
	}

	return signedRefreshToken, expiresAt, nil
}

// validating access token
func ValidateAccessToken(token string, cfg *cfg.Config) (*Claims, error) {
	return validateToken(token, cfg.JWTAccessSecret, "access")
}

// validating refresh token
func ValidateRefreshToken(token string, cfg *cfg.Config) (*Claims, error) {
	return validateToken(token, cfg.JWTRefreshSecret, "refresh")
}

func validateToken(tokenString, secret, expectedType string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	if claims.Type != expectedType {
		return nil, errors.New("invalid token type")
	}

	return claims, nil
}
