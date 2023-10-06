package repositories

import (
	"go-api/errors"
	"go-api/models"
)

type TodoRepositoryStruct struct {
	Registered []*models.TodoList
}

func (instance *TodoRepositoryStruct) FindAll() (*[]*models.TodoList, *errors.ErrorInterface) {
	return &instance.Registered, nil
}

func (instance *TodoRepositoryStruct) FindListById(listId string) (*models.TodoList, *errors.ErrorInterface) {
	for _, list := range instance.Registered {
		if list.Id == listId {
			return list, nil
		}
	}
	return nil, errors.NotFoundException("List not found")
}

func (instance *TodoRepositoryStruct) FindItemByListAndId(listId string, itemId string) (*models.TodoItem, *errors.ErrorInterface) {
	list, err := instance.FindListById(listId)
	if err != nil {
		return nil, err
	}
	for _, item := range list.List {
		if item.Id == itemId {
			return item, nil
		}
	}
	return nil, errors.NotFoundException("Item not found")
}

func (instance *TodoRepositoryStruct) CreateList(list *models.TodoList) (*models.TodoList, *errors.ErrorInterface) {
	var requested, err = instance.FindListById(list.Id)
	if requested != nil {
		return nil, errors.ConflictException("A list already exists with this identifier")
	}
	instance.Registered = append(instance.Registered, list)
	requested, err = instance.FindListById(list.Id)
	if err != nil {
		return nil, errors.InternalException("List creation failed")
	}
	return requested, nil
}

func (instance *TodoRepositoryStruct) CreateItem(listId string, item *models.TodoItem) (*models.TodoItem, *errors.ErrorInterface) {
	var requestedList, err = instance.FindListById(listId)
	if requestedList == nil {
		return nil, errors.NotFoundException("List not found")
	}
	var requestedItem, _ = instance.FindItemByListAndId(listId, item.Id)
	if requestedItem != nil {
		return nil, errors.ConflictException("An item already exists with this identifier in this list")
	}
	requestedList.List = append(requestedList.List, item)
	requestedItem, err = instance.FindItemByListAndId(listId, item.Id)
	if err != nil {
		return nil, errors.InternalException("Item creation failed")
	}
	return requestedItem, nil
}

var TodoRepository = TodoRepositoryStruct{
	[]*models.TodoList{},
}
