package main

import (
	"net/http"
	"strings"

	"ApiGo/views"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Counter = views.Counter

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	// Counter state
	count := Counter{Count: 0}

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
		count.Count++
		return c.Redirect(http.StatusSeeOther, "/")
	})

	// Start the server
	e.Logger.Fatal(e.Start(":1323"))
}
