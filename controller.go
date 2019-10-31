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

func homeHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("home.html")
	t.Execute(w, nil)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("viewNotes.html")
	t.Execute(w, nil)
}

func createHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("createNote.html")
	t.Execute(w, nil)
}

func updateHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("updateNote.html")
	t.Execute(w, nil)
}
