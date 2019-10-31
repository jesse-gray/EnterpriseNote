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
	r.HandleFunc("/api/home", homeHandler)
	r.HandleFunc("/api/viewNotes", viewNotesHandler)
	r.HandleFunc("/api/viewUsers", viewUsersHandler)
	r.HandleFunc("/api/createNote", createHandler)
	r.HandleFunc("/api/updateNote", updateHandler)
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
	//JavaScript and CSS handlers
	r.PathPrefix("/javascript/").Handler(http.StripPrefix("/javascript/", http.FileServer(http.Dir("./javascript"))))
	r.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir("./css"))))
	log.Fatal(http.ListenAndServe(":8000", r))
}
