package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
)

// Use the uuid library to create 128 bit numbers for session id's, these will be unique
// cookies will be used to ensure a valid user is using the program can evaluate his notes
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
	cookieID, err := r.Cookie("_cookie")
	// set cookie value empty
	cookieID = &http.Cookie{
		Name:	"_cookie",
		Value:	"",
		Age:	-1,
	}
	// set cookie using http.Setcookie with cookie ID which is now blank
	http.SetCookie(w, cookieID)
}

// a function to see if the user is logged in by getting the cookie
func getCookie(r *http.Request) (cookieID string) {

	cookieTracer, err := r.Cookie("_cookie")
	if err != nil {		// if error occurs return nothing as ID
		cookieID = " "
		return cookieID
	}
	cookieID = cookieTracer.Value	// return cookie from function on successful read
	return cookieID
}

// a function to remove cookie from client (to use in secure logout function)
func deleteCookie(w http.ResponseWriter, r *http.Request) {
	cookieID, _ := r.Cookie("_cookie")
	
	// reads the cookie
	// set cookie value empty
	cookieID = &http.Cookie{
		Name: "_cookie"
		Value: "",
		Age: -1,
	}
	http.SetCookie(w, cookieID) // sets cookie value on client header to empty
}

// function to return userID if using cookieID

func findUserID(req *http.Request) (user_id int) {
	db := opendb()
	defer db.Close()
	// get the cookieID
	cookieTracer, err := r.Cookie("_cookie")
	if err != nil {		// if error occurs return nothing as ID
		cookieTracer = cookieID
	}
	cookieID := cookieTracer.Value
	// sql search of user table for matching cookie
	sqlStatement := `SELECT user_id FROM "user" WHERE cookie_id=$1`
	rows, err := db.Query(sqlStatement, cookieID)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		err = rows.Scan(&user_id)
		if err!=nil{
			panic(err)
		}
	}
	return user_id
}

func isUseridLoggedIn(req *http.Request) bool {
	
	var user_id int
	cookieID, err := req.Cookie("_cookie")
	if err != nil {
		return false //user is not logged in
	}

	db := opendb()
	defer db.Close()
	// statement to pull userid where cookies match
	sqlStatement := `SELECT user_id FROM "user" WHERE cookie_id=$1`

	// if no rows match
	row = db.QueryRow(sqlStatement, cookieID.Value) // pull user id row
	switch err := row.Scan(&user_id); err {
	case sql.ErrNoRows:
		return false // no matches user not logged in
	case nil:
		return true // a match user is logged in
	default:
		panic(err)
	}
}

	// think that is all the cookie functionality we need
	// use the isUseridLoggedIn function for security checks before  user can do anything, ie put a if statement before edit view notes etc
	// also added below function to log out to be added to users




