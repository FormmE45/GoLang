package main

import (
	"log"
)

func main() {
	db, err := NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}
	db.createTableAccount()
	server := newAPIServer(":3000", db)
	server.Run()
}
