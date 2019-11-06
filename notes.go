package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Note Struct
type Note struct {
	NoteID   string `json:"noteid"`
	NoteText string `json:"notetext"`
	AuthorID int    `json:"authorid"`
}

//===============Get ALL notes===============

func getNotes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//Make database connection
	db := opendb()
	defer db.Close()

	//Create and execute SQL statement
	sqlStatement := `SELECT DISTINCT note.note_id, note_text, author_id, note_user.cookie_id, write_permission = true OR note_user.cookie_id = $1 FROM note LEFT JOIN permissions ON note.note_id = permissions.note_id JOIN "user" AS note_user ON note.author_id = note_user.user_id LEFT JOIN "user" AS permissions_user ON permissions.user_id = permissions_user.user_id WHERE note_user.cookie_id = $1 OR (permissions_user.cookie_id = $1 AND permissions.read_permission = true)`
	rows, err := db.Query(sqlStatement, getCookie(r))
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	//Declare arrays
	var allNotes [3][]Note
	var myNotes []Note
	var writeNotes []Note
	var readNotes []Note
	var writePerm bool

	//Format results from database
	for rows.Next() {
		var note Note
		var cookie string
		err = rows.Scan(&note.NoteID, &note.NoteText, &note.AuthorID, &cookie, &writePerm)
		if err != nil {
			panic(err)
		}
		//Split into 3 categories(personal, shared/write and shared/read)
		if cookie != "" {
			myNotes = append(myNotes, note)
		} else if writePerm {
			writeNotes = append(writeNotes, note)
		} else {
			readNotes = append(readNotes, note)
		}
	}

	//Add slices to the multidimensional array
	allNotes[0] = myNotes
	allNotes[1] = writeNotes
	allNotes[2] = readNotes

	err = rows.Err()
	if err != nil {
		panic(err)
	}

	//Send back to web app
	json.NewEncoder(w).Encode(allNotes)
}

//===============Get single note===============

func getNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	//Make database connection
	db := opendb()
	defer db.Close()

	//Create and execute SQL statement
	sqlStatement := `SELECT note.note_id, note_text, author_id FROM note LEFT JOIN permissions ON note.note_id = permissions.note_id JOIN "user" AS note_user ON note.author_id = note_user.user_id LEFT JOIN "user" AS permissions_user ON permissions.user_id = permissions_user.user_id WHERE note.note_id = $1 AND (note_user.cookie_id = $2 OR (permissions_user.cookie_id = $2 AND permissions.read_permission = true))`

	//Format results from database
	var note Note
	row := db.QueryRow(sqlStatement, params["id"], getCookie(r))

	//Send back to web app
	switch err := row.Scan(&note.NoteID, &note.NoteText, &note.AuthorID); err {
	case sql.ErrNoRows:
		json.NewEncoder(w).Encode(&note)
	case nil:
		json.NewEncoder(w).Encode(note)
	default:
		panic(err)
	}
}

//===============Create a new note===============

func createNote(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	//Recieve data from web app
	var note Note
	_ = json.NewDecoder(r.Body).Decode(&note)

	//Make database connection
	db := opendb()
	defer db.Close()

	//Create and execute SQL statement
	sqlStatement := `INSERT INTO note (note_text, author_id) SELECT $1 AS note_text, user_id FROM "user" WHERE cookie_id = $2`
	_, err := db.Exec(sqlStatement, note.NoteText, getCookie(r))
	if err != nil {
		panic(err)
	}

	//Check if author wants to use their saved sharing settings
	if params["bool"] == "true" {
		var noteID string

		//Find the ID of the latest Note
		sqlStatement := `SELECT MAX(note_id) FROM note`
		err := db.QueryRow(sqlStatement).Scan(&noteID)

		//Create and execute SQL statement
		sqlStatement = `INSERT INTO permissions SELECT $1 AS note_id, favourite_id, read_permission, write_permission FROM favourites JOIN "user" ON favourites.author_id = "user".user_id WHERE cookie_id = $2`
		_, err = db.Exec(sqlStatement, noteID, getCookie(r))
		if err != nil {
			panic(err)
		}
	}
}

//===============Delete a note===============

func deleteNote(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	//Make database connection
	db := opendb()
	defer db.Close()

	//Create and execute SQL statement
	sqlStatement := `DELETE FROM note JOIN "user" ON note.author_id = "user".user_id WHERE note_id = $1 AND cookie_id = $2`
	_, err := db.Exec(sqlStatement, params["id"], getCookie(r))
	if err != nil {
		panic(err)
	}
}

//===============Update a note===============

func updateNote(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	//Recieve data from web app
	var note Note
	_ = json.NewDecoder(r.Body).Decode(&note)

	//Make database connection
	db := opendb()
	defer db.Close()

	//Create and execute SQL statement
	sqlStatement := `UPDATE note SET note_text = $1 FROM permissions JOIN "user" AS note_user ON permissions.user_id = note_user.user_id JOIN "user" AS permissions_user ON permissions.user_id = permissions_user.user_id WHERE note.note_id = $2 AND (note_user.cookie_id = $3 OR (permissions_user.cookie_id = $3 AND permissions.write_permission = true))`
	_, err := db.Exec(sqlStatement, note.NoteText, params["id"], getCookie(r))
	if err != nil {
		panic(err)
	}
}

//===============Search Notes===============

func searchSQL(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	//Make database connection
	db := opendb()
	defer db.Close()

	//Declare arrays
	var allNotes [3][]Note
	var myNotes []Note
	var writeNotes []Note
	var readNotes []Note
	var note Note

	//Get all notes that match the text and user is allowed to see
	//Create and execute SQL statement
	sqlStatement, err := db.Prepare("SELECT DISTINCT note.note_id, note.note_text, note.author_id, note_user.cookie_id, write_permission = true OR note_user.cookie_id = $1 FROM note LEFT OUTER JOIN permissions ON (note.note_id = permissions.note_id) JOIN \"user\" AS note_user ON note.author_id = note_user.user_id LEFT JOIN \"user\" AS permissions_user ON permissions.user_id = permissions_user.user_id WHERE note_text ~ $2 AND note_user.cookie_id = $1 OR (note_text ~ $2 AND permissions_user.cookie_id = $1 AND (permissions.read_permission = TRUE))")
	if err != nil {
		log.Fatal(err)
	}
	rows, err := sqlStatement.Query(getCookie(r), params["sql"])
	if err != nil {
		log.Fatal(err)
	}

	//Format results from database
	for rows.Next() {
		var cookie string
		var writePerm bool
		err = rows.Scan(&note.NoteID, &note.NoteText, &note.AuthorID, &cookie, &writePerm)
		if err != nil {
			log.Fatal(err)
		}
		if cookie != "" {
			myNotes = append(myNotes, note)
		} else if writePerm {
			writeNotes = append(writeNotes, note)
		} else {
			readNotes = append(readNotes, note)
		}
	}

	//Add slices to the multidimensional array
	allNotes[0] = myNotes
	allNotes[1] = writeNotes
	allNotes[2] = readNotes

	//Send back to web app
	json.NewEncoder(w).Encode(&allNotes)
}

//===============Analyse note===============

func analyseNote(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	//Make database connection
	db := opendb()
	defer db.Close()

	//Check note exists and user has permission
	//Create and execute SQL statement
	var noteExists bool
	sqlStatement := `SELECT EXISTS (SELECT 1 FROM note JOIN "user" AS note_user ON note.author_id = note_user.user_id LEFT JOIN permissions ON note.note_id = permissions.note_id LEFT JOIN "user" AS permissions_user ON permissions.user_id = permissions_user.user_id WHERE note.note_id = $1 AND note_user.cookie_id = $2 OR (note.note_id = $1 AND permissions_user.cookie_id = $2 AND permissions.read_permission = true))`
	err := db.QueryRow(sqlStatement, params["id"], getCookie(r)).Scan(&noteExists)
	if err != nil {
		panic(err)
	}

	//Count occurances
	//Create and execute SQL statement
	sqlStatement = `SELECT (length(str) - length(replace(str, replacestr, '')) )::int / length(replacestr) FROM (VALUES ((SELECT note_text FROM note JOIN "user" AS note_user ON note.author_id = note_user.user_id LEFT JOIN permissions ON note.note_id = permissions.note_id LEFT JOIN "user" AS permissions_user ON permissions.user_id = permissions_user.user_id WHERE note.note_id = $1 AND note_user.cookie_id = $2 OR (note.note_id = $1 AND permissions_user.cookie_id = $2 AND permissions.read_permission = true)), $3)) AS t(str, replacestr)`
	row := db.QueryRow(sqlStatement, params["id"], getCookie(r), params["sql"])

	//Format results from database
	var count int
	err = row.Scan(&count)

	//Send back to web app
	if noteExists {
		json.NewEncoder(w).Encode(count)
	}
}
