package http

import (
	"ApiGo/types"
	"ApiGo/views"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/labstack/echo/v4"
)

type Manga = views.Manga

func SetupRoutes(e *echo.Echo) {
	manga := types.Manga{}

	// Route for rendering the template
	e.GET("/", func(c echo.Context) error {
		mangaObj := views.Manga{
			Name: "fsaf",
		}
		template := views.Mangas(&mangaObj)
		var htmlBuilder strings.Builder
		err := template.Render(c.Request().Context(), &htmlBuilder)
		if err != nil {
			return err
		}
		return c.HTML(http.StatusOK, htmlBuilder.String())
	})

	e.POST("/search", func(c echo.Context) error {
		data := url.Values{}
		data.Set("name", "Jujutsu")

		req, err := http.NewRequest("GET", "https://api.mangadex.org/manga", bytes.NewBufferString(data.Encode()))
		if err != nil {
			return err
		}

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		// bodyBytes, err := io.ReadAll(resp.Body)
		// if err != nil {
		// 	return err
		// }
		// fmt.Println("🙏", string(bodyBytes))
		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("failed to refresh token: %s", resp.Status)
		}

		var result types.Manga
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			return err
		}

		// Extract the title from the response
		if len(result.Data) > 0 {
			manga = result
		}

		manga = result
		fmt.Println("🔍", fmt.Sprintf("%v", manga.Data[0].ID))

		mangaObj := views.Manga{
			Name: manga.Data[0].ID,
		}

		template := views.Title(&mangaObj)
		var htmlBuilder strings.Builder

		err = template.Render(c.Request().Context(), &htmlBuilder)
		if err != nil {
			return err
		}
		return c.HTML(http.StatusOK, htmlBuilder.String())
	})
}
