package main

import (
	"log"
	"net/http"

	"github.com/andymartinezot/api_go_project/project3_crud/handler"
	"github.com/andymartinezot/api_go_project/project3_crud/storage"
)

func main() {
	store := storage.NewMemory()
	mux := http.NewServeMux()

	handler.RoutePerson(mux, &store)

	log.Println("Server initialized at the 8080 port")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Printf("Error in the server: %v\n", err)
	}
}
