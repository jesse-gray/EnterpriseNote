package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

//Note Struct
type Note struct {
	NoteID   string `json:"noteid"`
	NoteText string `json:"notetext"`
	AuthorID int    `json:"authorid"`
}

//Get ALL notes
func getNotes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// params := mux.Vars(r)
	db := opendb()
	defer db.Close()

	//Get cookie
	c, err := r.Cookie("user_id")
	if err != nil {
		panic(err)
	}
	sqlStatement := `SELECT DISTINCT note.note_id, note_text, author_id FROM note LEFT JOIN permissions ON note.note_id = permissions.note_id JOIN "user" AS note_user ON note.author_id = note_user.user_id LEFT JOIN "user" AS permissions_user ON permissions.user_id = permissions_user.user_id WHERE note_user.cookie_id = $1 OR (permissions_user.cookie_id = $1 AND permissions.read_permission = true)`
	rows, err := db.Query(sqlStatement, c.Value)
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
	//Get cookie
	c, err := r.Cookie("user_id")
	if err != nil {
		panic(err)
	}
	sqlStatement := `SELECT note.note_id, note_text, author_id FROM note LEFT JOIN permissions ON note.note_id = permissions.note_id JOIN "user" AS note_user ON note.author_id = note_user.user_id JOIN "user" AS permissions_user ON permissions.user_id = permissions_user.user_id WHERE note.note_id = $1 AND (note_user.cookie_id = $2 OR (permissions_user.cookie_id = $2 AND permissions.read_permission = true))`
	var note Note
	row := db.QueryRow(sqlStatement, params["id"], c.Value)
	switch err := row.Scan(&note.NoteID, &note.NoteText, &note.AuthorID); err {
	case sql.ErrNoRows:
		json.NewEncoder(w).Encode(&note)
	case nil:
		json.NewEncoder(w).Encode(note)
	default:
		panic(err)
	}
}

//Create a new note
func createNote(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var note Note
	_ = json.NewDecoder(r.Body).Decode(&note)
	db := opendb()
	defer db.Close()
	//Get cookie
	c, err := r.Cookie("user_id")
	if err != nil {
		panic(err)
	}
	sqlStatement := `INSERT INTO note (note_text, author_id) SELECT $1 AS note_text, user_id FROM "user" WHERE cookie_id = $2`
	_, err = db.Exec(sqlStatement, note.NoteText, c.Value)
	if err != nil {
		panic(err)
	}

	//Favourites query
	if params["bool"] == "true" {
		var noteID string

		sqlStatement := `SELECT MAX(note_id) FROM note`
		err := db.QueryRow(sqlStatement).Scan(&noteID)

		sqlStatement = `INSERT INTO permissions SELECT $1 AS note_id, favourite_id, read_permission, write_permission FROM favourites JOIN "user" ON favourites.author_id = "user".user_id WHERE cookie_id = $2`
		_, err = db.Exec(sqlStatement, noteID, c.Value)
		if err != nil {
			panic(err)
		}
	}
}

//Delete a note
func deleteNote(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	db := opendb()
	defer db.Close()
	//Get cookie
	c, err := r.Cookie("user_id")
	if err != nil {
		panic(err)
	}
	sqlStatement := `DELETE FROM note JOIN "user" ON note.author_id = "user".user_id WHERE note_id = $1 AND cookie_id = $2`
	_, err = db.Exec(sqlStatement, params["id"], c.Value)
	if err != nil {
		panic(err)
	}
}

//Update a note
func updateNote(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	db := opendb()
	defer db.Close()
	//Get cookie
	c, err := r.Cookie("user_id")
	if err != nil {
		panic(err)
	}
	var note Note
	_ = json.NewDecoder(r.Body).Decode(&note)
	sqlStatement := `UPDATE note SET note_text = $1 FROM permissions JOIN "user" AS note_user ON permissions.user_id = note_user.user_id JOIN "user" AS permissions_user ON permissions.user_id = permissions_user.user_id WHERE note.note_id = $2 AND (note_user.cookie_id = $3 OR (permissions_user.cookie_id = $3 AND permissions.write_permission = true))`
	_, err = db.Exec(sqlStatement, note.NoteText, params["id"], c.Value)
	if err != nil {
		panic(err)
	}
}
