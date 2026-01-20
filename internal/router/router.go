package router

import (
	"net/http"

	"github.com/itsjhonale/budget-api/internal/handlers"
)

// SetupRoutes configura todas las rutas de la API
func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	// Health check
	mux.HandleFunc("/health", handlers.HealthHandler)

	return mux
}
