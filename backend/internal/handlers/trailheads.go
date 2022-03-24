package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/michaelpeterswa/trailheads/backend/internal/dao"
	"github.com/michaelpeterswa/trailheads/backend/internal/structs"
	"github.com/michaelpeterswa/trailheads/backend/internal/trailheads"
	"go.uber.org/zap"
)

type TrailheadsHandler struct {
	trailheadsDAO *dao.TrailheadsDAO
	logger        *zap.Logger
}

func NewTrailheadsHandler(d *dao.TrailheadsDAO, z *zap.Logger) *TrailheadsHandler {
	return &TrailheadsHandler{trailheadsDAO: d, logger: z}
}

func (t *TrailheadsHandler) CreateTrailhead(c echo.Context) error {
	var trailheadToCreate trailheads.Trailhead
	body := c.Request().Body
	defer body.Close()

	err := json.NewDecoder(body).Decode(&trailheadToCreate)
	if err != nil {
		t.logger.Error("error decoding trailhead body", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, structs.Success{
			Success: false,
		})
	}

	err = t.trailheadsDAO.CreateTrailhead(c.Request().Context(), &trailheadToCreate)
	if err != nil {
		t.logger.Error("error creating trailhead", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, structs.Success{
			Success: false,
		})
	}
	return c.JSON(http.StatusCreated, structs.Success{
		Success: true,
	})
}

func (t *TrailheadsHandler) GetTrailhead(c echo.Context) error {
	trailheadName := c.QueryParam("name")

	foundTrailhead, err := t.trailheadsDAO.GetTrailhead(c.Request().Context(), trailheadName)
	if err != nil {
		t.logger.Error("error getting trailhead", zap.Error(err))
		return c.JSON(http.StatusOK, structs.Success{
			Success: false,
		})
	}
	return c.JSON(http.StatusOK, foundTrailhead)
}

func (t *TrailheadsHandler) GetTrailheads(c echo.Context) error {
	foundTrailheads, err := t.trailheadsDAO.GetTrailheads(c.Request().Context())
	if err != nil {
		t.logger.Error("error getting trailheads", zap.Error(err))
		return c.JSON(http.StatusOK, structs.Success{
			Success: false,
		})
	}
	return c.JSON(http.StatusOK, foundTrailheads)
}

func (t *TrailheadsHandler) UpdateTrailhead(c echo.Context) error {
	var trailheadToUpdate trailheads.Trailhead
	body := c.Request().Body
	defer body.Close()

	err := json.NewDecoder(body).Decode(&trailheadToUpdate)
	if err != nil {
		t.logger.Error("error decoding updateTrailhead body", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, structs.Success{
			Success: false,
		})
	}

	err = t.trailheadsDAO.UpdateTrailhead(c.Request().Context(), &trailheadToUpdate)
	if err != nil {
		t.logger.Error("error updating trailhead", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, structs.Success{
			Success: false,
		})
	}
	return c.JSON(http.StatusCreated, structs.Success{
		Success: true,
	})
}

func (t *TrailheadsHandler) DeleteTrailhead(c echo.Context) error {
	trailheadName := c.QueryParam("name")

	err := t.trailheadsDAO.DeleteTrailhead(c.Request().Context(), trailheadName)
	if err != nil {
		t.logger.Error("error deleting trailhead", zap.Error(err))
		return c.JSON(http.StatusOK, structs.Success{
			Success: false,
		})
	}
	return c.JSON(http.StatusOK, structs.Success{
		Success: true,
	})
}
