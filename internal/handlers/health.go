package handlers

import (
	"net/http"

	"github.com/itsjhonalex/budget-api/internal/utils"
)

// HealthHandler godoc
// @Summary Health check
// @Description Verifica el estado del servicio
// @Tags health
// @Produce json
// @Success 200 {object} utils.Response
// @Router /health [get]
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	utils.RespondSuccess(w, map[string]string{
		"status":  "ok",
		"service": "budget-api",
	})
}
