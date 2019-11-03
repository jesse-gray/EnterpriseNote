package main

import (
	"encoding/json"
	"net/http"
)

// function to open database from jesse's code for testing

function opendbtest() *sql.DB {
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




// function to check that entered username is valid must complete this
func validateUser(user string) bool {
	var user string

	db := opendb()
	defer db.Close()
}

// function to execute text search in SQL



function searchSQL(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	db := opendb()
	defer db.Close()
	var notes []Note
	var anote Note
	sqlStatement := db.Prepare("SELECT note.note_id, note.note_text, note.author_id FROM note LEFT OUTER JOIN permissions ON (note.note_id = permissions.note_id) WHERE note_text ~ $2 AND note.author_id = $1 OR (permissions.user_id = $1 AND (permissions.read_permission = TRUE))")
	if err != nil {
		log.Fatal(err)
	}
	rows, err := sqlStatement.Query(getauthorIDcookie, searchtext) // need to figure out where to get text we are searching for
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next(){
		err = rows.Scan(&anote.NoteID, &anote.NoteText, &anote.AuthorID)
		if err != nil {
			log.Fatal(err)
		}
		notes = append(notes, note)
	}
	json.NewEncoder(w).Encode(&notes) // need to add error functionality here

}


