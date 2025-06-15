package models

import (
    "time"

    "github.com/google/uuid"
)

type Crop struct {
    ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
    Name      string    `gorm:"not null"`
    Category  string    `gorm:"not null"`
    Price     float64   `gorm:"type:numeric;not null"`
    Quantity  int       `gorm:"not null"`
    ImageURL  string
    FarmerID  uuid.UUID `gorm:"type:uuid;not null"`
    CreatedAt time.Time `gorm:"autoCreateTime"`
}
