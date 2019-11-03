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
	params := mux.Vars(r)
	db := opendb()
	defer db.Close()
	sqlStatement := `SELECT note.note_id, note_text, author_id FROM note LEFT JOIN permissions ON note.note_id = permissions.note_id WHERE author_id = $1 OR (permissions.user_id = $1 AND permissions.read_permission = true)`
	rows, err := db.Query(sqlStatement, params["id"])
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
	row := db.QueryRow(sqlStatement, params["id"])
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
	sqlStatement := `INSERT INTO note (note_text, author_id) VALUES ($1, $2)`
	_, err := db.Exec(sqlStatement, note.NoteText, params["id"])
	if err != nil {
		panic(err)
	}
}

//Delete a note
func deleteNote(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	db := opendb()
	defer db.Close()
	sqlStatement := `DELETE FROM note WHERE note_id = $1`
	_, err := db.Exec(sqlStatement, params["id"])
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
	sqlStatement := `UPDATE note SET note_text = $1 FROM permissions WHERE note.note_id = $2 AND (author_id = $3 OR (permissions.user_id = $3 AND permissions.write_permission = true))`
	_, err := db.Exec(sqlStatement, note.NoteText, params["id"])
	if err != nil {
		panic(err)
	}
}
