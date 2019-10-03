package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//Note Struct (Model)
type Note struct {
	NoteID   string `json:"noteid"`
	NoteText string `json:"notetext"`
	Author   *User  `json:"author"`
}

//User Struct
type User struct {
	UserID    string `json:"userid"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

//Initialise notes slice
var notes []Note

//Get ALL notes
func getNotes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(notes)
}

//Get single note
func getNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	//Loop through notes for the right id
	for _, item := range notes {
		if item.NoteID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Note{})
}

//Create a new note
func createNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var note Note
	_ = json.NewDecoder(r.Body).Decode(&note)
	note.NoteID = strconv.Itoa(rand.Intn(1000000)) //Mock ID - not safe
	notes = append(notes, note)
	json.NewEncoder(w).Encode(note)
}

//Delete a note
func deleteNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range notes {
		if item.NoteID == params["id"] {
			notes = append(notes[:index], notes[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(notes)
}

//Update a note
func updateNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range notes {
		if item.NoteID == params["id"] {
			notes = append(notes[:index], notes[index+1:]...)
			var note Note
			_ = json.NewDecoder(r.Body).Decode(&note)
			note.NoteID = params["id"]
			notes = append(notes, note)
			json.NewEncoder(w).Encode(note)
			return
		}
	}
	json.NewEncoder(w).Encode(notes)
}

func main() {
	//Initialise Router
	r := mux.NewRouter()

	//Mock data - @todo - implement db
	notes = append(notes, Note{NoteID: "1", NoteText: "This is sample text for the first note", Author: &User{UserID: "1", FirstName: "John", LastName: "Smith"}})
	notes = append(notes, Note{NoteID: "2", NoteText: "This is some more sample text, however this is for the second note", Author: &User{UserID: "2", FirstName: "Sharon", LastName: "Tomkins"}})

	//Route handlers
	r.HandleFunc("/api/notes", getNotes).Methods("GET")
	r.HandleFunc("/api/notes/{id}", getNote).Methods("GET")
	r.HandleFunc("/api/notes", createNote).Methods("POST")
	r.HandleFunc("/api/notes/{id}", updateNote).Methods("PUT")
	r.HandleFunc("/api/notes/{id}", deleteNote).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", r))

}
