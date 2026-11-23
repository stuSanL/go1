package main

import (
	"awesomeProject/dao"
	"context"
	"fmt"
	"log"
	"os"

	"firebase.google.com/go"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

func main() {

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	credentials := os.Getenv("CREDENTIALS_FILE")

	ctx := context.Background()
	sa := option.WithAuthCredentialsFile(option.ServiceAccount, credentials)

	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln("Erro ao criar novo aplicativo: ", err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln("Erro ao instanciar novo Firestore: ", err)
	}

	c, err := dao.ComentarioReadById("com2", ctx, client)
	if err != nil {
		log.Fatalln("Erro ao ler Comentario: ", err)
	}
	c.ID = "com3"
	fmt.Println(c)

	defer client.Close()
}
