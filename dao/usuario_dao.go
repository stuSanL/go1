package dao

import (
	"awesomeProject/model"
	"context"
	"log"

	"cloud.google.com/go/firestore"
)

func ReadById(id string, ctx context.Context, cliente *firestore.Client) *model.User {
	dsnap, err := cliente.Collection("users").Doc(id).Get(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	var u model.User
	if err = dsnap.DataTo(&u); err != nil {
		log.Fatalln(err)
	}
	u.Id = id
	return &u
}
