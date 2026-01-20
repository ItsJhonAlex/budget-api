package main

import (
	"net/http"

	"github.com/itsjhonale/budget-api/internal/router"
	"github.com/itsjhonale/budget-api/internal/utils"

	_ "github.com/itsjhonale/budget-api/docs"
)

// @title Budget API
// @version 1.0
// @description API para gestión de presupuestos y transacciones
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email support@budget-api.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /
func main() {
	// Inicializar logger
	utils.InitLogger()
	utils.InfoLogger.Println("Initializing Budget API...")

	// Configurar rutas
	mux := router.SetupRoutes()

	// Iniciar servidor
	port := "8080"
	utils.InfoLogger.Printf("Server starting on port %s", port)

	if err := http.ListenAndServe(":"+port, mux); err != nil {
		utils.ErrorLogger.Fatal("Server failed to start:", err)
	}
}
