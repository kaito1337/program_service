package base

import (
	"github.com/google/uuid"
	"time"
)

type BaseViewDTO struct {
	ID        uuid.UUID `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
