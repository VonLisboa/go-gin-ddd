package service

import (
	"errors"
	"fmt"

	"go-gin-ddd/src/domain/agendas/model"
	"go-gin-ddd/src/domain/agendas/repository"
)

type AgendaService interface {
	GetBy(disponibilidade string) (*model.Agenda, error)
	Create(agenda model.Agenda) (*model.Agenda, error)
}

type agendaService struct {
	repository repository.AgendaRepository
}

// constructor
func NewService(repository repository.AgendaRepository) AgendaService {
	return &agendaService{
		repository: repository,
	}
}

func (s *agendaService) Create(agenda model.Agenda) (*model.Agenda, error) {

	if err := agenda.ValidateCnpj(); err != nil {
		return nil, err
	}

	return s.repository.Create(agenda)
}

func (s *agendaService) GetBy(disponibilidade string) (*model.Agenda, error) {
	if len(disponibilidade) == 0 {
		return nil, errors.New("invalid params")
	}

	agenda, err := s.repository.GetBy(disponibilidade)

	if err != nil {
		agendaNotFoundErr := fmt.Sprintf("Nothing found for \"disponibilidade\": %s", disponibilidade)
		return nil, errors.New(agendaNotFoundErr)
	}
	return agenda, nil
}
