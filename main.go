package main

import (
	"ApiGo/http"
	"ApiGo/services"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// Initialize the MangaService
	mangaService := &services.MangaServiceImpl{}

	// Setup routes with the MangaService
	http.SetupRoutes(e, mangaService)

	// Start the server
	e.Start(":8080")
}
