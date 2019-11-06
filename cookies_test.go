package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateCookie(t *testing.T) {
	assert := assert.New(t)
	testcookie := createCookie()
	assert.NotNil(testcookie)
}
