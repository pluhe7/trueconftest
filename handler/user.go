package handler

import (
	"net/http"
	"strconv"

	"github.com/pluhe7/trueconftest/model"

	"github.com/labstack/echo"
)

func (h *Handler) addUser(c echo.Context) error {
	u := &model.User{
		ID: h.Seq,
	}
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusInternalServerError, "Can't create user")
	}
	h.Users[u.ID] = u
	h.Seq++
	if err := h.jsonUpdate(); err != nil {
		return c.JSON(http.StatusInternalServerError, "Can't create user")
	}
	return c.JSON(http.StatusCreated, u)
}

func (h *Handler) getUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, h.Users)
}

func (h *Handler) getUserByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if h.Users[id] == nil {
		return c.JSON(http.StatusNotFound, "User not found")
	}
	return c.JSON(http.StatusOK, h.Users[id])
}

func (h *Handler) updateUserByID(c echo.Context) error {
	u := new(model.User)
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusInternalServerError, "Can't update user")
	}
	id, _ := strconv.Atoi(c.Param("id"))
	if h.Users[id] == nil {
		return c.JSON(http.StatusNotFound, "User not found")
	}
	h.Users[id].Name = u.Name
	if err := h.jsonUpdate(); err != nil {
		return c.JSON(http.StatusInternalServerError, "Can't update user")
	}
	return c.JSON(http.StatusOK, h.Users[id])
}

func (h *Handler) deleteUserByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if h.Users[id] == nil {
		return c.JSON(http.StatusNotFound, "User not found")
	}
	delete(h.Users, id)
	if err := h.jsonUpdate(); err != nil {
		return c.JSON(http.StatusInternalServerError, "Can't delete user")
	}
	return c.JSON(http.StatusOK, "User deleted")
}
