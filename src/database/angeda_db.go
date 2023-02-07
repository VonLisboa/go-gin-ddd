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

// Get implements repository.AgendaRepository
func (*AgendaRepository) Get() ([]*model.Agenda, error) {
	var agendas []*model.Agenda
	cur, err := db.Collection("agendas").Find(dbCtx, nil)
	if err != nil {
		return agendas, err
	}

	// Iterate through the cursor and decode items
	for cur.Next(dbCtx) {
		var item model.Agenda
		err := cur.Decode(&item)
		if err != nil {
			return agendas, err
		}

		agendas = append(agendas, &item)
	}

	if err := cur.Err(); err != nil {
		return agendas, err
	}

	// close the cursor
	cur.Close(dbCtx)

	return agendas, nil
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
