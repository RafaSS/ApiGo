package http

import (
	"ApiGo/types"
	"ApiGo/views" // Added log package
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo, mangaService types.MangaService) {

	e.GET("/", func(c echo.Context) error {
		viewModel, err := mangaService.GetMangaList(1)
		if err != nil {
			return err
		}

		// Render the template with the view model
		template := views.Mangas(viewModel)
		var htmlBuilder strings.Builder
		if err := template.Render(c.Request().Context(), &htmlBuilder); err != nil {
			return err
		}

		return c.HTML(http.StatusOK, htmlBuilder.String())
	})

	e.POST("/search", func(c echo.Context) error {
		// Get the search term from the form data
		title := c.FormValue("title")

		// Fetch the manga list with search parameter
		viewModel, err := mangaService.GetMangaListWithTitle(title)
		if err != nil {
			return err
		}

		// Render the template with the view model
		template := views.Mangas(viewModel)
		var htmlBuilder strings.Builder

		if err := template.Render(c.Request().Context(), &htmlBuilder); err != nil {
			return err
		}

		return c.HTML(http.StatusOK, htmlBuilder.String())
	})

	e.GET("/list", func(c echo.Context) error {
		// Get the page number from query parameters, default to 1
		pageStr := c.QueryParam("page")
		page := 1
		if pageStr != "" {
			if p, err := strconv.Atoi(pageStr); err == nil {
				page = p
			}
		}

		// Fetch manga list from the service
		viewModel, err := mangaService.GetMangaList(page)
		if err != nil {
			return err
		}

		// Render the template with the view model
		template := views.Mangas(viewModel)
		var htmlBuilder strings.Builder
		if err := template.Render(c.Request().Context(), &htmlBuilder); err != nil {
			return err
		}

		return c.HTML(http.StatusOK, htmlBuilder.String())
	})
}
