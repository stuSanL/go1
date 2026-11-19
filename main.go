package main

import (
	"context"
	"fmt"
	"log"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func main() {

	ctx := context.Background()
	sa := option.WithAuthCredentialsFile("", "C:/Users/2025101202010007/Downloads/teste-d887e-firebase-adminsdk-fbsvc-5fbfff8d2b.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	u := ReadByID("com1", ctx, client)
	fmt.Print(u.ID, " ", u.Comentario, " ", u.Usuario)
	defer client.Close()

}
