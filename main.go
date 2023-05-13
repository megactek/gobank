package main

import (
	"log"
)

func main() {
	store, err := NewPostGresStore()
	if err != nil {
		log.Fatal(err)
	}
	if err := store.Init(); err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("%v+\n", store)
	server := NewApiServer(":3000", store)
	server.Run()
}
