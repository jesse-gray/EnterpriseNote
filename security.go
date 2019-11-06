package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// function to check that entered username is valid must complete this
func validateUser(userid string) bool {
	var user int

	db := opendb()
	defer db.Close()

	sqlStatement, err := db.Prepare("SELECT user_id FROM \"user\" WHERE user_id = $1;")
	if err != nil {
		log.Fatal(err)
	}

	exists := sqlStatement.QueryRow(userid).Scan(&user)
	// if exists does not return a match of userid to the database then no user with this name exists

	if exists == sql.ErrNoRows {
		return false
	}
	if exists != nil {
		log.Fatal(exists.Error())
	}
	return true
}

// function to check a password match may put this into pword.go
func checkPassword(toCheck User) bool {
	var pword string

	db := opendb()
	defer db.Close()

	// using sqlStatement to match use through code

	sqlStatement, err := db.Prepare("SELECT user_password FROM \"user\" WHERE user_id = $1 AND user_password = $2;")
	if err != nil {
		log.Fatal(err)
	}

	err = sqlStatement.QueryRow(toCheck.UserID, toCheck.Password).Scan(&pword)

	// if err == null ie nothing is returned from query, there is no matching password so yjrm er report false
	if err == sql.ErrNoRows {
		return false
	}

	if err != nil {
		log.Fatal(err)
	}

	// we have a match =)

	return true
}

// now to use the above 2 functions to allow a user to login to the application

func secureLogin(w http.ResponseWriter, r *http.Request) {
	var user User
	var result bool
	// var userPassword Pword
	streamUser, err := ioutil.ReadAll(r.Body) // parsing data from a post request
	if err != nil {
		panic(err)
	}

	json.Unmarshal(streamUser, &user) // unmarshal the data from the reader
	// json.Unmarshal(userPassword, &user)

	if validateUser(user.UserID) { //1st checks the user is valid
		if checkPassword(user) { // 2nd checks the passwords match
			CookieID := createCookie()
			if err != nil {
				panic(err)
			}
			attatchCookietoUser(CookieID, user) // sets cookie to db

			userCookie := &http.Cookie{ //creating the cookie for the user_id
				Name:  "user_id",
				Value: CookieID,
			}
			// set the cookie on client
			http.SetCookie(w, userCookie)
			fmt.Printf("Log in was Successful") // console use only
			result = true
			json.NewEncoder(w).Encode(result)
		} else {
			fmt.Printf("Log in failed, incorrect password") // console use only
		}
	} else {
		fmt.Printf("No such user exists") // need to replace with http message to interact with front end
	}
}

// function to execute text search in SQL
func searchSQL(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	db := opendb()
	defer db.Close()
	//Get cookie
	c, err := r.Cookie("user_id")
	if err != nil {
		panic(err)
	}
	var notes []Note
	var note Note
	sqlStatement, err := db.Prepare("SELECT note.note_id, note.note_text, note.author_id FROM note LEFT OUTER JOIN permissions ON (note.note_id = permissions.note_id) JOIN \"user\" AS note_user ON note.author_id = note_user.user_id LEFT JOIN \"user\" AS permissions_user ON permissions.user_id = permissions_user.user_id WHERE note_text ~ $2 AND note_user.cookie_id = $1 OR (note_text ~ $2 AND permissions_user.cookie_id = $1 AND (permissions.read_permission = TRUE))")
	if err != nil {
		log.Fatal(err)
	}
	rows, err := sqlStatement.Query(c.Value, params["sql"])
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		err = rows.Scan(&note.NoteID, &note.NoteText, &note.AuthorID)
		if err != nil {
			log.Fatal(err)
		}
		notes = append(notes, note)
	}
	json.NewEncoder(w).Encode(&notes) // need to add error functionality here
}

//Analyse note function
func analyseNote(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	db := opendb()
	defer db.Close()
	//Get cookie
	c, err := r.Cookie("user_id")
	if err != nil {
		panic(err)
	}
	//Check note exists and user has permission
	var noteExists bool
	sqlStatement := `SELECT EXISTS (SELECT 1 FROM note JOIN "user" AS note_user ON note.author_id = note_user.user_id LEFT JOIN permissions ON note.note_id = permissions.note_id LEFT JOIN "user" AS permissions_user ON permissions.user_id = permissions_user.user_id WHERE note.note_id = $1 AND note_user.cookie_id = $2 OR (note.note_id = $1 AND permissions_user.cookie_id = $2 AND permissions.read_permission = true))`
	err = db.QueryRow(sqlStatement, params["id"], c.Value).Scan(&noteExists)
	//Count occurances
	sqlStatement = `SELECT (length(str) - length(replace(str, replacestr, '')) )::int / length(replacestr) FROM (VALUES ((SELECT note_text FROM note JOIN "user" AS note_user ON note.author_id = note_user.user_id LEFT JOIN permissions ON note.note_id = permissions.note_id LEFT JOIN "user" AS permissions_user ON permissions.user_id = permissions_user.user_id WHERE note.note_id = $1 AND note_user.cookie_id = $2 OR (note.note_id = $1 AND permissions_user.cookie_id = $2 AND permissions.read_permission = true)), $3)) AS t(str, replacestr)`
	row := db.QueryRow(sqlStatement, params["id"], c.Value, params["sql"])
	var count int
	err = row.Scan(&count)
	if noteExists {
		json.NewEncoder(w).Encode(count)
	}
}
