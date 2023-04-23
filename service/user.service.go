package service

import (
	"net/http"

	"github.com/fazaalexander/go-gorm/config"
	"github.com/fazaalexander/go-gorm/model"

	"github.com/labstack/echo/v4"
)

// Interface User
type IUserService interface {
	GetUsers() ([]model.UserResponse, error)
	GetUserById(id string) (*model.UserResponse, error)
	CreateUser(*model.User) error
	DeleteUser(id string) error
	UpdateUser(user *model.User, id string, name string, email string, password string) error
}

// Struct untuk mengimplementasikan interface
type UserRepository struct {
	Func IUserService
}

var userRepository IUserService

func init() {
	ur := &UserRepository{}
	ur.Func = ur

	userRepository = ur
}

func GetUserRepository() IUserService {
	return userRepository
}

func SetUserRepository(ur IUserService) {
	userRepository = ur
}

func (u *UserRepository) GetUsers() ([]model.UserResponse, error) {
	var users []model.UserResponse
	// if err := config.DB.Find(&users).Error; err != nil {
	// 	return nil, echo.NewHTTPError(http.StatusBadRequest, err.Error())
	// }

	if err := config.DB.Model(&model.User{}).Select("id, name, email, password").Find(&users).Error; err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return users, nil
}

func (u *UserRepository) GetUserById(id string) (*model.UserResponse, error) {
	var user model.UserResponse
	// if err := config.DB.First(&user, id).Error; err != nil {
	// 	return nil, echo.NewHTTPError(http.StatusBadRequest, err.Error())
	// }

	if err := config.DB.Model(&model.User{}).Select("id, name, email, password").First(&user, id).Error; err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return &user, nil
}

func (u *UserRepository) CreateUser(user *model.User) error {
	if err := config.DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return nil
}

func (u *UserRepository) DeleteUser(id string) error {
	var user model.User
	if err := config.DB.First(&user, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := config.DB.Delete(&user, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return nil
}

func (u *UserRepository) UpdateUser(user *model.User, id string, name string, email string, password string) error {
	if err := config.DB.Model(user).Where("id = ?", id).Updates(model.User{Name: name, Email: email, Password: password}).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return nil
}
