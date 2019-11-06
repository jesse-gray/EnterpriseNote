package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOpendb(t *testing.T) {
	assert := assert.New(t)

	var testUser User
	testUser.FirstName = "Bob"
	testUser.LastName = "Testcase"
	testUser.Password = "password"

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

	assert.NotNil(db, "db not opened")

}
