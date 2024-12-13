package http

import (
	"ApiGo/types"
	"ApiGo/views"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {

	e.GET("/", func(c echo.Context) error {
		// Initialize MangaViewModel with empty Manga and AuthorName
		mangaObj := &types.MangaViewModel{
			Manga:      &types.Manga{},
			AuthorName: "",
		}
		template := views.Mangas(mangaObj)
		var htmlBuilder strings.Builder
		err := template.Render(c.Request().Context(), &htmlBuilder)
		if err != nil {
			return err
		}
		return c.HTML(http.StatusOK, htmlBuilder.String())
	})

	e.POST("/search", func(c echo.Context) error {
		// Get the search term from the form data
		title := c.FormValue("title")

		// Use query parameters in the URL
		queryParams := url.Values{}
		queryParams.Set("title", title)

		apiURL := "https://api.mangadex.org/manga?" + queryParams.Encode()

		req, err := http.NewRequest("GET", apiURL, nil)
		if err != nil {
			return err
		}

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("failed to get manga: %s", resp.Status)
		}

		// Decode the response into a struct that matches the API response
		var result struct {
			Data []types.Manga `json:"data"`
		}
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			return err
		}

		if len(result.Data) > 0 {
			manga := &result.Data[0]
			fmt.Println("🔍", result)

			// Create MangaViewModel with manga data and author name
			mangaObj := &types.MangaViewModel{
				Manga:      manga,
				AuthorName: getAuthorName(manga.Relationships),
			}

			template := views.Title(mangaObj)
			var htmlBuilder strings.Builder

			err = template.Render(c.Request().Context(), &htmlBuilder)
			if err != nil {
				return err
			}
			return c.HTML(http.StatusOK, htmlBuilder.String())
		}

		return c.String(http.StatusNotFound, "No manga found")
	})
}

// Helper function to get the author's name
func getAuthorName(relationships []types.Relationship) string {
	for _, rel := range relationships {
		if rel.Type == "author" {
			fmt.Println("👩‍🎨", rel)
			return rel.ID
		}
	}
	return "Unknown Author"
}
