

package models

import (
	"time"

	"github.com/google/uuid"
)

type CartItem struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID    uint      `gorm:"not null"` // bigint in Postgres
	CropID    uuid.UUID `gorm:"type:uuid;not null"`
	Quantity  int       `gorm:"not null"`
	CreatedAt time.Time
}
