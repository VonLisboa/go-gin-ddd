package model

import (
	"errors"
	"strconv"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Agenda struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Horario  string             `json:"horario" bson:"horario" binding:"required"`
	Empresa  Empresa            `bson:"empresa" json:"empresa"`
	Horarios []Horario          `bson:"horarios,omitempty" json:"horarios,omitempty"`
}

type Empresa struct {
	Cnpj string `json:"cnpj" bson:"cnpj" binding:"required"`
	Nome string `json:"nome,omitempty" bson:"nome,omitempty"`
}

type Horario struct {
	Inicio     string `json:"inicio" bson:"inicio"`
	Fim        string `json:"fim" bson:"fim"`
	Disponivel bool   `json:"disponivel" bson:"disponivel"`
}

func (agenda *Agenda) ValidateCnpj() error {

	return nil
}

func (agenda *Agenda) ValidateHorario() error {
	numbers := strings.Split(agenda.Horario, ":")
	hour, err := strconv.Atoi(numbers[0])
	if err != nil {
		return err
	}

	minute, err := strconv.Atoi(numbers[1])
	if err != nil {
		return err
	}

	if hour < 0 || hour > 24 || minute < 0 || minute > 59 {
		return errors.New("invalid time")
	}

	return nil
}
