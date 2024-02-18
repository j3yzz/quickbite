package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Health struct{}

type SuccessResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func (h Health) Handle(c echo.Context) error {
	return c.JSON(http.StatusOK, SuccessResponse{
		Success: true,
		Message: "everything is fine",
	})
}

func (h Health) Register(g *echo.Group) {
	g.GET("/health-check", h.Handle)
}
