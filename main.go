package main

import (
	"awesomeProject/dao"
	"context"
	"fmt"
	"log"
	"os"

	firebase "firebase.google.com/go"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

func main() {

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	credentials := os.Getenv("CREDENTIALS_FILE")

	ctx := context.Background()
	sa := option.WithAuthCredentialsFile("", credentials)

	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	u := dao.ReadByID("com1", ctx, client)
	fmt.Print(u.ID, " ", u.Comentario, " ", u.Usuario)

	defer client.Close()
}
