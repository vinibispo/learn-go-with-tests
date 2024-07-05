package main

import (
	"log"
	"net/http"
	"os"
)

const dbFileName = "game.db.json"

func main() {
	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		log.Fatalf("problem opening %s %v", dbFileName, err)
	}

	store := NewFileSystemPlayerStore(db)
	server := NewPlayerServer(store)

	//setting 8000 because MacOS Control Center uses 5000
	if err := http.ListenAndServe(":8000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
