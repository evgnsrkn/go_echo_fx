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

// Для методов структуры используется указатель на структуру
func (uh UserHandler) GetUserById(c echo.Context) error {
	var user model.User

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		// не так важно, но как в проде будешь узнавать состояния приложения?
		// надо в файлик (тоже) логировать
		uh.logger.Debug("Invalid uuid format")
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// 1) необходимо указать ссылку на `&user`;
	// 2) также некритично, здесь можно сокращение использовать:
	// if errFind := db.Find(...); err != nil {
	//	...
	// }
	err = uh.db.First(user, id).Error
	if err != nil {
		// тут мелочь, return'а нету
		uh.logger.Debug(fmt.Sprintf("User with id %s does not exist", id))
	}

	return c.JSON(http.StatusOK, user)
}

func (uh UserHandler) GetAllUsers(c echo.Context) error {
	// Здесь лучше `var users []User`, поскольку 1) ты именно так выше указал
	// 2) запись является общепринятой. Можешь по популярным репам глянуть
	users := []model.User{}

	uh.db.Find(&users)

	return c.JSON(http.StatusOK, users)
}
