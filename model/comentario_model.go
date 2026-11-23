package model

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
)

type Comentario struct {
	ID         string
	Comentario string
	Usuario    User
}

type ComentarioDTO struct {
	ID         string                 `firestore:"id"`
	Comentario string                 `firestore:"comentario"`
	Usuario    *firestore.DocumentRef `firestore:"usuario"`
}

func (c ComentarioDTO) ToEntity(ctx context.Context) (Comentario, error) {
	if c.Usuario == nil {
		return Comentario{}, fmt.Errorf("[ComentarioDTO] referência de usuário nula para o comentário %q", c.ID)
	}

	uId := c.Usuario.ID
	usnap, err := c.Usuario.Get(ctx)

	if err != nil {
		return Comentario{}, fmt.Errorf("[ComentarioDTO] erro ao obter snapshot de usuário (%s): %w", uId, err)
	}

	if !usnap.Exists() {
		return Comentario{}, fmt.Errorf("[ComentarioDTO] usuário (%s) não encontrado", uId)
	}

	var uDTO UserDTO
	if err := usnap.DataTo(&uDTO); err != nil {
		return Comentario{}, fmt.Errorf("[ComentarioDTO] erro ao converter snapshot do usuário (%s) para DTO: %w", uId, err)
	}

	u := uDTO.ToEntity()
	u.Id = uId

	return Comentario{
		ID:         c.ID,
		Comentario: c.Comentario,
		Usuario:    u,
	}, nil
}

func (c Comentario) ToMap(client *firestore.Client) (string, map[string]any) {
	id := c.ID
	uref := client.Collection("comentario").Doc(c.Usuario.Id)
	m := map[string]any{
		"comentario": c.Comentario,
		"usuario":    uref,
	}
	return id, m
}
