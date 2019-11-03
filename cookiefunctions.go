package main

import (
	"database/sql"
	"log"
	"net/http"
	"github.com/lib/pq"
	"github.com/satori/go.uuid"
)

// Use the uuid library to create 128 bit numbers for session id's, these will be unique

// function to create the cookie uses the uuid library to give it a unique id, based on time stamp and MAC address

func createCookie() string {
	uuid.Init()
	cookieID := uuid.NewV1()
	// testing fmt.Printf("version %d variant %d: %d\n", u1.Version(), u1.Variant(), u1)
	// return id for use
	return cookieID.String()
}

// function adds the session to the logged in user on the backend

func attatchCookietoUser(cookieID string, user User) bool {
	db := opendb()
	defer db.Close()

	// add cookie to database to show user is logged in
	sqlStatement, err := db.Prepare("UPDATE "user" SET cookie_id = $1 WHERE user_id = $2;")
	if err != nil {
		panic(err)
	}
	sqlStatement.Exec(cookieID, user.UserID)

	// show success of cookie placement
	return true
}

// function to remove the session id from the user when no longer logged in

func removeCookieFromUser(w http.ResponseWriter, r *http.Request) {
	// get cookieID
	cookieID, err := r.Cookie("cookie_ID")

	cookieID = &http.Cookie{
		Name:	"session",
		Value:	"",
		Age:	-1,
	}
	// set cookie using http.Setcookie with cookie ID which is now blank
	http.SetCookie(w, cookieID)
}

// a function to see if the user is logged in by getting the cookie
func getCookie(r *http.Request) (cookieID string) {

	cookieTracer, err := r.Cookie("cookie_id")
	if err != nil {
		cookieID = " "
		return cookieID
	}
	cookieID = cookieTracer.Value
	return cookieID
}