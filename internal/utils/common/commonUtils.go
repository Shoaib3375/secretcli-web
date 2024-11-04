package common

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/mahinops/secretcli-web/internal/utils/crypto"
	"github.com/mahinops/secretcli-web/model"
)

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

func RespondWithSuccess(w http.ResponseWriter, code int, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(model.SuccessResponse{
		Code:    code,
		Message: message,
		Data:    data,
	})
}
