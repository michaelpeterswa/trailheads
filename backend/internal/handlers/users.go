package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/michaelpeterswa/trailheads/backend/internal/dao"
	"github.com/michaelpeterswa/trailheads/backend/internal/structs"
)

type UsersHandler struct {
	usersDAO *dao.UsersDAO
}

func NewUsersHandler(d *dao.UsersDAO) *UsersHandler {
	return &UsersHandler{usersDAO: d}
}

func (h *UsersHandler) GetUser(c echo.Context) error {
	foundUser, err := h.usersDAO.GetUser(c.Request().Context(), "example@test.com")
	if err != nil {
		return c.JSON(http.StatusOK, structs.Success{
			Success: false,
		})
	}
	return c.JSON(http.StatusOK, foundUser)
}
