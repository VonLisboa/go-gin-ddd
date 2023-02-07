package database

import (
	"go-gin-ddd/src/domain/agendas/model"
	repo "go-gin-ddd/src/domain/agendas/repository"
)

// constructor
func NewAgendaDB() repo.AgendaRepository {
	return &dbRepository{}
}

type dbRepository struct {
}

// GetBy implements repository.AgendaRepository
func (*dbRepository) GetBy(disponibilidade string) (*model.Agenda, error) {
	panic("unimplemented")
}

// Create implements repository.AgendaRepository
func (repo *dbRepository) Create(agenda model.Agenda) (*model.Agenda, error) {
	_, err := Mongo.Collection("agendas").InsertOne(MongoContext, agenda)
	if err != nil {
		return nil, err
	}
	return &agenda, nil
}
