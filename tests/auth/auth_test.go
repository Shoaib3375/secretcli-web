package auth_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mahinops/secretcli-web/cmd"
	"github.com/stretchr/testify/assert"
)

// TestUserRegistration tests the user registration endpoint
func TestUserRegistration(t *testing.T) {
	// Initialize the application
	app, err := cmd.NewApp(".env.test")
	assert.NoError(t, err)
	defer app.CloseDatabase()

	// Prepare the registration request with the correct JSON key
	user := map[string]string{
		"name":     "testuser", // Correcting to match Auth struct
		"email":    "testuser@example.com",
		"password": "password123",
	}
	jsonUser, _ := json.Marshal(user)

	req, err := http.NewRequest(http.MethodPost, "/auth/register", bytes.NewBuffer(jsonUser))
	assert.NoError(t, err)

	// Create a ResponseRecorder to capture the response
	rr := httptest.NewRecorder()

	// Serve the request using the Router
	app.Router.ServeHTTP(rr, req)

	// Check the response code and body
	assert.Equal(t, http.StatusCreated, rr.Code)

	var response map[string]interface{}
	json.Unmarshal(rr.Body.Bytes(), &response)
	assert.Contains(t, response, "message")

	// Adjust based on your actual response logic
	assert.Equal(t, response["message"], "User testuser created successfully")
}
