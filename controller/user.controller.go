package controller

import (
	"log"
	"net/http"
	"time"

	"github.com/fazaalexander/go-gorm/config"
	"github.com/fazaalexander/go-gorm/model"
	"github.com/fazaalexander/go-gorm/service"
	"github.com/fazaalexander/go-gorm/utils"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type jwtCustomClaims struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

type Controller struct {
}

func Login(c echo.Context) error {
	username := c.FormValue("name")
	password := c.FormValue("password")

	var user model.User
	if err := config.DB.Model(&user).Where("name = ?", username).First(&user).Error; err != nil {
		log.Println("username tidak ditemukan")
		return echo.ErrUnauthorized
	}

	if err := utils.ComparePassword(user.Password, password); err != nil {
		log.Println("password salah")
		return echo.ErrUnauthorized
	}

	// Set custom claims
	claims := &jwtCustomClaims{
		user.Name,
		user.Email,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}

// get all users
func (m *Controller) GetUsersController(c echo.Context) error {
	var users []model.UserResponse

	users, err := service.GetUserRepository().GetUsers()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   users,
	})
}

// get user by id
func (m *Controller) GetUserController(c echo.Context) error {
	// your solution here
	var user *model.UserResponse
	id := c.Param("id")

	user, err := service.GetUserRepository().GetUserById(id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get user by id",
		"user":    user,
	})
}

// create new user
func (m *Controller) CreateUserController(c echo.Context) error {
	user := model.User{}
	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "failed",
		})
	}

	if err = service.GetUserRepository().CreateUser(&user); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success create new user",
		"user":    user,
	})
}

// delete user by id
func (m *Controller) DeleteUserController(c echo.Context) error {
	// your solution here
	// var user *model.User
	id := c.Param("id")

	err := service.GetUserRepository().DeleteUser(id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete user by id",
	})
}

// update user by id
func (m *Controller) UpdateUserController(c echo.Context) error {
	// your solution here
	var user model.User
	id := c.Param("id")
	name := c.FormValue("name")
	email := c.FormValue("email")
	password := c.FormValue("password")

	err := service.GetUserRepository().UpdateUser(&user, id, name, email, password)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	// if err := config.DB.Model(&user).Where("id = ?", id).Updates(model.User{Name: name, Email: email, Password: password}).Error; err != nil {
	// 	return c.JSON(http.StatusBadRequest, map[string]string{
	// 		"error": err.Error(),
	// 	})
	// }

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Success update user",
	})
}
