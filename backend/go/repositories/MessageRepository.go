package repositories

import (
	"go-api/errors"
	"go-api/models"
)

type MessageRepositoryStruct struct {
	Registered []*models.Message
}

func (instance *MessageRepositoryStruct) FindAll() (*[]*models.Message, *errors.ErrorInterface) {
	return &instance.Registered, nil
}

func (instance *MessageRepositoryStruct) FindById(messageId string) (*models.Message, *errors.ErrorInterface) {
	for _, message := range instance.Registered {
		if message.Id == messageId {
			return message, nil
		}
	}
	return nil, errors.NotFoundException("Message not found")
}

func (instance *MessageRepositoryStruct) CreateOne(message *models.Message) (*models.Message, *errors.ErrorInterface) {
	var requested, err = instance.FindById(message.Id)
	if requested != nil {
		return nil, errors.ConflictException("A message already exists with this identifier")
	}
	instance.Registered = append(instance.Registered, message)
	requested, err = instance.FindById(message.Id)
	if err != nil {
		return nil, errors.InternalException("Message creation failed")
	}
	return requested, nil
}

var MessageRepository = MessageRepositoryStruct{
	[]*models.Message{},
}
