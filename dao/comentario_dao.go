package dao

import (
	"awesomeProject/model"
	"context"
	"log"

	"cloud.google.com/go/firestore"
)

func ComentarioCreate(c model.Comentario, ctx context.Context, client *firestore.Client) {
	id, m := c.ToMap(client)
	_, err := client.Collection("comentário").Doc(id).Set(ctx, m)
	if err != nil {
		log.Fatalln("Erro ao criar Comentário: ", err)
	}
}

func ComentarioReadById(ID string, ctx context.Context, client *firestore.Client) (model.Comentario, error) {
	dsnap, err := client.Collection("comentário").Doc(ID).Get(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	var cDTO model.ComentarioDTO
	if err = dsnap.DataTo(&cDTO); err != nil {
		log.Fatalln(err)
	}
	cDTO.ID = ID
	return cDTO.ToEntity(ctx)
}
