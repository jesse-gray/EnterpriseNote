package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateNote(t *testing.T) {

}

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
