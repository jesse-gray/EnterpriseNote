package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

//Permission Struct
type Permission struct {
	NoteID          int  `json:"noteid"`
	UserID          int  `json:"userid"`
	ReadPermission  bool `json:"readpermission"`
	WritePermission bool `json:"writepermission"`
}

//===============Update a permission===============

func updatePermission(w http.ResponseWriter, r *http.Request) {
	//Recieve data from web app
	var permission Permission
	_ = json.NewDecoder(r.Body).Decode(&permission)

	//Make databse connection
	db := opendb()
	defer db.Close()

	//Find author of note
	//Create and execute SQL statement
	sqlStatement := `SELECT cookie_id FROM "user" JOIN note ON "user".user_id = note.author_id WHERE note_id = $1`
	var cookie string
	err := db.QueryRow(sqlStatement, permission.NoteID).Scan(&cookie)
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}

	//Check user to be updated exists
	//Create and execute SQL statement
	sqlStatement = `SELECT EXISTS (SELECT 1 FROM "user" WHERE user_id = $1)`
	var userExists bool
	err = db.QueryRow(sqlStatement, permission.UserID).Scan(&userExists)
	if err != nil {
		panic(err)
	}

	//Execute the update if user exists and current user is author
	if cookie == getCookie(r) && userExists {
		//Create and execute SQL statement
		sqlStatement = `INSERT INTO permissions VALUES ($3, $4, $1, $2) ON CONFLICT (note_id, user_id) DO UPDATE SET read_permission = $1, write_permission = $2`
		_, err = db.Exec(sqlStatement, permission.ReadPermission, permission.WritePermission, permission.NoteID, permission.UserID)
		if err != nil {
			panic(err)
		}
	}
}

//===============Get all favourites===============

func getFavourites(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//Make databse connection
	db := opendb()
	defer db.Close()

	//Create and execute SQL statement
	sqlStatement := `SELECT "user".user_id, "user".user_first_name, "user".user_last_name, "user".cookie_id, "user".user_password FROM "user" JOIN favourites ON "user".user_id = favourites.favourite_id JOIN "user" AS author ON favourites.author_id = author.user_id WHERE author.cookie_id = $1`
	rows, err := db.Query(sqlStatement, getCookie(r))
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	//Format results from database
	var users []User
	for rows.Next() {
		var user User
		err = rows.Scan(&user.UserID, &user.FirstName, &user.LastName, &user.CookieID, &user.Password)
		if err != nil {
			panic(err)
		}
		users = append(users, user)
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}

	//Send back to web app
	json.NewEncoder(w).Encode(users)

}

//===============Create a new favourite===============

func createFavourite(w http.ResponseWriter, r *http.Request) {
	//Recieve data from web app
	var newPermission Permission
	_ = json.NewDecoder(r.Body).Decode(&newPermission)

	//Make databse connection
	db := opendb()
	defer db.Close()

	//Get currently logged in user ID
	//Create and execute SQL statement
	sqlStatement := `SELECT user_id FROM "user" WHERE cookie_id = $1`
	row := db.QueryRow(sqlStatement, getCookie(r))

	//Format results from database
	var authorID int
	err := row.Scan(&authorID)

	//Insert into table
	//Create and execute SQL statement
	sqlStatement = `INSERT INTO favourites VALUES ($1, $2, $3, $4)`
	_, err = db.Exec(sqlStatement, authorID, newPermission.UserID, newPermission.ReadPermission, newPermission.WritePermission)
	if err != nil {
		panic(err)
	}
}

//===============Delete a favourite===============

func deleteFavourite(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	//Make databse connection
	db := opendb()
	defer db.Close()

	//Create and execute SQL statement
	sqlStatement := `DELETE FROM favourites USING "user" WHERE favourites.author_id = "user".user_id AND cookie_id = $1 AND favourite_id = $2`
	_, err := db.Exec(sqlStatement, getCookie(r), params["id"])
	if err != nil {
		panic(err)
	}
}
