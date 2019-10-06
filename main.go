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
	AuthorID int    `json:"authorid"`
}

//User Struct
type User struct {
	UserID    string `json:"userid"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

//Get ALL notes
func getNotes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
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
	sqlStatement := `SELECT note.note_id, note_text, author_id FROM note LEFT JOIN permissions ON note.note_id = permissions.note_id WHERE author_id = $1 OR (permissions.user_id = $1 AND permissions.read_permission = true)`
	rows, err := db.Query(sqlStatement, 1) //@todo get author_id from cookie (currently logged on user)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var notes []Note
	for rows.Next() {
		var note Note
		err = rows.Scan(&note.NoteID, &note.NoteText, &note.AuthorID)
		if err != nil {
			panic(err)
		}
		notes = append(notes, note)
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(notes)
}

//Get single note
func getNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
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
	sqlStatement := `SELECT note.note_id, note_text, author_id FROM note LEFT JOIN permissions ON note.note_id = permissions.note_id WHERE note.note_id = $1 AND (author_id = $2 OR (permissions.user_id = $2 AND permissions.read_permission = true))`
	var note Note
	row := db.QueryRow(sqlStatement, params["id"], 1) //@todo get author_id from cookie (currently logged on user)
	switch err := row.Scan(&note.NoteID, &note.NoteText, &note.AuthorID); err {
	case sql.ErrNoRows:
		fmt.Println("No notes were found!")
	case nil:
		json.NewEncoder(w).Encode(note)
	default:
		panic(err)
	}
}

//Create a new note
func createNote(w http.ResponseWriter, r *http.Request) {
	var note Note
	_ = json.NewDecoder(r.Body).Decode(&note)
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
	sqlStatement := `INSERT INTO "note" (note_text, author_id) VALUES ($1, $2)`
	_, err = db.Exec(sqlStatement, note.NoteText, 1) //@todo get author_id from cookie (currently logged on user)
	if err != nil {
		panic(err)
	}
}

//Delete a note
func deleteNote(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
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
	sqlStatement := `DELETE FROM "note" WHERE note_id = $1 AND author_id = $2`
	_, err = db.Exec(sqlStatement, params["id"], 1) //@todo get author_id from cookie (currently logged on user)
	if err != nil {
		panic(err)
	}
}

//Update a note
func updateNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	// for index, item := range notes {
	// 	if item.NoteID == params["id"] {
	// 		notes = append(notes[:index], notes[index+1:]...)
	// 		var note Note
	// 		_ = json.NewDecoder(r.Body).Decode(&note)
	// 		note.NoteID = params["id"]
	// 		notes = append(notes, note)
	// 		json.NewEncoder(w).Encode(note)
	// 		return
	// 	}
	// }
	//json.NewEncoder(w).Encode(notes)
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
	var note Note
	_ = json.NewDecoder(r.Body).Decode(&note)
	sqlStatement := `UPDATE "note" SET note_text = $1 WHERE note_id = $2 AND author_id = $2`
	_, err = db.Exec(sqlStatement, params["id"], 1) //@todo get author_id from cookie (currently logged on user)
	if err != nil {
		panic(err)
	}
}

//Get ALL users
var users []User

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
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
	//Initialise Router
	r := mux.NewRouter()

	//Route handlers
	r.HandleFunc("/api/notes", getNotes).Methods("GET")
	r.HandleFunc("/api/notes/{id}", getNote).Methods("GET")
	r.HandleFunc("/api/notes", createNote).Methods("POST")
	r.HandleFunc("/api/notes/{id}", updateNote).Methods("PUT")
	r.HandleFunc("/api/notes/{id}", deleteNote).Methods("DELETE")
	r.HandleFunc("/api/users", getUsers).Methods("GET")
	r.HandleFunc("/api/users", createUser).Methods("POST")
	r.HandleFunc("/api/users/{id}", updateUser).Methods("PUT")
	r.HandleFunc("/api/users/{id}", deleteUser).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", r))
}
