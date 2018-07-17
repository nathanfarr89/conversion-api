package main

import (
	"io"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"alive": true}`)
}

func main() {
	//port := ":" + os.Getenv("PORT")
	port := ":7777"
	router := mux.NewRouter()
	router.HandleFunc("/", HealthCheckHandler)
	router.HandleFunc("/weight/{id}", returnWeightConv)
	router.HandleFunc("/temp/{id}", returnTempConv)
	router.HandleFunc("/measure/{id}", returnMeasureConv)
	router.HandleFunc("/feetmeters/{id}", returnFeetMetersConv)
	log.Fatal(http.ListenAndServe(port, handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router)))
}
