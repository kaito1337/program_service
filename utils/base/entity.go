package base

import (
	"github.com/google/uuid"
	"time"
)

type BaseEntity struct {
	ID        uuid.UUID   `json:"id" db:"id"`
	CreatedAt time.Time   `json:"created_at" db:"created_at"`
	UpdatedAt time.Time   `json:"updated_at" db:"updated_at"`
	DeletedAt interface{} `json:"deleted_at" db:"deleted_at"`
}
