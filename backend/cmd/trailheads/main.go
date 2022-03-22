package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/michaelpeterswa/trailheads/backend/internal/cache"
	"github.com/michaelpeterswa/trailheads/backend/internal/dao"
	"github.com/michaelpeterswa/trailheads/backend/internal/db"
	"github.com/michaelpeterswa/trailheads/backend/internal/handlers"
	"github.com/michaelpeterswa/trailheads/backend/internal/logging"
	"github.com/michaelpeterswa/trailheads/backend/internal/structs"
	"go.uber.org/zap"
)

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

	mongoClient, err := db.InitMongo(ctx, "mongodb://root:example@mongo:27017")
	if err != nil {
		logger.Error("unable to acquire mongo client", zap.Error(err))
	}

	// Users DAO and Handler
	usersDAO := dao.NewUsersDAO(mongoClient)
	usersHandler := handlers.NewUsersHandler(usersDAO, logger)
	// Trailheads DAO and Handler

	e := echo.New()
	e.Use(middleware.Static("dist"))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))
	p := prometheus.NewPrometheus("echo", nil)
	p.Use(e)

	e.GET("/healthcheck", func(c echo.Context) error {
		return c.JSON(http.StatusOK, structs.HealthCheck{
			Healthy: "ok",
		})
	})

	apiGroup := e.Group("/api")
	apiGroup.Use(middleware.BasicAuth(func(username string, apikey string, c echo.Context) (bool, error) {
		user, err := usersDAO.GetUser(ctx, username)
		if err != nil {
			fmt.Println(err)
			return false, err
		}

		if user != nil && user.APIKey == apikey {
			return true, nil
		}
		return false, nil
	}))

	apiGroup.GET("/user", usersHandler.GetUser)
	apiGroup.POST("/user", usersHandler.CreateUser)
	apiGroup.PUT("/user", usersHandler.UpdateUser)
	apiGroup.DELETE("/user", usersHandler.DeleteUser)

	e.Any("/*", func(c echo.Context) error {
		return c.File("dist/index.html")
	})

	err = e.Start(":8080")
	logger.Fatal("failed to start echo", zap.Error(err))
}
