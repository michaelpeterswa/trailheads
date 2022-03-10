package main

import (
	"context"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/michaelpeterswa/wildfire/backend/internal/cache"
	"github.com/michaelpeterswa/wildfire/backend/internal/db"
	"github.com/michaelpeterswa/wildfire/backend/internal/logging"
	"go.uber.org/zap"
)

type HealthCheck struct {
	Healthy string `json:"healthy"`
}

func main() {
	ctx := context.Background()

	logger, err := logging.InitZapLogger()
	if err != nil {
		log.Fatal("unable to acquire zap logger")
	}

	_, err = cache.InitRedis(ctx, "redis", 6379)
	if err != nil {
		logger.Error("unable to acquire redis client", zap.Error(err))
	}

	_, err = db.InitMongo(ctx, "mongodb://localhost:27017")
	if err != nil {
		logger.Error("unable to acquire mongo client", zap.Error(err))
	}

	e := echo.New()
	e.Use(middleware.Static("dist"))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	e.GET("/healthcheck", func(c echo.Context) error {
		return c.JSON(http.StatusOK, HealthCheck{
			Healthy: "ok`",
		})
	})

	e.Any("/*", func(c echo.Context) error {
		return c.File("dist/index.html")
	})

	err = e.Start(":8080")
	logger.Fatal("failed to start echo", zap.Error(err))
}
