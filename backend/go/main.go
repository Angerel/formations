package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	REST "go-api/controllers/rest"
	"strconv"
)

func main() {
	// Default allowed origins and proxies
	allowedOrigins := []string{"http://localhost:4200"}
	// Port of execution
	port := 8080

	/* Initializing Gin router and proxy policy */
	router := gin.Default()
	router.ForwardedByClientIP = true
	err := router.SetTrustedProxies(allowedOrigins)
	if err != nil {
		return
	}

	/* CORS config */
	config := cors.DefaultConfig()
	config.AllowOrigins = allowedOrigins
	router.Use(cors.New(config))

	/* REST API endpoints */
	router.GET("/", REST.PingController.Ping)

	err = router.Run("localhost:" + strconv.Itoa(port))
	if err != nil {
		return
	}
}
