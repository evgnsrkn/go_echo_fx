package db

import (
	"management/model"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabaseConnection(logger *zap.Logger) *gorm.DB {
	logger.Info("Setting up database connection")

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

	logger.Info("Migrating User model")
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		panic("Failed migration")
	}

	return db
}
