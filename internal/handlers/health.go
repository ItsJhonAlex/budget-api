package handlers

import (
	"net/http"

	"github.com/itsjhonale/budget-api/internal/utils"
)

// HealthHandler maneja el endpoint de health check
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	utils.RespondSuccess(w, map[string]string{
		"status": "ok",
		"service": "budget-api",
	})
}
