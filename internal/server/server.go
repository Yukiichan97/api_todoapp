package server

import (
	"backend_todoapp/internal/tasks"

	"github.com/labstack/echo/v4"
)

func NewServer() *echo.Echo {
	e := echo.New()
	repo := tasks.NewRepository()
	service := tasks.NewService(repo)
	handler := tasks.NewHandler(service)

	handler.RegisterRoutes(e)

	return e
}
