package common

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/mahinops/secretcli-web/internal/utils/crypto"
	"github.com/mahinops/secretcli-web/model"
)

// ParseAndValidatePayload reads, validates, and unmarshals JSON payload into the provided struct.
func ParseAndValidatePayload(r *http.Request, target interface{}) error {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := crypto.ValidatePayload(data, target); err != nil {
		return err
	}

	return json.Unmarshal(data, target)
}

// RespondWithSuccess sends a JSON response for successful operations like login or registration.
func RespondWithSuccess(w http.ResponseWriter, code int, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(model.SuccessResponse{
		Code:    code,
		Message: message,
		Data:    data,
	})
}
