package server

import (
	"backend_todoapp/internal/tasks"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewServer() *echo.Echo {
	e := echo.New()

	e.Use(middleware.CORS())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
	}))

	repo := tasks.NewRepository()
	service := tasks.NewService(repo)
	handler := tasks.NewHandler(service)

	handler.RegisterRoutes(e)

	return e
}
