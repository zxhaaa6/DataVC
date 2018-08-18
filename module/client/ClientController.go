package client

import "github.com/gin-gonic/gin"

type Controller struct {
}

func InitController(route *gin.RouterGroup) {
	controller := Controller{}
	route.POST("/connect", controller.connect)
}

func (r *Controller) connect(c *gin.Context) {
	var connectOption ConnectOptionType
	c.BindJSON(&connectOption)
	err := HandleConnection(connectOption)
	if err != nil {
		c.JSON(500, false)
		return
	}
	c.JSON(200, true)
}
