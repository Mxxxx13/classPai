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

	login.POST("/classroom",controller.CreateClassroom)
	login.DELETE("/classroom/:id",controller.DeleteClassroom)

	login.GET("/homework/:id",controller.ShowHomework)
	login.POST("/homework", controller.UploadHomework)
	login.PUT("/homework/:id", controller.AlterHomework)
	login.DELETE("/homework/:id", controller.DeleteHomework)

	login.POST("/sign/teacher",controller.UploadSign)
	login.POST("/sign", controller.Sign)
	login.GET("/sign/:id", controller.ShowSign)

	login.GET("/score", controller.ShowScore)
	login.POST("/score", controller.UploadScore)
	login.PUT("/score/:id", controller.AlterScore)
	login.DELETE("/score/:id", controller.DeleteScore)

	r.Run(":8080")
}
