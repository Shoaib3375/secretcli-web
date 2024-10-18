// health.go

package health

import (
	"net/http"
)

// Handler checks the health of the application
func Handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
