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
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, nil)
}

func createHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("create.html")
	t.Execute(w, nil)
}
