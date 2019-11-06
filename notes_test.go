package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetNote(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/note/{id}/{user}", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getNote)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	//check the response body is what we expect.

	expected := `[{"note_id":1, "note_text":"This is sample text for the first note", "author_id":1}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected data got %v want %v", rr.Body.String(), expected)
	}
}

func TestCreateNote(t *testing.T) {

	var jsonStr = []byte(`{"note_text":"Hello World", "author_id":1}`)

	req, err := http.NewRequest("Post", "/api/notes/{bool}", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(createNote)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	expected := `{"note_text":"Hello World", "author_id":1}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestDeleteNote(t *testing.T) {

	req, err := http.NewRequest("DELETE", "/entry", nil)
	if err != nil {
		t.Fatal(err)
	}

	q := req.URL.Query()
	q.Add("note_id", "1")
	req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(deleteNote)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	expected := `{"note_id":1,"note_text":"This is a sample text for the first note","author_id":1}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestUpdateNote(t *testing.T) {

	var jsonStr = []byte(`{"note_id":1, "note_text":"Hello World", "author_id":1}`)

	req, err := http.NewRequest("PUT", "/api/notes/{id}", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(UpdateNote)
	handler := ServeHttp(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{"note_id":1, "note_text":"Hello World", "author_id":1}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

// func TestGetNotes -- Need to find out how to test recieving an array of objects
