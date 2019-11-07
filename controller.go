package main

import (
	"html/template"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/index.html")
	t.Execute(w, nil)
}

func signUpHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/signUp.html")
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

func viewNoteHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/viewNote.html")
	t.Execute(w, nil)
}

func updateNoteHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/updateNote.html")
	t.Execute(w, nil)
}

func updateUserHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/updateUser.html")
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

func analyseNoteHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/analyseNote.html")
	t.Execute(w, nil)
}

func viewFavouritesHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/viewFavourites.html")
	t.Execute(w, nil)
}
