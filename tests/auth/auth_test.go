package auth_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/mahinops/secretcli-web/cmd"
	"github.com/stretchr/testify/assert"
)

const testConfigPath = "./../../.env.test"

// TestUserRegistration tests the user registration endpoint
func TestUserRegistration(t *testing.T) {
	// Initialize the application
	app, err := cmd.NewApp(testConfigPath)
	assert.NoError(t, err)
	defer app.CloseDatabase()

	// Generate a unique email for this test run
	email := fmt.Sprintf("testuser%d@example.com", time.Now().UnixNano())

	// Prepare the registration request with the correct JSON key
	user := map[string]string{
		"name":     "testuser",
		"email":    email, // Use unique email
		"password": "password123",
	}
	jsonUser, _ := json.Marshal(user)

	req, err := http.NewRequest(http.MethodPost, "/auth/register", bytes.NewBuffer(jsonUser))
	assert.NoError(t, err)

	// Create a ResponseRecorder to capture the response
	rr := httptest.NewRecorder()

	// Serve the request using the Router
	app.Router.ServeHTTP(rr, req)

	// Check the response code
	if rr.Code != http.StatusCreated {
		t.Errorf("Expected status %d but got %d. Response: %s", http.StatusCreated, rr.Code, rr.Body.String())
		return
	}

	var response map[string]interface{}
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Contains(t, response, "message")
	assert.Equal(t, response["message"], "User "+user["name"]+" created successfully")
}

// TestUserLogin tests the user login endpoint
func TestUserLogin(t *testing.T) {
	// Initialize the application
	app, err := cmd.NewApp(testConfigPath)
	assert.NoError(t, err)
	defer app.CloseDatabase()

	// First, register the user
	TestUserRegistration(t) // Ensure the user is registered

	// Prepare the login request
	loginUser := map[string]string{
		"email":    "testuser1@example.com", // Make sure this matches the email used during registration
		"password": "password123",
	}
	jsonUser, _ := json.Marshal(loginUser)

	req, err := http.NewRequest(http.MethodPost, "/auth/login", bytes.NewBuffer(jsonUser))
	assert.NoError(t, err)

	// Create a ResponseRecorder to capture the response
	rr := httptest.NewRecorder()

	// Serve the request using the Router
	app.Router.ServeHTTP(rr, req)

	// Check the response code and body
	assert.Equal(t, http.StatusOK, rr.Code)

	var response map[string]interface{}
	json.Unmarshal(rr.Body.Bytes(), &response)
	assert.Contains(t, response, "token") // Check if token is returned
	assert.NotEmpty(t, response["token"]) // Ensure the token is not empty
}
