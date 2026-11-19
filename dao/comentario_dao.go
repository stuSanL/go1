package dao

import (
	"awesomeProject/model"
	"context"
	"log"

	"cloud.google.com/go/firestore"
)

func ReadByID(ID string, ctx context.Context, client *firestore.Client) *model.Comentario {
	dsnap, err := client.Collection("comentário").Doc(ID).Get(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	var c model.Comentario
	if err = dsnap.DataTo(&c); err != nil {
		log.Fatalln(err)
	}
	return &c
}
