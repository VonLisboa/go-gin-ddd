package model

type Agenda struct {
	ID string `json:"id"`
}

func (agenda *Agenda) ValidateCnpj() error {

	return nil
}
