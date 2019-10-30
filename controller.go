package main

import (
	"html/template"
	"net/http"
)

//Index handler
func indexHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, nil)
}
