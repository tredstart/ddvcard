package routes

import (
	"ddcard/internal/models"
	"ddcard/internal/views"
	"log"
	"net/http"
	"strconv"

	"github.com/a-h/templ"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func PatientPage(c echo.Context) error {
	patient, err := models.FetchPatient(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Coś poszło nie tak: "+err.Error())
	}

	return views.PatientPage(patient, map[int]models.Tooth{}).Render(c.Request().Context(), c.Response())
}

func UpdateOrCreateTooth(c echo.Context) error {
	patient := c.Param("patient_id")
	tooth := c.Param("tooth_id")
	num, err := strconv.ParseInt(c.FormValue("num"), 10, 32)
	if err != nil {
		log.Println(c.Request())
		return c.HTML(http.StatusBadRequest, "<div><p>Zły numer zęba</p></div>")
	}
	var t models.Tooth
	if tooth == "" {
		c.Request().ParseForm()
		t.Id = uuid.NewString()
		t.Number = int(num)
		t.Patient = patient
	} else {
		if t, err = models.FetchTooth(tooth); err != nil {
			return c.HTML(http.StatusBadRequest, "<div><p>Coś poszło nie tak</p></div>")
		}
	}
	side, err := strconv.ParseInt(c.FormValue("side"), 10, 32)
	if err != nil {
		return c.HTML(http.StatusBadRequest, "<div><p>Źle zaznaczona strona</p></div>")
	}
	proc, err := strconv.ParseInt(c.FormValue("proc"), 10, 32)
	if err != nil {
		return c.HTML(http.StatusBadRequest, "<div><p>Źle zaznaczona procedura</p></div>")
	}
	tooth_type := c.FormValue("type")
	var view func(string, models.Tooth, int) templ.Component
	switch side {
	case int64(models.Top):
		t.T = int(proc)
		if tooth_type == "back" {
			view = views.BackTop
		} else {
			view = views.FrontTop
		}
	case int64(models.Right):
		t.R = int(proc)
		if tooth_type == "back" {
			view = views.BackRight
		} else {
			view = views.FrontRight
		}
	case int64(models.Bottom):
		t.B = int(proc)
		if tooth_type == "back" {
			view = views.BackBottom
		} else {
			view = views.FrontBottom
		}
	case int64(models.Left):
		t.L = int(proc)
		if tooth_type == "back" {
			view = views.BackLeft
		} else {
			view = views.FrontLeft
		}
	case int64(models.Center):
		t.Center = int(proc)
		if tooth_type == "back" {
			view = views.BackCenter
		}
	}
	if tooth == "" {
		if err = models.CreateTooth(t); err != nil {
			return c.HTML(http.StatusBadRequest, "<div><p>Coś poszło nie tak</p></div>")
		}
	} else {
		if err = models.UpdateTooth(t); err != nil {
			return c.HTML(http.StatusBadRequest, "<div><p>Coś poszło nie tak</p></div>")
		}
	}

	return view(patient, t, int(num)).Render(c.Request().Context(), c.Response())
}
