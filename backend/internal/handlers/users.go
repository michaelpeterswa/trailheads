package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/michaelpeterswa/trailheads/backend/internal/dao"
	"github.com/michaelpeterswa/trailheads/backend/internal/structs"
	"go.uber.org/zap"
)

type UsersHandler struct {
	usersDAO *dao.UsersDAO
	logger   *zap.Logger
}

func NewUsersHandler(d *dao.UsersDAO, z *zap.Logger) *UsersHandler {
	return &UsersHandler{usersDAO: d, logger: z}
}

func (h *UsersHandler) CreateUser(c echo.Context) error {
	var userToCreate structs.User
	body := c.Request().Body
	defer body.Close()

	err := json.NewDecoder(body).Decode(&userToCreate)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, structs.Success{
			Success: false,
		})
	}

	err = h.usersDAO.CreateUser(c.Request().Context(), &userToCreate)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, structs.Success{
			Success: false,
		})
	}
	return c.JSON(http.StatusCreated, structs.Success{
		Success: true,
	})
}

func (h *UsersHandler) GetUser(c echo.Context) error {
	user := c.QueryParam("username")

	foundUser, err := h.usersDAO.GetUser(c.Request().Context(), user)
	if err != nil {
		return c.JSON(http.StatusOK, structs.Success{
			Success: false,
		})
	}
	return c.JSON(http.StatusOK, foundUser)
}

func (h *UsersHandler) UpdateUser(c echo.Context) error {
	var userToUpdate structs.User
	body := c.Request().Body
	defer body.Close()

	err := json.NewDecoder(body).Decode(&userToUpdate)
	if err != nil {
		h.logger.Error("error decoding updateUser body", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, structs.Success{
			Success: false,
		})
	}

	err = h.usersDAO.UpdateUser(c.Request().Context(), &userToUpdate)
	if err != nil {
		h.logger.Error("error updating user", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, structs.Success{
			Success: false,
		})
	}
	return c.JSON(http.StatusCreated, structs.Success{
		Success: true,
	})
}

func (h *UsersHandler) DeleteUser(c echo.Context) error {
	user := c.QueryParam("username")

	err := h.usersDAO.DeleteUser(c.Request().Context(), user)
	if err != nil {
		return c.JSON(http.StatusOK, structs.Success{
			Success: false,
		})
	}
	return c.JSON(http.StatusOK, structs.Success{
		Success: true,
	})
}
