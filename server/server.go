package server

import (
	"context"
	"management/handler"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func NewServer(handler *handler.UserHandler, db *gorm.DB, lc fx.Lifecycle, logger *zap.Logger) *http.Server {
	logger.Info("Setting up server")
	e := echo.New()
	e.GET("/users/:id", handler.GetUserById)
	e.GET("/users", handler.GetAllUsers)

	// Зачем, если переменная `e` (*echo.Echo) поддерживает запуск `e.Start(:8000)` ??
	// гл. Quick Start - https://echo.labstack.com/guide/#quick-start
	srv := &http.Server{Addr: ":5555", Handler: e}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			// Более того, раз ты инжектишь логгер, то его надо использовать
			// и для логирования запросов сервака. На сайте эха всё есть
			// гл. Quick Start и Middleware > Logger
			logger.Info("Server started")
			go srv.ListenAndServe()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})
	return srv
}
