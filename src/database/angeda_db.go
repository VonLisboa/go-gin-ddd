package database

import (
	"go-gin-ddd/src/domain/agendas/model"
	repo "go-gin-ddd/src/domain/agendas/repository"
)

// constructor
func NewAgendaDB() repo.AgendaRepository {
	return &AgendaRepository{}
}

type AgendaRepository struct {
}

// GetBy implements repository.AgendaRepository
func (*AgendaRepository) GetBy(disponibilidade string) (*model.Agenda, error) {
	panic("unimplemented")
}

// Create implements repository.AgendaRepository
func (repo *AgendaRepository) Create(agenda model.Agenda) (*model.Agenda, error) {
	_, err := db.Collection("agendas").InsertOne(dbCtx, agenda)
	if err != nil {
		return nil, err
	}
	return &agenda, nil
}
