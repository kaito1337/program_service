package repository

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"program_service/internal/core/program/entity"
)

type ProgramRepository struct {
	dbClient *sqlx.DB
}

func New(dbClient *sqlx.DB) *ProgramRepository {
	return &ProgramRepository{
		dbClient: dbClient,
	}
}

func (r *ProgramRepository) Create(dto entity.ProgramCreationDTO) (*entity.ProgramViewDTO, error) {
	return nil, nil
}

func (r *ProgramRepository) Get(id uuid.UUID) (*entity.ProgramViewDTO, error) {
	return nil, nil
}

func (r *ProgramRepository) Update(id uuid.UUID, data entity.ProgramUpdateDTO) (*entity.ProgramViewDTO, error) {
	return nil, nil
}

func (r *ProgramRepository) Delete(id uuid.UUID) (bool, error) {
	return false, nil
}
