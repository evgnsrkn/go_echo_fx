package handler

import (
	"fmt"
	"management/model"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserHandler struct {
	logger *zap.Logger
	db     *gorm.DB
}

func NewUserHandler(db *gorm.DB, logger *zap.Logger) *UserHandler {
	logger.Info("Setting up user handler")
	return &UserHandler{
		logger: logger,
		db:     db,
	}
}

func (uh UserHandler) GetUserById(c echo.Context) error {
	var user model.User

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		uh.logger.Debug("Invalid uuid format")
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = uh.db.First(user, id).Error
	if err != nil {
		uh.logger.Debug(fmt.Sprintf("User with id %s does not exist", id))
	}

	return c.JSON(http.StatusOK, user)
}

func (uh UserHandler) GetAllUsers(c echo.Context) error {
	users := []model.User{}

	uh.db.Find(&users)

	return c.JSON(http.StatusOK, users)
}
