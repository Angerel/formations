package services

import (
	"github.com/gorilla/websocket"
	"go-api/errors"
	"go-api/models"
)
import "go-api/repositories"

type UserServiceStruct struct{}

func (instance UserServiceStruct) FindAll() (*[]*models.ConnectedUser, *errors.ErrorInterface) {
	return repositories.UserRepository.FindAll()
}

func (instance UserServiceStruct) FindById(userId string) (*models.ConnectedUser, *errors.ErrorInterface) {
	return repositories.UserRepository.FindById(userId)
}

func (instance UserServiceStruct) FindByName(userName string) (*models.ConnectedUser, *errors.ErrorInterface) {
	return repositories.UserRepository.FindByName(userName)
}

func (instance UserServiceStruct) FindByConnection(connection *websocket.Conn) (*models.ConnectedUser, *errors.ErrorInterface) {
	return repositories.UserRepository.FindByConnection(connection)
}

func (instance UserServiceStruct) CreateOne(user *models.ConnectedUser) (*models.ConnectedUser, *errors.ErrorInterface) {
	return repositories.UserRepository.CreateOne(user)
}

func (instance UserServiceStruct) LinkConnection(userId string, connection *websocket.Conn) (*models.ConnectedUser, *errors.ErrorInterface) {
	return repositories.UserRepository.LinkConnection(userId, connection)
}

func (instance UserServiceStruct) UnLinkConnection(userId string) (*models.ConnectedUser, *errors.ErrorInterface) {
	return repositories.UserRepository.UnLinkConnection(userId)
}

var UserService = UserServiceStruct{}
