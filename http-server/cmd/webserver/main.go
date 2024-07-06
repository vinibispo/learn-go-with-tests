package main

import (
	"log"
	"net/http"
	poker "github.com/vinibispo/learn-go-with-tests/http-server"
)

const dbFileName = "game.db.json"

func main() {
	store, close, err := poker.FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatal(err)
	}
	defer close()

	server := poker.NewPlayerServer(store)

	//setting 8000 because MacOS Control Center uses 5000
	if err := http.ListenAndServe(":8000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
