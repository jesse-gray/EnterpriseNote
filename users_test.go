package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateUser(t *testing.T) {

	var jsonStr = []byte(`{"user_first_name":"Harlod", "user_last_name":"Boomkin", "cookie_id":"", "user_password":"password"}`)

	req, err := http.NewRequest("Post", "/api/notes/{bool}", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(createUser)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	expected := `{"user_first_name":"Harlod", "user_last_name":"Boomkin", "cookie_id":"", "user_password":"password"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestDeleteUser(t *testing.T) {

	req, err := http.NewRequest("DELETE", "/entry", nil)
	if err != nil {
		t.Fatal(err)
	}

	q := req.URL.Query()
	q.Add("user_id", "1")
	req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(deleteNote)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	expected := `{"user_first_name":"John", "user_last_name":"Acers", "cookie_id":"", "user_password":"password"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestUpdateNote(t *testing.T) {

	var jsonStr = []byte(`{"user_id":1, "user_first_name":"Gerald", "user_last_name":"Hopkins", "cookie_id":"", "user_password":"password"}`)

	req, err := http.NewRequest("PUT", "/api/notes/{id}", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(updateNote)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{"user_id":1, "user_first_name":"Gerald", "user_last_name":"Hopkins", "cookie_id":"", "user_password":"password"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

// need to add test for getUsers
