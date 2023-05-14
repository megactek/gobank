package main

import (
	"flag"
	"fmt"
	"log"
)

func SeedAccount(store storage, fname, lname, passw string) *Account {

	acc, err := NewAccount(fname, lname, passw)
	if err != nil {
		log.Fatal(err)
	}
	if err := store.CreateAccount(acc); err != nil {
		log.Fatal(err)
	}
	fmt.Println("new account -> ", acc.Number)
	return acc
}
func SeedAccounts(s storage) {
	SeedAccount(s, "megac", "sneh", "megacky667")
}

func main() {

	seed := flag.Bool("seed", false, "seed the db")
	flag.Parse()
	store, err := NewPostGresStore()
	if err != nil {
		log.Fatal(err)
	}
	if err := store.Init(); err != nil {
		log.Fatal(err)
	}
	if *seed {
		fmt.Printf("seeding accounts to the database\n")
		SeedAccounts(store)
	}
	server := NewApiServer(":3000", store)
	server.Run()
}
