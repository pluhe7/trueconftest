package handler

import (
	"github.com/labstack/echo"
)

func (h *Handler) Register(e *echo.Echo) {
	h.jsonInit()
	e.POST("/users", h.addUser)
	e.GET("/users", h.getUsers)
	e.GET("/users/:id", h.getUserByID)
	e.PUT("/users/:id", h.updateUserByID)
	e.DELETE("/users/:id", h.deleteUserByID)
}
