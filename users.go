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
	Password  string `json:"password"`
}

//Login
func logIn(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")
	// db := opendb()
	// defer db.Close()
	// sqlStatement := `SELECT * FROM "user"`
	// rows, err := db.Query(sqlStatement)
	// if err != nil {
	// 	panic(err)
	// }
	// defer rows.Close()
	// var users []User
	// for rows.Next() {
	// 	var user User
	// 	err = rows.Scan(&user.UserID, &user.FirstName, &user.LastName)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	users = append(users, user)
	// }
	// err = rows.Err()
	// if err != nil {
	// 	panic(err)
	// }
	// json.NewEncoder(w).Encode(users)
}

//Sign Up
func signUp(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")
	// db := opendb()
	// defer db.Close()
	// sqlStatement := `SELECT * FROM "user"`
	// rows, err := db.Query(sqlStatement)
	// if err != nil {
	// 	panic(err)
	// }
	// defer rows.Close()
	// var users []User
	// for rows.Next() {
	// 	var user User
	// 	err = rows.Scan(&user.UserID, &user.FirstName, &user.LastName)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	users = append(users, user)
	// }
	// err = rows.Err()
	// if err != nil {
	// 	panic(err)
	// }
	// json.NewEncoder(w).Encode(users)
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
	db := opendb()
	defer db.Close()
	sqlStatement := `DELETE FROM "user" WHERE user_id = $1`
	_, err := db.Exec(sqlStatement, 1) //@todo get author_id from cookie (currently logged on user)
	if err != nil {
		panic(err)
	}
}

//Update a user
func updateUser(w http.ResponseWriter, r *http.Request) {
	db := opendb()
	defer db.Close()
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	sqlStatement := `UPDATE "user" SET user_first_name = $1, user_last_name = $2, user_password = $3 WHERE user_id = $3`
	_, err := db.Exec(sqlStatement, user.FirstName, user.LastName, user.Password, 1) //@todo get author_id from cookie (currently logged on user)
	if err != nil {
		panic(err)
	}
}
