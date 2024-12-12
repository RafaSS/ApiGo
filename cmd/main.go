package main

import (
	"ApiGo/auth"
	"ApiGo/http"
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Error loading .env file: %v\n", err)
		return
	}

	username := os.Getenv("USERNAMEMANGA")
	password := os.Getenv("PASSWORD")
	clientID := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")

	_, refreshTokenValue, err := auth.Authenticate(username, password, clientID, clientSecret)
	if err != nil {
		fmt.Printf("Error authenticating: %v\n", err)
		return
	}
	ticker := time.NewTicker(10 * time.Minute)
	defer ticker.Stop()

	e := echo.New()
	e.Use(middleware.Logger())

	http.SetupRoutes(e)

	// Start the server
	go func() {
		e.Logger.Fatal(e.Start(":1323"))
	}()

	for range ticker.C {
		_, newRefreshToken, err := auth.RefreshToken(refreshTokenValue, clientID, clientSecret)
		if err != nil {
			fmt.Printf("Error refreshing token: %v\n", err)
			return
		}
		refreshTokenValue = newRefreshToken
		fmt.Println("Token refreshed successfully")
	}
}
