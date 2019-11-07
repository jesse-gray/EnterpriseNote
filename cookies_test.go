package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//func TestCreateCookie(t *testing.T) {
//	assert := assert.New(t)
//	testcookie := createCookie()
//	assert.NotNil(testcookie)
//}

func TestAttatchCookietoUser(t *testing.T) {
	assert := assert.New(t)

	var testUser User
	testUser.FirstName = "Bob"
	testUser.LastName = "Testcase"
	testUser.Password = "password"
	testUser.CookieID = "1"

	var testNote Note
	testNote.AuthorID = 1
	testNote.NoteText = "Sample Note"

	var testPermissions Permission
	testPermissions.NoteID = 1
	testPermissions.UserID = 1
	testPermissions.ReadPermission = true
	testPermissions.WritePermission = true

	db := opendb()
	defer db.Close()

	if assert.NotNil(db) {
		// test attatchCookietoUser
		assert.True(attatchCookietoUser("a", testUser), "Cookie has been added")
		
		// test createcookie
		testcookie := createCookie()
		assert.NotNil(testcookie)
			
		// test getCookie
		
	}
}
