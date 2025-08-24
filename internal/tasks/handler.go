package tasks

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	service *Service
}

func NewHandler(s *Service) *Handler {
	return &Handler{service: s}
}

func (h *Handler) RegisterRoutes(e *echo.Echo) {
	e.POST("/tasks", h.createTask)
	e.GET("/tasks", h.getTasks)
	e.PUT("/tasks/:id", h.updateTask)
	e.DELETE("/tasks/:id", h.deleteTask)
}

func (h *Handler) createTask(c echo.Context) error {
	t := new(Task)
	if err := c.Bind(t); err != nil {
		return err
	}
	created := h.service.CreateTask(*t)
	return c.JSON(http.StatusCreated, created)
}

func (h *Handler) getTasks(c echo.Context) error {
	list := h.service.GetTasks()
	return c.JSON(http.StatusOK, list)
}

func (h *Handler) updateTask(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	t := new(Task)
	if err := c.Bind(t); err != nil {
		return err
	}
	updated, ok := h.service.UpdateTask(id, *t)
	if !ok {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Task not found"})
	}
	return c.JSON(http.StatusOK, updated)
}

func (h *Handler) deleteTask(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	ok := h.service.DeleteTask(id)
	if !ok {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Task not found"})
	}
	return c.NoContent(http.StatusNoContent)
}
