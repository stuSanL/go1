package dao

import (
	"awesomeProject/model"
	"context"
	"log"

	"cloud.google.com/go/firestore"
)

func UserCreate(u model.User, ctx context.Context, client *firestore.Client) {
	id, m := u.ToMap()
	_, err := client.Collection("users").Doc(id).Set(ctx, m, firestore.MergeAll)
	if err != nil {
		log.Fatalln("Erro ao criar usuário: ", err)
	}
}

func UserReadById(id string, ctx context.Context, client *firestore.Client) (model.User, error) {
	dsnap, err := client.Collection("users").Doc(id).Get(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	var uDTO model.UserDTO
	if err = dsnap.DataTo(&uDTO); err != nil {
		log.Fatalln(err)
	}
	uDTO.Id = id
	return uDTO.ToEntity(), err
}
