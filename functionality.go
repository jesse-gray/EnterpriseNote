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
func 