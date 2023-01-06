package db

import (
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabaseConnection(logger *zap.Logger) *gorm.DB {
	logger.Info("Setting up database connection")

	// вообще по-хорошему конфигурационные данные вынести и сокрыть их здесь
	// кст лишний пробел, поставь trailing space или другие аналоги
	dsn := `host=localhost
			user=srkn
			password=1234
			dbname=go_example
			port=5432
			sslmode=disable
			TimeZone=Europe/Moscow`

	logger.Info("Opening database connection")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("DB connection failed")
	}

	return db
}
