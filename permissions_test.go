package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUpdatePermission(t *testing.T) {

	var jsonStr = []byte(`{"note_id":1, "user_id":1, "read_permission":TRUE, "write_permission":TRUE}`)

	req, err := http.NewRequest("PUT", "/api/permission", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(updatePermission)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{"note_id":1, "user_id":1, "read_permission":TRUE, "write_permission":TRUE}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestCreateFavourite(t *testing.T) {

	var jsonStr = []byte(`{"note_id":9, "user_id":9, "read_permission":TRUE, "write_permission":TRUE}`)

	req, err := http.NewRequest("Post", "/api/favourite", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(createFavourite)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	expected := `{"note_id":9, "user_id":9, "read_permission":TRUE, "write_permission":TRUE}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestDeleteFavourite(t *testing.T) {

	req, err := http.NewRequest("DELETE", "/api/favourite/{id}", nil)
	if err != nil {
		t.Fatal(err)
	}

	q := req.URL.Query()
	q.Add("user_id", "1")
	q.Add("note_id", "1")
	req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(deleteUser)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	expected := `{"user_first_name":"John", "user_last_name":"Acers", "cookie_id":"", "user_password":"password"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
