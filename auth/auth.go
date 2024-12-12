package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

func Authenticate(username, password, clientID, clientSecret string) (string, string, error) {
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

func RefreshToken(refreshToken, clientID, clientSecret string) (string, string, error) {
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
