package main

import "cloud.google.com/go/firestore"

type Comentario struct {
	ID         string                 `firestore:"id"`
	Comentario string                 `firestore:"comentario"`
	Usuario    *firestore.DocumentRef `firestore:"usuario"`
}
