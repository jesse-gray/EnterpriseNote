package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
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
