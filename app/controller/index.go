package controller

import (
	"gee/app/controller/context"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.JSON(200, gin.H{"code": 200, "message": "msg", "data": nil})
}

func Test(c *context.Context) {
	//c.JSON(200, gin.H{"code": 200, "message": "msg", "data": nil})
	c.Success("sdfsdfdsf")
}
