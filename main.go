package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "enterprisedb"
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

//Initialise slices
var notes []Note
var users []User

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

//Get ALL users
func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

//Get single user
func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	//Loop through users for the right id
	for _, user := range users {
		if user.UserID == params["id"] {
			json.NewEncoder(w).Encode(user)
			return
		}
	}
	json.NewEncoder(w).Encode(&User{})
}

//Create a new user
func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	user.UserID = strconv.Itoa(rand.Intn(1000000)) //Mock ID - not safe
	users = append(users, user)
	json.NewEncoder(w).Encode(user)
}

//Delete a user
func deleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, user := range users {
		if user.UserID == params["id"] {
			users = append(users[:index], users[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(users)
}

//Update a user
func updateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, user := range users {
		if user.UserID == params["id"] {
			users = append(users[:index], users[index+1:]...)
			var user User
			_ = json.NewDecoder(r.Body).Decode(&user)
			user.UserID = params["id"]
			users = append(users, user)
			json.NewEncoder(w).Encode(user)
			return
		}
	}
	json.NewEncoder(w).Encode(users)
}

func main() {
	//Connect to postgres db
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	//Initialise Router
	r := mux.NewRouter()

	//Mock data - @todo - implement db
	notes = append(notes, Note{NoteID: "1", NoteText: "This is sample text for the first note", Author: &User{UserID: "1", FirstName: "John", LastName: "Smith"}})
	notes = append(notes, Note{NoteID: "2", NoteText: "This is some more sample text, however this is for the second note", Author: &User{UserID: "2", FirstName: "Sharon", LastName: "Tomkins"}})

	users = append(users, User{UserID: "1", FirstName: "John", LastName: "Smith"})
	users = append(users, User{UserID: "2", FirstName: "Sharon", LastName: "Tomkins"})

	//Route handlers
	r.HandleFunc("/api/notes", getNotes).Methods("GET")
	r.HandleFunc("/api/notes/{id}", getNote).Methods("GET")
	r.HandleFunc("/api/notes", createNote).Methods("POST")
	r.HandleFunc("/api/notes/{id}", updateNote).Methods("PUT")
	r.HandleFunc("/api/notes/{id}", deleteNote).Methods("DELETE")
	r.HandleFunc("/api/users", getUsers).Methods("GET")
	r.HandleFunc("/api/users/{id}", getUser).Methods("GET")
	r.HandleFunc("/api/users", createUser).Methods("POST")
	r.HandleFunc("/api/users/{id}", updateUser).Methods("PUT")
	r.HandleFunc("/api/users/{id}", deleteUser).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", r))
}
