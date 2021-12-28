// @Title : main
// @Description ://TODO: Add Description
// @Author : MX
// @Update : 2021/12/22 15:58 

package main

import (
	"classPai/controller"
	"classPai/dao"
	"classPai/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	dao.ConnDB()
	r := gin.Default()
	login := r.Group("")
	login.Use(middleware.LoginRequired)

	r.POST("/user", controller.Register)
	r.GET("/user/login", controller.Login)
	r.GET("/user/:id",controller.ShowUser)
	login.PUT("/user/:id",controller.AlterUser)

	r.GET("/topic/:id", controller.ShowTopic)
	login.POST("/topic", controller.UploadTopic)
	login.PUT("/topic/:id", controller.AlterTopic)
	login.DELETE("/topic/:id", controller.DeleteTopic)
	login.PUT("/topic/like/:id", controller.LikeTopic)

	login.POST("/comment",controller.UploadComment)
	login.DELETE("/comment/:id",controller.DeleteComment)

	r.Run(":8080")
}
