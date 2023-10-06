package REST

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type PingControllerStruct struct{}

func (instance PingControllerStruct) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, "Pong !")
}

var PingController = PingControllerStruct{}
