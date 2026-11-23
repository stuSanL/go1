package main

import (
	"awesomeProject/dao"
	"context"
	"fmt"
	"log"
	"os"

	"cloud.google.com/go/firestore"
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

	u := dao.UserReadByFirst("Ada", ctx, client)
	us := u[0]
	fmt.Println(us)

	cms := dao.FindComentariosByUser(us, ctx, client)

	fmt.Println(cms)

	defer func(client *firestore.Client) {
		err := client.Close()
		if err != nil {
			log.Println("Erro ao encerrar client: ", err)
		}
	}(client)
}
