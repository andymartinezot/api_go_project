package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	//Register the path and the handler
	http.HandleFunc("/Saludar", Saludar)
	http.HandleFunc("/Despedir", Despedir)
	//Raising a server using the 8080 port
	log.Println("Server running on http://127.0.0.1:8080")
	http.ListenAndServe(":8080", nil)
}

func Saludar(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world")
}

func Despedir(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Bye world")
}
