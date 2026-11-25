package dao

import (
	"awesomeProject/model"
	"context"
	"log"

	"cloud.google.com/go/firestore"
)

func JogoSet(j model.Jogo, ctx context.Context, client *firestore.Client) {
	id, mapa := j.ToMap()
	if _, err := client.Collection("jogos").Doc(id).Set(ctx, mapa, firestore.MergeAll); err != nil {
		log.Println("[JogoSet] Erro ao setar jogo: ", err)
	}
}

func JogoReadById(id string, ctx context.Context, client *firestore.Client) model.Jogo {
	dsnap, err := client.Collection("jogos").Doc(id).Get(ctx)
	if err != nil {
		log.Println("[JogoReadById] Erro ao ler jogo: ", err)
	}
	jogo := model.JogoDTO{}
	if err := dsnap.DataTo(&jogo); err != nil {
		log.Println("[JogoReadById] Erro ao ler jogo: ", err)
	}
	jogo.Id = id
	return jogo.ToEntity()
}
