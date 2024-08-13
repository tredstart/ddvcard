package routes

import (
	"ddcard/internal/views"

	"github.com/labstack/echo/v4"
)

func IndexPage(c echo.Context) error {
	return views.IndexPage().Render(c.Request().Context(), c.Response())
}
