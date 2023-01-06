package server

import (
	"context"
	"management/pkg/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func NewServer(handler *handler.UserHandler, db *gorm.DB, lc fx.Lifecycle, logger *zap.Logger) {
	logger.Info("Setting up server")
	e := echo.New()
	e.GET("/users/:id", handler.GetUserById)
	e.GET("/users", handler.GetAllUsers)

	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			logger.Info("request",
				zap.String("URI", v.URI),
				zap.Int("status", v.Status),
			)

			return nil
		},
	}))

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			logger.Info("Server started")
			go e.Start(":5555")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return e.Shutdown(ctx)
		},
	})
}
