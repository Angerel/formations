package services

import (
	"go-api/errors"
	"go-api/models"
)
import "go-api/repositories"

type MessageServiceStruct struct{}

func (instance MessageServiceStruct) FindAll() (*[]*models.Message, *errors.ErrorInterface) {
	return repositories.MessageRepository.FindAll()
}
func (instance MessageServiceStruct) FindById(messageId string) (*models.Message, *errors.ErrorInterface) {
	return repositories.MessageRepository.FindById(messageId)
}

func (instance MessageServiceStruct) CreateOne(message *models.Message) (*models.Message, *errors.ErrorInterface) {
	return repositories.MessageRepository.CreateOne(message)
}

var MessageService = MessageServiceStruct{}
