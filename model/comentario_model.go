package model

import (
	"context"
	"log"

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
	uId := c.Usuario.ID
	usnap, err := c.Usuario.Get(ctx)
	if err != nil {
		log.Fatalln("Erro ao obter snapshot de Usuário: ", err)
	}
	var uDTO UserDTO
	if err = usnap.DataTo(&uDTO); err != nil {
		log.Fatalln("Erro ao criar DTO de usuário: ", err)
	}
	u, err := uDTO.ToEntity()
	if err != nil {
		log.Fatalln("Erro ao criar usuário: ", err)
	}
	u.Id = uId
	return Comentario{
		ID:         c.ID,
		Comentario: c.Comentario,
		Usuario:    u,
	}, err

}
