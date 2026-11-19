package main

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
)

func ReadById(id string, ctx context.Context, cliente *firestore.Client) *User {
	dsnap, err := cliente.Collection("users").Doc(id).Get(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	var u User
	if err = dsnap.DataTo(&u); err != nil {
		log.Fatalln(err)
	}
	return &u
}
