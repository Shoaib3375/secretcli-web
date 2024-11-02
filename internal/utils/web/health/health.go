// health.go

package health

import (
	"net/http"
)

// HealthCheck handles the health check request.
//
//	@Summary		Health Check
//	@Description	Check the health of the service.
//	@Tags			health
//	@Success		200	{string}	string	"OK"
//	@Router			/health [get]
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
