package main

import (
	"log"
	"net/http"
)

func main() {
	server := &PlayerServer{}
	handler := http.HandlerFunc(server.ServeHTTP)
	log.Fatal(http.ListenAndServe(":8000", handler)) //setting 8000 because MacOS Control Center uses 5000
}
