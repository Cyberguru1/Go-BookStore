package main


import (
	"log"
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/Cyberguru1/Go-BookStore/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	fmt.Println("Runnig on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
	
}