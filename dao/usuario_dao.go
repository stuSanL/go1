package dao

import (
	"awesomeProject/model"
	"context"
	"errors"
	"log"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
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

func UserReadByFirst(first string, ctx context.Context, client *firestore.Client) []model.User {
	dsnap := client.Collection("users").Where("first", "==", first).Documents(ctx)
	users := make([]model.User, 0)
	defer dsnap.Stop()
	for {
		doc, err := dsnap.Next()
		if errors.Is(err, iterator.Done) {
			break
		}
		if err != nil {
			log.Println("Erro ao iterar snapshots: ", err)
		}
		user := model.UserDTO{}
		if err := doc.DataTo(&user); err != nil {
			return nil
		}
		user.Id = doc.Ref.ID
		users = append(users, user.ToEntity())
	}
	return users
}

func UserReplace(u model.User, ctx context.Context, client *firestore.Client) {
	id, m := u.ToMap()
	_, err := client.Collection("users").Doc(id).Set(ctx, m, firestore.MergeAll)
	if err != nil {
		log.Println(err)
	}
}

func UserUpdateBornDate(u model.User, ctx context.Context, client *firestore.Client) {
	_, err := client.Collection("users").Doc(u.Id).Update(ctx, []firestore.Update{{Path: "born", Value: u.Born}})
	if err != nil {
		log.Println("Erro ao atualizar usuário: ", err)
	}
}

func UserDeleteById(id string, ctx context.Context, client *firestore.Client) {
	_, err := client.Collection("users").Doc(id).Delete(ctx)
	if err != nil {
		log.Println("Erro ao deletar usuario: ", err)
	}
}
