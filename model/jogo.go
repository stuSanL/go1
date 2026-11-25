package model

import (
	"time"
)

type Jogo struct {
	Id         string
	Nome       string
	Lancamento time.Time
}

type JogoDTO struct {
	Id         string    `firestore:"id"`
	Nome       string    `firestore:"nome"`
	Lancamento time.Time `firestore:"lancamento"`
}

func (j JogoDTO) ToEntity() Jogo {
	return Jogo{
		Id:         j.Id,
		Nome:       j.Nome,
		Lancamento: j.Lancamento,
	}
}

func (j Jogo) ToMap() (string, map[string]any) {
	id := j.Id
	mapa := map[string]any{
		"nome":       j.Nome,
		"lancamento": j.Lancamento,
	}
	return id, mapa
}
