package routes

import (
	"ddcard/internal/models"
	"ddcard/internal/views"
	"net/http"
	"strings"

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
		var name, surname string
		split := strings.Split(name_surname, " ")
		if len(split) > 1 {
			surname = split[len(split)-1]
			name = strings.Join(split[:len(split)-2], " ")
		} else {
			name = split[0]
			surname = split[0]
		}
		patients, err := models.FetchPatientsByName(name, surname)
		if err != nil {
			return c.HTML(http.StatusNotFound, "<div><p>Nie znaleziono żadnych pacjentów</p></div>")
		}
		return views.Results(patients).Render(c.Request().Context(), c.Response())
	}
	return c.HTML(http.StatusNotFound, "<div><p>Nie znaleziono żadnych pacjentów</p></div>")
}
