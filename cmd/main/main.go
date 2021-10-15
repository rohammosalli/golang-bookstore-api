package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rohammosalli/golang-bookstore-api/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	fmt.Println("Server is starting...")
	err := http.ListenAndServe("localhost:8080", r)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
		fmt.Println("Server stopped")
	}

}
