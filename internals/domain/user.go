package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID             uuid.UUID
	RestaurantName string
	Username       string
	Email          string
	Password       string
	EmailVerified  bool
	CreatedAt      time.Time
}
