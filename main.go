package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homepage")
}

func main() {
	port := ":" + os.Getenv("PORT")
	router := mux.NewRouter()
	router.HandleFunc("/", homePage)
	router.HandleFunc("/weight/{id}", returnWeightConv)
	router.HandleFunc("/temp/{id}", returnTempConv)
	router.HandleFunc("/measure/{id}", returnMeasureConv)
	router.HandleFunc("/feetmeters/{id}", returnFeetMetersConv)
	log.Fatal(http.ListenAndServe(port, handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router)))
}
