package service

import (
	"errors"
	"fmt"
	"strconv"

	"go-gin-ddd/src/domain/agendas/model"
	"go-gin-ddd/src/domain/agendas/repository"
)

type AgendaService interface {
	Get() ([]*model.Agenda, error)
	GetBy(disponibilidade string) ([]*model.Agenda, error)
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

// Get implements AgendaService
func (s *agendaService) Get() ([]*model.Agenda, error) {
	// implement things if necessary
	return s.repository.Get()
}

func (s *agendaService) Create(agenda model.Agenda) (*model.Agenda, error) {
	if err := agenda.ValidateCnpj(); err != nil {
		return nil, err
	}

	if err := agenda.ValidateHorario(); err != nil {
		return nil, err
	}

	return s.repository.Create(agenda)
}

func (s *agendaService) validateDisponibilidade(disponibilidade string) error {
	_, err := strconv.Atoi(disponibilidade)
	return err
}

func (s *agendaService) GetBy(disponibilidade string) ([]*model.Agenda, error) {

	if err := s.validateDisponibilidade(disponibilidade); err != nil {
		return nil, err
	}

	agenda, err := s.repository.GetBy(disponibilidade)

	if err != nil {
		agendaNotFoundErr := fmt.Sprintf("Nothing found for \"disponibilidade\": %s", disponibilidade)
		return nil, errors.New(agendaNotFoundErr)
	}
	return agenda, nil
}
