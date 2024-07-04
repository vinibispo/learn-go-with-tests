package main

import (
	"log"
	"net/http"
)

func main() {
	handler := http.HandlerFunc(PlayerServer)
	log.Fatal(http.ListenAndServe(":8000", handler)) //setting 8000 because MacOS Control Center uses 5000
}
