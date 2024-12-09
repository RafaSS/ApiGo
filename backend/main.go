package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func Log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		begin := time.Now()
		next.ServeHTTP(w, r)
		fmt.Println(r.URL.String(), r.Method, time.Since(begin))
	})
}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc(
		"/api/users/{foo}",
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "hello, world")
		},
	)
	srv := &http.Server{
		Addr:                         ":8080",
		Handler:                      Log(mux),
		DisableGeneralOptionsHandler: false,
		ReadTimeout:                  10 * time.Second,
		WriteTimeout:                 10 * time.Second,
		IdleTimeout:                  1 * time.Minute,
	}
	if err := srv.ListenAndServe(); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			panic(err)
		}
	}
}
func autenticate() {
	data := url.Values{}
	data.Set("grant_type", "password")
	data.Set("username", "RafaelSS")
	data.Set("password", "fila4545")
	data.Set("client_id", "personal-client-642bd0b4-7288-4fb7-ad53-45da5f9b78eb-b18c2790")
	data.Set("client_secret", "i4papsLZE7225MCKmRUkZynhSLu7gM9k")
	resp, err := http.Post("https://auth.mangadex.org/realms/mangadex/protocol/openid-connect/token", "application/x-www-form-urlencoded", strings.NewReader(data.Encode()))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: status code", resp.StatusCode)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	fmt.Println(string(body))
}

func init() {
	go autenticate()
}
