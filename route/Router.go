package route

import (
	"github.com/gin-gonic/gin"
	"github.com/zxhaaa6/DataVC/module/client"
)

func InitRouter(router *gin.Engine) {
	client.InitController(router.Group("client"))
}
