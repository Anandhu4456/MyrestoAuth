package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	RestaurantName string    `gorm:"type:varchar(255);not null" json:"restaurant_name"`
	Username       string    `gorm:"type:varchar(255);not null" json:"username"`
	Email          string    `gorm:"type:varchar(255);uniqueIndex;not null" json:"email"`
	Password       string    `gorm:"type:varchar(255)" json:"-"`
	EmailVerified  bool      `gorm:"default:false" json:"email_verified"`
	CreatedAt      time.Time `json:"created_at"`
}
