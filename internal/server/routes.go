package server

import (
	"ddcard/internal/server/routes"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Static("/static", "static")
	e.GET("/", routes.IndexPage)
	e.GET("/new-patient", routes.NewPatientForm)
	e.POST("/new-patient", routes.NewPatient)
    e.POST("/search", routes.Search)

	return e
}
