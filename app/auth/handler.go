package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mahinops/secretcli-web/model"
)

type AuthHandler struct {
	usecase model.AuthUsecase
}

func NewAuthHandler(usecase model.AuthUsecase) *AuthHandler {
	return &AuthHandler{usecase: usecase}
}

func (h *AuthHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user model.Auth
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	ctx := context.Background()

	name, err := h.usecase.Create(ctx, user)
	fmt.Println(err)
	if err != nil {
		http.Error(w, "Error registering user"+err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with a success message and the user's name
	w.WriteHeader(http.StatusCreated) // Set the status to 201 Created
	response := map[string]string{
		"message": "User " + name + " created successfully",
	}
	json.NewEncoder(w).Encode(response) // Encode the response as JSON

}
