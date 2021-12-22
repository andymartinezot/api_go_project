package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("public")) //Creating handler indicating that the folder with the index is "public"
	http.Handle("/", fs)                      //Register handler and calling the fs variable
	log.Println("Server running in http://127.0.0.1:8081")
	http.ListenAndServe(":8081", nil) //Create a server using the 8081 port.
}
