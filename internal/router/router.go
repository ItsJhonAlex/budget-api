package router

import (
	"net/http"

	"github.com/itsjhonalex/budget-api/internal/handlers"
	httpSwagger "github.com/swaggo/http-swagger"
)

// SetupRoutes configura todas las rutas de la API
func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	// Health check
	mux.HandleFunc("/health", handlers.HealthHandler)

	// Swagger UI
	mux.HandleFunc("/swagger/", httpSwagger.WrapHandler)

	return mux
}
