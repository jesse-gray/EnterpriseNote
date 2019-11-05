package main

import (
	"encoding/json"
	"net/http"
)

//User Struct
type User struct {
	UserID    string `json:"userid"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	CookieID  string `json:"cookieid"`
	Password  string `json:"password"`
}

//Get ALL users
func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := opendb()
	defer db.Close()
	// sqlStatement := `SELECT user_id, user_first_name, user_last_name, user_password FROM "user"`
	sqlStatement := `SELECT * FROM "user"`
	rows, err := db.Query(sqlStatement)
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

//Create a new user
func createUser(w http.ResponseWriter, r *http.Request) {
	var newUser User
	_ = json.NewDecoder(r.Body).Decode(&newUser)
	db := opendb()
	defer db.Close()
	sqlStatement := `INSERT INTO "user" (user_first_name, user_last_name, user_password) VALUES ($1, $2, $3)`
	_, err := db.Exec(sqlStatement, newUser.FirstName, newUser.LastName, newUser.Password)
	if err != nil {
		panic(err)
	}
}

//Delete a user
func deleteUser(w http.ResponseWriter, r *http.Request) {
	db := opendb()
	defer db.Close()

	//Get cookie
	c, err := r.Cookie("user_id")
	if err != nil {
		panic(err)
	}

	sqlStatement := `DELETE FROM "user" WHERE cookie_id = $1`
	_, err = db.Exec(sqlStatement, c.Value)
	if err != nil {
		panic(err)
	}
}

//Update a user
func updateUser(w http.ResponseWriter, r *http.Request) {
	db := opendb()
	defer db.Close()

	//Get cookie
	c, err := r.Cookie("user_id")
	if err != nil {
		panic(err)
	}

	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	sqlStatement := `UPDATE "user" SET user_first_name = $1, user_last_name = $2, user_password = $3 WHERE cookie_id = $4`
	_, err = db.Exec(sqlStatement, user.FirstName, user.LastName, user.Password, c.Value)
	if err != nil {
		panic(err)
	}
}
