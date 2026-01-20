package main

import (
	"net/http"

	"github.com/itsjhonale/budget-api/internal/router"
	"github.com/itsjhonale/budget-api/internal/utils"
)

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
