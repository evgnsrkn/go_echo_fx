package main

import (
	"management/db"
	"management/handler"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

func main() {
	fx.New(
		fx.Provide(
			db.NewDatabaseConnection,
			zap.NewProduction,
			handler.NewUserHandler,
			NewServer,
		),
		fx.Invoke(
			NewServer,
		),
	).Run()
}
