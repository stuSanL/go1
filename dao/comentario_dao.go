package dao

import (
	"awesomeProject/model"
	"context"
	"log"

	"cloud.google.com/go/firestore"
)

//func Create(c model.Comentario, ctx context.Context, client *firestore.Client) error {
//	client.Collection("").Add(ctx, c)
//}

func ReadByID(ID string, ctx context.Context, client *firestore.Client) (model.Comentario, error) {
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
