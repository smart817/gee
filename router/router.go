package router

import (
	"gee/router/middleware"
	"gee/util"

	"gee/app/controller"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	r.Use(middleware.Log())
	r.GET("/test", func(c *gin.Context) { c.JSON(200, gin.H{"app_name": util.AppName, "version": util.AppVersion}) })
	r.GET("/index", controller.Index)
	r.GET("/test2", controller.Test)
}
