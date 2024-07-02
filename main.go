package main

import (
	"fmt"
	"log"
)

func main() {
	db, err := NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v", db)
	// server := newAPIServer(":3000")
	// server.Run()
}
