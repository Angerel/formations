package repositories

import (
	"github.com/gorilla/websocket"
	"go-api/errors"
	"go-api/models"
)

type UserRepositoryStruct struct {
	Registered []*models.ConnectedUser
}

func (instance *UserRepositoryStruct) FindAll() (*[]*models.ConnectedUser, *errors.ErrorInterface) {
	return &instance.Registered, nil
}

func (instance *UserRepositoryStruct) FindById(userId string) (*models.ConnectedUser, *errors.ErrorInterface) {
	for _, user := range instance.Registered {
		if user.User.Id == userId {
			return user, nil
		}
	}
	return nil, errors.NotFoundException("User not found")
}

func (instance *UserRepositoryStruct) FindByName(userName string) (*models.ConnectedUser, *errors.ErrorInterface) {
	for _, user := range instance.Registered {
		if user.User.Name == userName {
			return user, nil
		}
	}
	return nil, errors.NotFoundException("User not found")
}

func (instance *UserRepositoryStruct) FindByConnection(userConnection *websocket.Conn) (*models.ConnectedUser, *errors.ErrorInterface) {
	for _, user := range instance.Registered {
		if user.Connection == userConnection {
			return user, nil
		}
	}
	return nil, errors.NotFoundException("User not found")
}

func (instance *UserRepositoryStruct) CreateOne(user *models.ConnectedUser) (*models.ConnectedUser, *errors.ErrorInterface) {
	var requested, err = instance.FindById(user.User.Id)
	if requested != nil {
		return nil, errors.ConflictException("A user already exists with this identifier")
	}
	instance.Registered = append(instance.Registered, user)
	requested, err = instance.FindById(user.User.Id)
	if err != nil {
		return nil, errors.InternalException("User creation failed")
	}
	return requested, nil
}

func (instance *UserRepositoryStruct) LinkConnection(userId string, connection *websocket.Conn) (*models.ConnectedUser, *errors.ErrorInterface) {
	var user, err = instance.FindById(userId)
	if err != nil {
		return nil, err
	}
	user.Connection = connection
	return user, nil
}

func (instance *UserRepositoryStruct) UnLinkConnection(userId string) (*models.ConnectedUser, *errors.ErrorInterface) {
	var user, err = instance.FindById(userId)
	if err != nil {
		return nil, err
	}
	user.Connection = nil
	return user, nil
}

// UserRepository - Initializes the repository's data store
var UserRepository = UserRepositoryStruct{
	[]*models.ConnectedUser{},
}
