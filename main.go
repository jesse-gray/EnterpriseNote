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
	//Web app routing
	r.HandleFunc("/api/", indexHandler)
	r.HandleFunc("/api/signUp", signUpHandler)
	r.HandleFunc("/api/home", homeHandler)
	r.HandleFunc("/api/viewNotes", viewNotesHandler)
	r.HandleFunc("/api/viewUsers", viewUsersHandler)
	r.HandleFunc("/api/createNote", createHandler)
	r.HandleFunc("/api/viewNote", viewNoteHandler)
	r.HandleFunc("/api/updateNote", updateNoteHandler)
	r.HandleFunc("/api/updateUser", updateUserHandler)
	r.HandleFunc("/api/updatePerms", updatePermsHandler)
	r.HandleFunc("/api/findNote", findNoteHandler)
	r.HandleFunc("/api/analyseNote", analyseNoteHandler)
	r.HandleFunc("/api/viewFavourites", viewFavouritesHandler)

	//API routing
	r.HandleFunc("/api/login", secureLogin).Methods("POST")
	r.HandleFunc("/api/logout", logout).Methods("POST")
	r.HandleFunc("/api/notes", getNotes).Methods("GET")
	r.HandleFunc("/api/note/{id}", getNote).Methods("GET")
	r.HandleFunc("/api/notes/{bool}", createNote).Methods("POST")
	r.HandleFunc("/api/notes/{id}", updateNote).Methods("PUT")
	r.HandleFunc("/api/notes/{id}", deleteNote).Methods("DELETE")
	r.HandleFunc("/api/notes/{sql}", searchSQL).Methods("GET")
	r.HandleFunc("/api/notes/{id}/{sql}", analyseNote).Methods("GET")
	r.HandleFunc("/api/users", getUsers).Methods("GET")
	r.HandleFunc("/api/users", createUser).Methods("POST")
	r.HandleFunc("/api/users", deleteUser).Methods("DELETE")
	r.HandleFunc("/api/users", updateUser).Methods("PUT")
	r.HandleFunc("/api/permission", updatePermission).Methods("PUT")
	r.HandleFunc("/api/favourite", getFavourites).Methods("GET")
	r.HandleFunc("/api/favourite", createFavourite).Methods("POST")
	r.HandleFunc("/api/favourite/{id}", deleteFavourite).Methods("DELETE")

	//JavaScript handler
	r.PathPrefix("/javascript/").Handler(http.StripPrefix("/javascript/", http.FileServer(http.Dir("./javascript"))))

	//Start server on port 8000
	log.Fatal(http.ListenAndServe(":8000", r))
}
