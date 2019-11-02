package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

type Pword struct {
	PwordID 	int 	'json:"pwordID"'
	UserID 		int 	'json:"userID"'
	Pword		string	'json:"pword"'
}

func checkPassword (password string) bool {
	var password string

	db := opendb()
	defe db.Close()

}