package client

import (
	"github.com/gin-gonic/gin"
	"github.com/zxhaaa6/DataVC/util"
	"github.com/zxhaaa6/DataVC/module/common_type"
)

type Controller struct {
}

func InitController(route *gin.RouterGroup) {
	controller := Controller{}
	route.POST("/connect", controller.connect)
	route.GET("/x", controller.ping)
}

func (r *Controller) connect(c *gin.Context) {
	var connectOption common_type.ConnectOptionType
	c.BindJSON(&connectOption)
	err := HandleConnection(connectOption)
	if err != nil {
		c.JSON(500, false)
		return
	}
	c.JSON(200, true)
}

func (r *Controller) ping(c *gin.Context) {
	c.JSON(200, util.GenSimpleJson("aaa"))
}
