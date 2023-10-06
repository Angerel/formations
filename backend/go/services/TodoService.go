package services

import (
	"go-api/errors"
	"go-api/models"
)
import "go-api/repositories"

type TodoServiceStruct struct{}

func (instance TodoServiceStruct) FindAll() (*[]*models.TodoList, *errors.ErrorInterface) {
	return repositories.TodoRepository.FindAll()
}

func (instance TodoServiceStruct) FindListById(listId string) (*models.TodoList, *errors.ErrorInterface) {
	return repositories.TodoRepository.FindListById(listId)
}

func (instance TodoServiceStruct) FindItemByListAndId(listId string, itemId string) (*models.TodoItem, *errors.ErrorInterface) {
	return repositories.TodoRepository.FindItemByListAndId(listId, itemId)
}

func (instance TodoServiceStruct) CreateList(list *models.TodoList) (*models.TodoList, *errors.ErrorInterface) {
	return repositories.TodoRepository.CreateList(list)
}

func (instance TodoServiceStruct) CreateItem(listId string, item *models.TodoItem) (*models.TodoItem, *errors.ErrorInterface) {
	return repositories.TodoRepository.CreateItem(listId, item)
}

var TodoService = TodoServiceStruct{}
