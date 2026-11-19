package main

type User struct {
	Id    string `firestore:"id"`
	First string `firestore:"first"`
	Last  string `firestore:"last"`
	Born  int    `firestore:"born"`
}
