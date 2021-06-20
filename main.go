package main

import (
	"gee/router"
	"gee/util"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(util.AppMode)
	app := gin.Default()
	app.GET("/v", func(c *gin.Context) { c.JSON(200, gin.H{"app_name": util.AppName, "version": util.AppVersion}) })
	router.InitRouter(app)
	app.Run(":" + util.AppPort)
}
