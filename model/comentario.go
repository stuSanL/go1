package model

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
)

type Comentario struct {
	ID         string
	Comentario string
	Usuario    string
}

type ComentarioDTO struct {
	ID         string                 `firestore:"id"`
	Comentario string                 `firestore:"comentario"`
	Usuario    *firestore.DocumentRef `firestore:"usuario"`
}

func (c ComentarioDTO) ToEntity(ctx context.Context) (Comentario, error) {
	u, err := c.Usuario.Get(ctx)
	if err != nil {
		log.Println("[Comentario ToEntity] Erro ao obter usuário: ", err)
	}
	id := u.Ref.ID
	return Comentario{
		ID:         c.ID,
		Comentario: c.Comentario,
		Usuario:    id,
	}, nil
}

func (c Comentario) ToMap() (string, map[string]any) {
	id := c.ID
	m := map[string]any{
		"comentario": c.Comentario,
		"usuario":    c.Usuario,
	}
	return id, m
}
