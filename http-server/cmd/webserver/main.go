package main

import (
	poker "github.com/vinibispo/learn-go-with-tests/http-server"
	"log"
	"net/http"
)

const dbFileName = "game.db.json"

func main() {
	store, close, err := poker.FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatal(err)
	}
	defer close()

	game := poker.NewTexasHoldem(poker.BlindAlerterFunc(poker.Alerter), store)

	server, err := poker.NewPlayerServer(store, game)

	if err != nil {
		log.Fatal(err)
	}

	//setting 8000 because MacOS Control Center uses 5000
	if err := http.ListenAndServe(":8000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
