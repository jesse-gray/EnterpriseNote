package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

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

//Permission Struct
type Permission struct {
	NoteID          int  `json:"noteid"`
	UserID          int  `json:"userid"`
	ReadPermission  bool `json:"readpermission"`
	WritePermission bool `json:"writepermission"`
}

//Connect to postgres db
func opendb() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}

//Get ALL notes
func getNotes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := opendb()
	defer db.Close()
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
	db := opendb()
	defer db.Close()
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
	db := opendb()
	defer db.Close()
	sqlStatement := `INSERT INTO "note" (note_text, author_id) VALUES ($1, $2)`
	_, err := db.Exec(sqlStatement, note.NoteText, 1) //@todo get author_id from cookie (currently logged on user)
	if err != nil {
		panic(err)
	}
}

//Delete a note
func deleteNote(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	db := opendb()
	defer db.Close()
	sqlStatement := `DELETE FROM "note" WHERE note_id = $1 AND author_id = $2`
	_, err := db.Exec(sqlStatement, params["id"], 1) //@todo get author_id from cookie (currently logged on user)
	if err != nil {
		panic(err)
	}
}

//Update a note
func updateNote(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	db := opendb()
	defer db.Close()
	var note Note
	_ = json.NewDecoder(r.Body).Decode(&note)
	sqlStatement := `UPDATE "note" SET note_text = $1 FROM permissions WHERE note.note_id = $2 AND (author_id = $3 OR (permissions.user_id = $3 AND permissions.write_permission = true))`
	_, err := db.Exec(sqlStatement, note.NoteText, params["id"], 2) //@todo get author_id from cookie (currently logged on user)
	if err != nil {
		panic(err)
	}
}

//Get ALL users
func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := opendb()
	defer db.Close()
	sqlStatement := `SELECT * FROM "user"`
	rows, err := db.Query(sqlStatement)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var users []User
	for rows.Next() {
		var user User
		err = rows.Scan(&user.UserID, &user.FirstName, &user.LastName)
		if err != nil {
			panic(err)
		}
		users = append(users, user)
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(users)
}

//Create a new user
func createUser(w http.ResponseWriter, r *http.Request) {
	var newUser User
	_ = json.NewDecoder(r.Body).Decode(&newUser)
	db := opendb()
	defer db.Close()
	sqlStatement := `INSERT INTO "user" (user_first_name, user_last_name) VALUES ($1, $2)`
	_, err := db.Exec(sqlStatement, newUser.FirstName, newUser.LastName)
	if err != nil {
		panic(err)
	}
}

//Update a user
func updateUser(w http.ResponseWriter, r *http.Request) {
	db := opendb()
	defer db.Close()
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	sqlStatement := `UPDATE "user" SET user_first_name = $1, user_last_name = $2 WHERE user_id = $3`
	_, err := db.Exec(sqlStatement, user.FirstName, user.LastName, 1) //@todo get author_id from cookie (currently logged on user)
	if err != nil {
		panic(err)
	}
}

//Update a permission
// @todo only note author can update
func updatePermission(w http.ResponseWriter, r *http.Request) {
	db := opendb()
	defer db.Close()
	var permission Permission
	_ = json.NewDecoder(r.Body).Decode(&permission)
	sqlStatement := `INSERT INTO permissions VALUES ($3, $4, $1, $2) ON CONFLICT (note_id, user_id) DO UPDATE SET read_permission = $1, write_permission = $2`
	_, err := db.Exec(sqlStatement, permission.ReadPermission, permission.WritePermission, permission.NoteID, permission.UserID)
	if err != nil {
		panic(err)
	}
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
	r.HandleFunc("/api/users", updateUser).Methods("PUT")
	r.HandleFunc("/api/permission", updatePermission).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8000", r))
}
