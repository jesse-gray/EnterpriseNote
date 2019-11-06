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

//===============Get ALL users from database===============

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//Make db connection
	db := opendb()
	defer db.Close()

	//Create and execute SQL statement
	sqlStatement := `SELECT * FROM "user"`
	rows, err := db.Query(sqlStatement)
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

//===============Create a new user in database===============

func createUser(w http.ResponseWriter, r *http.Request) {
	//Recieve data from web app
	var newUser User
	_ = json.NewDecoder(r.Body).Decode(&newUser)

	//Make database connection
	db := opendb()
	defer db.Close()

	//Create and execute SQL statement
	sqlStatement := `INSERT INTO "user" (user_first_name, user_last_name, user_password) VALUES ($1, $2, $3)`
	_, err := db.Exec(sqlStatement, newUser.FirstName, newUser.LastName, newUser.Password)
	if err != nil {
		panic(err)
	}
}

//===============Delete a user===============

func deleteUser(w http.ResponseWriter, r *http.Request) {
	//Make database connection
	db := opendb()
	defer db.Close()

	//Create and execute SQL statement
	sqlStatement := `DELETE FROM "user" WHERE cookie_id = $1`
	_, err := db.Exec(sqlStatement, getCookie(r))
	if err != nil {
		panic(err)
	}
}

//===============Update a user===============

func updateUser(w http.ResponseWriter, r *http.Request) {
	//Recieve data from web app
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)

	//Make databse connection
	db := opendb()
	defer db.Close()

	//Create and execute SQL statement
	sqlStatement := `UPDATE "user" SET user_first_name = $1, user_last_name = $2, user_password = $3 WHERE cookie_id = $4`
	_, err := db.Exec(sqlStatement, user.FirstName, user.LastName, user.Password, getCookie(r))
	if err != nil {
		panic(err)
	}
}
