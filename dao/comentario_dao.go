package dao

import (
	"awesomeProject/model"
	"context"
	"errors"
	"log"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

func ComentarioSet(c model.Comentario, ctx context.Context, client *firestore.Client) {
	id, m := c.ToMap()
	_, err := client.Collection("comentário").Doc(id).Set(ctx, m)
	if err != nil {
		log.Fatalln("Erro ao criar Comentário: ", err)
	}
}

func ComentarioReadById(id string, ctx context.Context, client *firestore.Client) (model.Comentario, error) {
	dsnap, err := client.Collection("comentário").Doc(id).Get(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	var cDTO model.ComentarioDTO
	if err = dsnap.DataTo(&cDTO); err != nil {
		log.Fatalln(err)
	}
	cDTO.ID = id
	return cDTO.ToEntity(ctx)
}

func ComentariosFindByUser(u model.User, ctx context.Context, client *firestore.Client) []model.Comentario {
	uref := client.Collection("users").Doc(u.Id)
	snap := client.Collection("comentário").Where("usuario", "==", uref).Documents(ctx)
	comentarios := make([]model.Comentario, 0)
	defer snap.Stop()
	for {
		doc, err := snap.Next()
		if errors.Is(err, iterator.Done) {
			break
		}
		if err != nil {
			log.Println("Erro ao iterar comentários: ", err)
		}
		var cDTO model.ComentarioDTO
		if err = doc.DataTo(&cDTO); err != nil {
			log.Println(err)
		}
		cDTO.ID = doc.Ref.ID
		err = doc.DataTo(&cDTO)
		if err != nil {
			log.Println("Erro ao construir ComentarioDTO: ", err)
		}
		c, err := cDTO.ToEntity(ctx)
		if err != nil {
			log.Println("Erro ao construir Comentario: ", err)
		}
		comentarios = append(comentarios, c)
	}
	return comentarios
}

func ComentarioDeleteById(id string, ctx context.Context, client *firestore.Client) {
	_, err := client.Collection("comentário").Doc(id).Delete(ctx)
	if err != nil {
		log.Println("Erro ao excluir comentário: ", err)
	}
}
