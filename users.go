package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

//User Struct
type User struct {
	UserID    string `json:"userid"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Password  string `json:"password"`
}

//Login
func login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	db := opendb()
	defer db.Close()
	sqlStatement := `SELECT user_id FROM "user" WHERE user_id = $1 AND user_password = $2`
	row := db.QueryRow(sqlStatement, user.UserID, user.Password)
	var logon int
	switch err := row.Scan(&logon); err {
	case sql.ErrNoRows:
		json.NewEncoder(w).Encode(&logon)
	case nil:
		json.NewEncoder(w).Encode(logon)
	default:
		panic(err)
	}
}

//Sign Up
func signUp(w http.ResponseWriter, r *http.Request) {

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
		err = rows.Scan(&user.UserID, &user.FirstName, &user.LastName, &user.Password)
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
	sqlStatement := `INSERT INTO "user" (user_first_name, user_last_name, user_password) VALUES ($1, $2 $3)`
	_, err := db.Exec(sqlStatement, newUser.FirstName, newUser.LastName, newUser.Password)
	if err != nil {
		panic(err)
	}
}

//Delete a user
func deleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	db := opendb()
	defer db.Close()
	sqlStatement := `DELETE FROM "user" WHERE user_id = $1`
	_, err := db.Exec(sqlStatement, params["id"])
	if err != nil {
		panic(err)
	}
}

//Update a user
func updateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	db := opendb()
	defer db.Close()
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	sqlStatement := `UPDATE "user" SET user_first_name = $1, user_last_name = $2, user_password = $3 WHERE user_id = $4`
	_, err := db.Exec(sqlStatement, user.FirstName, user.LastName, user.Password, params["id"])
	if err != nil {
		panic(err)
	}
}
