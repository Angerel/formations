package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	REST "go-api/controllers/rest"
	WS "go-api/controllers/websocket"
	"strconv"
)

func main() {
	// Default allowed origins and proxies
	allowedProxies := []string{"127.0.0.1"}
	allowedOrigins := []string{"http://127.0.0.1:5173", "http://localhost:5173"}
	// Port of execution
	port := 8080

	/* Initializing Gin router and proxy policy */
	router := gin.Default()
	router.ForwardedByClientIP = true
	err := router.SetTrustedProxies(allowedProxies)
	if err != nil {
		fmt.Printf("An unexpected error occurred : %+v\n", err)
		return
	}

	/* CORS config */
	config := cors.DefaultConfig()
	config.AllowOrigins = allowedOrigins
	router.Use(cors.New(config))

	/* REST API endpoints */
	router.GET("/", REST.PingController.Ping)
	router.GET("/todos", REST.TodoController.GetAllTodoList)
	router.GET("/todos/:list-id", REST.TodoController.GetOneTodoList)
	router.GET("/todos/:list-id/:item-id", REST.TodoController.GetOneTodoItem)
	router.POST("/todos", REST.TodoController.AddOneTodoList)
	router.POST("/todos/:list-id", REST.TodoController.AddOneTodoItem)

	/* WS Endpoint */
	router.GET("/ws/live-chat", WS.ChatController.JoinChat)

	err = router.Run("localhost:" + strconv.Itoa(port))
	if err != nil {
		fmt.Printf("An unexpected error occurred : %+v\n", err)
		return
	}
}
