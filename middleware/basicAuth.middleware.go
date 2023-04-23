package middleware

import (
	"github.com/fazaalexander/go-gorm/config"
	"github.com/fazaalexander/go-gorm/model"
	"github.com/fazaalexander/go-gorm/utils"

	"github.com/labstack/echo/v4/middleware"

	"github.com/labstack/echo/v4"
)

// Basic Auth
func BasicAuthMiddleware(e *echo.Echo) {
	e.Use(middleware.BasicAuth(func(username, password string, ctx echo.Context) (bool, error) {
		var user model.User
		id := ctx.Param("id")
		config.DB.First(&user, id)

		hashedPassword, err := utils.HashPassword(password)
		if err != nil {
			return false, nil
		}

		if username != user.Name {
			err = utils.ComparePassword(hashedPassword, user.Password)
			if err != nil {
				return false, err
			}
		}

		return true, nil
	}))
}
