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

func TestSecurity(t *testing.T) {
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
		// test validate user

		assert.Equal(validateUser("1"), true, "User Exists")

		// test check password

		assert.Equal(checkPassword(testUser), true, "Password checked")

	}
}
