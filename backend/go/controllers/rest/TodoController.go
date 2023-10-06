package REST

import (
	"github.com/gin-gonic/gin"
	"go-api/models"
	"go-api/services"
	"net/http"
)

type TodoControllerStruct struct{}

func (instance TodoControllerStruct) GetAllTodoList(c *gin.Context) {
	allLists, err := services.TodoService.FindAll()
	if err != nil {
		c.JSON(err.StatusCode, err.Message)
		return
	}
	c.JSON(http.StatusOK, allLists)
}

func (instance TodoControllerStruct) GetOneTodoList(c *gin.Context) {
	listId := c.Param("list-id")
	list, err := services.TodoService.FindListById(listId)
	if err != nil {
		c.JSON(err.StatusCode, err.Message)
		return
	}
	c.JSON(http.StatusOK, list)
}

func (instance TodoControllerStruct) GetOneTodoItem(c *gin.Context) {
	listId := c.Param("list-id")
	itemId := c.Param("item-id")

	item, err := services.TodoService.FindItemByListAndId(listId, itemId)
	if err != nil {
		c.JSON(err.StatusCode, err.Message)
		return
	}
	c.JSON(http.StatusOK, item)
}

func (instance TodoControllerStruct) AddOneTodoList(c *gin.Context) {
	var newList models.TodoList
	if err := c.BindJSON(&newList); err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	item, err := services.TodoService.CreateList(&newList)
	if err != nil {
		c.JSON(err.StatusCode, err.Message)
		return
	}
	c.JSON(http.StatusOK, item)
}

func (instance TodoControllerStruct) AddOneTodoItem(c *gin.Context) {
	listId := c.Param("list-id")
	var newItem models.TodoItem
	if err := c.BindJSON(&newItem); err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	item, err := services.TodoService.CreateItem(listId, &newItem)
	if err != nil {
		c.JSON(err.StatusCode, err.Message)
		return
	}
	c.JSON(http.StatusOK, item)
}

var TodoController = TodoControllerStruct{}
