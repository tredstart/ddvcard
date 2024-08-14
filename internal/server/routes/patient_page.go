package routes

import (
	"ddcard/internal/models"
	"ddcard/internal/views"
	"net/http"

	"github.com/labstack/echo/v4"
)

func PatientPage(c echo.Context) error {
	patient, err := models.FetchPatient(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Coś poszło nie tak: "+err.Error())
	}

	return views.PatientPage(patient).Render(c.Request().Context(), c.Response())
}
