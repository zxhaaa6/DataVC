package main

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/zxhaaa6/DataVC/util"
)

func main() {
	/* ==== init [config | env | log] ==== */
	util.LoadEnv()
	util.InitLog()
	config := util.InitConfig()

	/* ==== init GIN APP ==== */
	gin.SetMode(util.GetEnv("GIN_MODE"))
	app := gin.New()

	/* ==== middleware ==== */
	if util.GetEnv("GIN_ENV") != "production" {
		app.Use(gin.Logger())
	}
	app.Use(gin.Recovery())

	/* ==== router ==== */

	/* ==== startUp ==== */
	port := config.Port
	log.Infoln("[GIN]Service Listening on port:", port)

	app.Run(":" + port)
}
