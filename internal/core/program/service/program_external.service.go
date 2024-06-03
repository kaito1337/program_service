package service

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"program_service/internal/core/program/entity"
)

type ProgramExternalService struct {
	URL string
}

func NewProgramExternalService(url string) *ProgramExternalService {
	return &ProgramExternalService{
		URL: url,
	}
}

func (s *ProgramExternalService) Get(id uuid.UUID, secret string) (*entity.ProgramViewDTO, error) {
	var program entity.ProgramViewDTO
	client := http.Client{}
	req, err := http.NewRequest("GET", s.URL, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("id", id.String())
	q.Add("secret", secret)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	err = json.NewDecoder(resp.Body).Decode(&program)
	if err != nil {
		return nil, err
	}

	return &program, nil
}
