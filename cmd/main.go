package main

import (
	"management/pkg/db"
	"management/pkg/handler"
	"management/pkg/server"

	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func main() {
	fx.New(
		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log}
		}),
		fx.Provide(
			db.NewDatabaseConnection,
			zap.NewProduction,
			handler.NewUserHandler,
		),
		fx.Invoke(
			server.NewServer,
		),
	).Run()
}

// юзай линтер, он прям желательно должен быть во всех проектах

// вынеси функции запуска проекта и другие штуки в Makefile (что супер неочевидно выноси в `./scripts/*`)
// и опиши запуск в `README.md` (ведение уже по себе является хорошей практикой), хотя бы так:
// `make build` - сборка проекта
// `make daemon` - запуск демона

// опирайся лучше на стандартный project layout, можешь cmd, internal, pkg, config
// и др директории спокойно использовать

// тесты нужны
