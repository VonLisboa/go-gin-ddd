package model

type Agenda struct {
	ID   string `json:"id,omitempty"`
	Cnpj string `json:"teste,omitempty" binding:"required"`
}

func (agenda *Agenda) ValidateCnpj() error {

	return nil
}
