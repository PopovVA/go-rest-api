package model

import (
	"testing"
)

//TestUser is simple user for unit tests
func TestUser(t *testing.T) *User {
	return &User{
		Email:    "user@example.org",
		Password: "password",
	}
}
