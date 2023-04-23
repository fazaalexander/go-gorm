package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"

	"github.com/fazaalexander/go-gorm/model"
	"github.com/fazaalexander/go-gorm/service"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetUsersController(t *testing.T) {
	userRepository := &service.UserRepositoryMock{Mock: mock.Mock{}}
	service.SetUserRepository(userRepository)

	type args struct {
		c echo.Context
	}
	type UserResponse struct {
		Message string               `json: "message"`
		Users   []model.UserResponse `json: "users"`
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "success",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := []model.UserResponse{{ID: 1, Name: "John Doe", Email: "johndoe@example.com", Password: "johndoe1234"}, {ID: 2, Name: "Jane Doe", Email: "janedoe@example.com", Password: "janedoe1234"}}
			userRepository.Mock.On("GetUsers").Return(data, nil)

			e := echo.New()

			req := httptest.NewRequest(http.MethodGet, "/users/", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			controller := Controller{}
			controller.GetUsersController(c)

			assert.Equal(t, http.StatusOK, rec.Code)
			var users UserResponse
			err := json.Unmarshal(rec.Body.Bytes(), &users)
			fmt.Println(users)
			assert.NoError(t, err)
			assert.Equal(t, "John Doe", users.Users[0].Name)
		})
	}
}

func TestGetUserController(t *testing.T) {
	userRepository := &service.UserRepositoryMock{Mock: mock.Mock{}}
	service.SetUserRepository(userRepository)

	type args struct {
		c echo.Context
	}

	type UserResponse struct {
		Message string             `json:"message"`
		User    model.UserResponse `json: "users"`
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "success",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := []model.UserResponse{{ID: 1, Name: "John Doe", Email: "johndoe@example.com", Password: "johndoe1234"}, {ID: 2, Name: "Jane Doe", Email: "janedoe@example.com", Password: "janedoe1234"}}
			id := "1"
			userRepository.Mock.On("GetUserById", id).Return(&data, nil)

			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/users/:"+id, nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetParamNames("id")
			c.SetParamValues("1")

			controller := Controller{}
			err := controller.GetUserController(c)

			assert.NoError(t, err)
			assert.Equal(t, http.StatusOK, rec.Code)

			var user UserResponse
			err = json.Unmarshal(rec.Body.Bytes(), &user)
			fmt.Println(user)
			assert.NoError(t, err)
			assert.Equal(t, "John Doe", user.User.Name)
		})
	}
}

func TestCreateUserController(t *testing.T) {
	userRepository := &service.UserRepositoryMock{Mock: mock.Mock{}}
	service.SetUserRepository(userRepository)

	type args struct {
		c echo.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "success",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := model.User{}
			userRepository.Mock.On("CreateUser", &data).Return(nil)

			e := echo.New()

			bData, _ := json.Marshal(data)
			req := httptest.NewRequest(http.MethodPost, "/users/", bytes.NewReader(bData))
			req.Header.Set("content-type", "application/json")
			fmt.Println(string(bData))
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			controller := Controller{}
			controller.CreateUserController(c)

			assert.Equal(t, http.StatusOK, rec.Code)
		})
	}
}

func TestDeleteUserController(t *testing.T) {
	userRepository := &service.UserRepositoryMock{Mock: mock.Mock{}}
	service.SetUserRepository(userRepository)

	type args struct {
		c echo.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "success",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id := "1"
			userRepository.Mock.On("DeleteUser", id).Return(nil)

			e := echo.New()
			req := httptest.NewRequest(http.MethodDelete, "/users/:"+id, nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetParamNames("id")
			c.SetParamValues("1")

			controller := Controller{}
			err := controller.DeleteUserController(c)

			assert.NoError(t, err)
			assert.Equal(t, http.StatusOK, rec.Code)
		})
	}
}

func TestController_UpdateUserController(t *testing.T) {
	userRepository := &service.UserRepositoryMock{Mock: mock.Mock{}}
	service.SetUserRepository(userRepository)

	type args struct {
		c echo.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "success",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user := &model.User{Name: "John Doe", Email: "johndoe@example.com", Password: "johndoe1234"}
			// if err := service.GetUserRepository().CreateUser(user); err != nil {
			// 	t.Errorf("Failed to create user: %v", err)
			// }

			id := strconv.Itoa(int(user.ID))
			userRepository.Mock.On("UpdateUser", user, id, "Updated Name", "updated@example.com", "updatedpassword").Return(nil)

			e := echo.New()
			form := url.Values{}
			form.Set("name", "Updated Name")
			form.Set("email", "updated@example.com")
			form.Set("password", "updatedpassword")
			req := httptest.NewRequest(http.MethodPut, "/users/:"+id, nil)
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetParamNames("id")
			c.SetParamValues(id)

			controller := Controller{}
			err := controller.UpdateUserController(c)

			assert.NoError(t, err)
			assert.Equal(t, http.StatusOK, rec.Code)
		})
	}
}
