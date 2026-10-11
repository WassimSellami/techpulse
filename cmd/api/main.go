package main

import (
	"log"
	"net/http"
	"techpulse/internal/handler"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/brief", handler.Brief)

	log.Println("TechPulse listening on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
