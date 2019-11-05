package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
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

//Update a permission
func updatePermission(w http.ResponseWriter, r *http.Request) {
	db := opendb()
	defer db.Close()
	//Get cookie
	c, err := r.Cookie("user_id")
	if err != nil {
		panic(err)
	}
	var permission Permission
	fmt.Println(permission)
	_ = json.NewDecoder(r.Body).Decode(&permission)
	//Find author of note
	var cookie string
	sqlStatement := `SELECT cookie_id FROM "user" JOIN note ON "user".user_id = note.author_id WHERE note_id = $1`
	err = db.QueryRow(sqlStatement, permission.NoteID).Scan(&cookie)
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}
	fmt.Println(cookie)
	//Check user to be updated exists
	var userExists bool
	sqlStatement = `SELECT EXISTS (SELECT 1 FROM "user" WHERE user_id = $1)`
	err = db.QueryRow(sqlStatement, permission.UserID).Scan(&userExists)
	if err != nil {
		panic(err)
	}
	fmt.Println(userExists)
	//Execute the update if user exists and current user is author
	if cookie == c.Value && userExists { //@todo get author_id from cookie (currently logged on user)
		sqlStatement = `INSERT INTO permissions VALUES ($3, $4, $1, $2) ON CONFLICT (note_id, user_id) DO UPDATE SET read_permission = $1, write_permission = $2`
		_, err = db.Exec(sqlStatement, permission.ReadPermission, permission.WritePermission, permission.NoteID, permission.UserID)
		if err != nil {
			panic(err)
		}
	}
}

//Get all favourites
func getFavourites(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := opendb()
	defer db.Close()
	//Get cookie
	c, err := r.Cookie("user_id")
	if err != nil {
		panic(err)
	}
	sqlStatement := `SELECT "user".user_id, "user".user_first_name, "user".user_last_name, "user".cookie_id, "user".user_password FROM "user" JOIN favourites ON "user".user_id = favourites.favourite_id JOIN "user" AS author ON favourites.author_id = author.user_id WHERE author.cookie_id = $1`
	rows, err := db.Query(sqlStatement, c.Value)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
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
	json.NewEncoder(w).Encode(users)

}

//Create a new favourite
func createFavourite(w http.ResponseWriter, r *http.Request) {
	var newPermission Permission
	_ = json.NewDecoder(r.Body).Decode(&newPermission)
	db := opendb()
	defer db.Close()
	//Get cookie
	c, err := r.Cookie("user_id")
	if err != nil {
		panic(err)
	}
	//Get currently logged in user ID
	var authorID int
	sqlStatement := `SELECT user_id FROM "user" WHERE cookie_id = $1`
	row := db.QueryRow(sqlStatement, c.Value)
	err = row.Scan(&authorID)
	//Insert into table
	sqlStatement = `INSERT INTO favourites VALUES ($1, $2, $3, $4)`
	_, err = db.Exec(sqlStatement, authorID, newPermission.UserID, newPermission.ReadPermission, newPermission.WritePermission)
	if err != nil {
		panic(err)
	}
}

//Delete a favourite
func deleteFavourite(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	db := opendb()
	defer db.Close()
	//Get cookie
	c, err := r.Cookie("user_id")
	if err != nil {
		panic(err)
	}
	sqlStatement := `DELETE FROM favourites USING "user" WHERE favourites.author_id = "user".user_id AND cookie_id = $1 AND favourite_id = $2`
	_, err = db.Exec(sqlStatement, c.Value, params["id"])
	if err != nil {
		panic(err)
	}
}
