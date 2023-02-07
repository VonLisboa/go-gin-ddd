package repository

import "go-gin-ddd/src/domain/agendas/model"

type AgendaRepository interface {
	Get() ([]*model.Agenda, error)
	GetBy(disponibilidade string) (*model.Agenda, error)
	Create(agenda model.Agenda) (*model.Agenda, error)
}
