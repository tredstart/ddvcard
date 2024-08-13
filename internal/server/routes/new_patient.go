package routes

import (
	"ddcard/internal/models"
	"ddcard/internal/views"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func NewPatientForm(c echo.Context) error {
	return views.PatientForm(models.Patient{}).Render(c.Request().Context(), c.Response())
}

func GetIllness(illness []string) int {
	var ill int

	for _, i := range illness {
		ill |= models.Illness[i]
	}

	return ill
}

func NewPatient(c echo.Context) error {
	c.Request().ParseForm()
	r := c.Request()
	pesel := r.FormValue("pesel")
	if patients, err := models.FetchPatientsByPesel(pesel); err == nil || len(patients) > 0 {
		c.Response().Header().Add("Hx-Retarget", "#errors")
		c.Response().Header().Add("Hx-Reswap", "innerHTML")
		return c.HTML(http.StatusBadRequest, `<div><p>Pacjent z tym numerem pesel już istnieje</p></div>`)
	}
	patient := models.Patient{}
	patient.Name = r.FormValue("name")
	patient.Surname = r.FormValue("surname")
	patient.Address = r.FormValue("address")
	if r.FormValue("child") == "on" {
		patient.Child = 1
	}
	patient.Birthdate = r.FormValue("birthdate")
	patient.Pesel = pesel
	patient.Registered = time.Now().Format("02.01.2006")
	patient.Phone = r.FormValue("phone")
	patient.Id = uuid.NewString()
	patient.Medicines = r.FormValue("medicines")
	patient.Illness = GetIllness(r.Form["illness"])
	if err := models.PatientCreate(patient); err != nil {
		c.Response().Header().Add("Hx-Retarget", "#errors")
		c.Response().Header().Add("Hx-Reswap", "innerHTML")
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.String(http.StatusOK, "Pacjent został zapisany poprawnie.")
}
