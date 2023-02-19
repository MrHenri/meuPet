package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestNewUser(t *testing.T) {
	email := "example@example.com"
	password := "password123"
	phone := "123-456-7890"

	user, err := NewUser(email, password, phone)

	assert.NotNil(t, user, "NewUser should return a non-nil User pointer")
	assert.NoError(t, err, "NewUser should not return an error")

	assert.Equal(t, email, user.Email, "User email should match input email")
	assert.Equal(t, phone, user.Phone, "User phone should match input phone")
	assert.NotEmpty(t, user.ID, "User ID should not be empty")

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	assert.NoError(t, err, "User password should be a valid bcrypt hash of input password")
}

func TestAuthentication(t *testing.T) {
	password := "password123"

	// create a user with a known password hash
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	assert.NoError(t, err, "bcrypt hash generation should not return an error")
	user := &User{Password: string(hash)}

	// test correct password
	assert.True(t, user.Authentication(password), "Authentication should return true for correct password")

	// test incorrect password
	assert.False(t, user.Authentication("wrongpassword"), "Authentication should return false for incorrect password")
}
