package routes

import (
	"ddcard/internal/models"
	"ddcard/internal/views"
	"net/http"

	"github.com/labstack/echo/v4"
)

func IndexPage(c echo.Context) error {
	return views.IndexPage().Render(c.Request().Context(), c.Response())
}

func Search(c echo.Context) error {
	pesel := c.FormValue("pesel")
	name_surname := c.FormValue("name")
	if pesel != "" {
		patients, err := models.FetchPatientsByPesel(pesel)
		if err != nil {
			return c.HTML(http.StatusNotFound, "<div><p>Nie znaleziono żadnych pacjentów</p></div>")
		}
		return views.Results(patients).Render(c.Request().Context(), c.Response())
	}
	if name_surname != "" {
	}
	return c.HTML(http.StatusNotFound, "<div><p>Nie znaleziono żadnych pacjentów</p></div>")
}
