package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

var dbc *sql.DB

func main() {
	dbc = dbConnect()
	r := mux.NewRouter()
	r.HandleFunc("/api/word", addNewWord).Methods("POST")
	r.HandleFunc("/api/task", getTask).Methods("POST")
	r.HandleFunc("/api/check", checkTask).Methods("POST")
	r.HandleFunc("/api/lessons", getLessons).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", r))
}
