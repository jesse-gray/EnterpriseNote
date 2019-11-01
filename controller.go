package main

import (
	"html/template"
	"net/http"
)

//Index handler
// func indexHandler(w http.ResponseWriter, r *http.Request) {
// 	if /*User already logged on*/ {
// 		t, _ := template.ParseFiles("main.html")
// 	} else {
// 		t, _ := template.ParseFiles("index.html")
// 	}
// 	t.Execute(w, nil)
// }

func indexHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/index.html")
	t.Execute(w, nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/home.html")
	t.Execute(w, nil)
}

func viewNotesHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/viewNotes.html")
	t.Execute(w, nil)
}

func viewUsersHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/viewUsers.html")
	t.Execute(w, nil)
}

func createHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/createNote.html")
	t.Execute(w, nil)
}

func updateHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/updateNote.html")
	t.Execute(w, nil)
}

func updatePermsHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/updatePerms.html")
	t.Execute(w, nil)
}

func findNoteHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/findNote.html")
	t.Execute(w, nil)
}
