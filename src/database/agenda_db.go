package database

import (
	"go-gin-ddd/src/domain/agendas/model"
	repo "go-gin-ddd/src/domain/agendas/repository"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
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

	// filter for get all
	filter := bson.D{{}}

	cur, err := db.Collection("agendas").Find(dbCtx, filter)
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
func (*AgendaRepository) GetBy(disponibilidade string) ([]*model.Agenda, error) {
	var agendas []*model.Agenda

	disp, _ := strconv.Atoi(disponibilidade)

	filter := bson.D{
		{Key: "$and",
			Value: bson.A{
				bson.D{{Key: "horarios.inicio", Value: bson.D{{Key: "$lt", Value: disp}}}},
				bson.D{{Key: "horarios.fim", Value: bson.D{{Key: "$gt", Value: disp}}}},
			},
		}}

	cur, err := db.Collection("agendas").Find(dbCtx, filter)
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

// Create implements repository.AgendaRepository
func (repo *AgendaRepository) Create(agenda model.Agenda) (*model.Agenda, error) {
	_, err := db.Collection("agendas").InsertOne(dbCtx, agenda)
	if err != nil {
		return nil, err
	}
	return &agenda, nil
}
