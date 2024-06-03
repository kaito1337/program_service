package entity

import (
	"github.com/google/uuid"
	"program_service/utils/base"
)

type ProgramEntity struct {
	base.BaseEntity
	Name      string `json:"name" db:"name"`
	NameEn    string `json:"name_en" db:"name_en"`
	IsPublic  bool   `json:"is_public" db:"is_public"`
	ProjectID string `json:"project_id" db:"project_id"`
}

type ProgramCreationDTO struct {
	ID        uuid.UUID `json:"id,omitempty" db:"id"`
	Name      string    `json:"name" db:"name"`
	NameEn    string    `json:"name_en" db:"name_en"`
	IsPublic  bool      `json:"is_public" db:"is_public"`
	ProjectID string    `json:"project_id" db:"project_id"`
}

type ProgramUpdateDTO struct {
	Name      string `json:"name" db:"name"`
	NameEn    string `json:"name_en" db:"name_en"`
	IsPublic  bool   `json:"is_public" db:"is_public"`
	ProjectID string `json:"project_id" db:"project_id"`
}

type ProgramViewDTO struct {
	base.BaseViewDTO
	Name      string `json:"name" db:"name"`
	NameEn    string `json:"name_en" db:"name_en"`
	IsPublic  bool   `json:"is_public" db:"is_public"`
	ProjectID string `json:"project_id" db:"project_id"`
}
