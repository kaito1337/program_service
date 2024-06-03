package service

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"program_service/internal/core/program/entity"
	"program_service/internal/core/program/repository"
)

type ProgramRepository interface {
	Create(dto entity.ProgramCreationDTO) (*entity.ProgramViewDTO, error)

	Get(id uuid.UUID) (*entity.ProgramViewDTO, error)

	Update(id uuid.UUID, dto entity.ProgramUpdateDTO) (*entity.ProgramViewDTO, error)

	Delete(id uuid.UUID) (bool, error)
}

type ProgramService struct {
	repository ProgramRepository
}

func New(dbClient *sqlx.DB) *ProgramService {
	return &ProgramService{
		repository: repository.New(dbClient),
	}
}

func (s *ProgramService) Create() {

}
