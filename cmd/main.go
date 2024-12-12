package main

import (
	"net/http"
	"strings"

	"ApiGo/views"

	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
	"time"

	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Counter = views.Counter

func authenticate(username, password, clientID, clientSecret string) (string, string, error) {
	data := url.Values{}
	data.Set("grant_type", "password")
	data.Set("username", username)
	data.Set("password", password)
	data.Set("client_id", clientID)
	data.Set("client_secret", clientSecret)

	req, err := http.NewRequest("POST", "https://auth.mangadex.org/realms/mangadex/protocol/openid-connect/token", bytes.NewBufferString(data.Encode()))
	if err != nil {
		return "", "", err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", "", fmt.Errorf("failed to authenticate: %s", resp.Status)
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", "", err
	}

	accessToken, ok := result["access_token"].(string)
	if !ok {
		return "", "", fmt.Errorf("failed to get access token")
	}

	refreshToken, ok := result["refresh_token"].(string)
	if !ok {
		return "", "", fmt.Errorf("failed to get refresh token")
	}

	return accessToken, refreshToken, nil
}

func refreshToken(refreshToken, clientID, clientSecret string) (string, string, error) {
	data := url.Values{}
	data.Set("grant_type", "refresh_token")
	data.Set("refresh_token", refreshToken)
	data.Set("client_id", clientID)
	data.Set("client_secret", clientSecret)

	req, err := http.NewRequest("POST", "https://auth.mangadex.org/realms/mangadex/protocol/openid-connect/token", bytes.NewBufferString(data.Encode()))
	if err != nil {
		return "", "", err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", "", fmt.Errorf("failed to refresh token: %s", resp.Status)
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", "", err
	}

	newAccessToken, ok := result["access_token"].(string)
	if !ok {
		return "", "", fmt.Errorf("failed to get new access token")
	}

	newRefreshToken, ok := result["refresh_token"].(string)
	if !ok {
		return "", "", fmt.Errorf("failed to get new refresh token")
	}

	return newAccessToken, newRefreshToken, nil
}

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

	_, refreshTokenValue, err := authenticate(username, password, clientID, clientSecret)
	if err != nil {
		fmt.Printf("Error authenticating: %v\n", err)
		return
	}
	ticker := time.NewTicker(10 * time.Minute)
	defer ticker.Stop()

	e := echo.New()
	e.Use(middleware.Logger())

	// Counter state
	count := Counter{Count: "pass"}

	// Route for rendering the template
	e.GET("/", func(c echo.Context) error {
		template := views.Counts(&count)
		var htmlBuilder strings.Builder
		err := template.Render(c.Request().Context(), &htmlBuilder)
		if err != nil {
			return err
		}
		return c.HTML(http.StatusOK, htmlBuilder.String())
	})

	e.POST("/count", func(c echo.Context) error {
		count.Count = refreshTokenValue
		template := views.Conter(&count)
		var htmlBuilder strings.Builder
		err := template.Render(c.Request().Context(), &htmlBuilder)
		if err != nil {
			return err
		}
		return c.HTML(http.StatusOK, htmlBuilder.String())
	})

	// Start the server
	go func() {
		e.Logger.Fatal(e.Start(":1323"))
	}()

	for range ticker.C {
		_, newRefreshToken, err := refreshToken(refreshTokenValue, clientID, clientSecret)
		if err != nil {
			fmt.Printf("Error refreshing token: %v\n", err)
			return
		}
		refreshTokenValue = newRefreshToken
		fmt.Println("Token refreshed successfully")
	}
}
