package service

import (
	"errors"

	"github.com/fazaalexander/go-gorm/model"

	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	Mock mock.Mock
}

func (um *UserRepositoryMock) GetUsers() ([]model.UserResponse, error) {
	args := um.Mock.Called()
	return args.Get(0).([]model.UserResponse), args.Error(1)
}

func (um *UserRepositoryMock) GetUserById(id string) (*model.UserResponse, error) {
	args := um.Mock.Called(id)
	users := args.Get(0).(*[]model.UserResponse)
	if len(*users) == 0 {
		return nil, args.Error(1)
	}
	return &(*users)[0], args.Error(1)
}

func (um *UserRepositoryMock) CreateUser(user *model.User) error {
	args := um.Mock.Called(user)
	if args.Get(0) == nil {
		return errors.New("error")
	} else {
		return nil
	}
}

func (um *UserRepositoryMock) DeleteUser(id string) error {
	args := um.Mock.Called(id)
	return args.Error(0)
}

func (um *UserRepositoryMock) UpdateUser(user *model.User, id string, name string, email string, password string) error {
	args := um.Mock.Called(user, id, name, email, password)
	return args.Error(0)
}
