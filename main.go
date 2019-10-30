package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	//Initialise Router
	r := mux.NewRouter()

	//Route handlers
	r.HandleFunc("/api/index", indexHandler)
	r.HandleFunc("/api/create", createHandler)
	r.HandleFunc("/api/notes", getNotes).Methods("GET")
	r.HandleFunc("/api/notes/{id}", getNote).Methods("GET")
	r.HandleFunc("/api/notes", createNote).Methods("POST")
	r.HandleFunc("/api/notes/{id}", updateNote).Methods("PUT")
	r.HandleFunc("/api/notes/{id}", deleteNote).Methods("DELETE")
	r.HandleFunc("/api/users", getUsers).Methods("GET")
	r.HandleFunc("/api/users", createUser).Methods("POST")
	r.HandleFunc("/api/users", deleteUser).Methods("DELETE")
	r.HandleFunc("/api/users", updateUser).Methods("PUT")
	r.HandleFunc("/api/permission", updatePermission).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8000", r))
}
