// @Title : user
// @Description ://TODO: Add Description
// @Author : MX
// @Update : 2021/12/22 15:52 

package controller

import (
	"net/http"

	"classPai/resp"
	"classPai/service"
	"classPai/util"
	"github.com/gin-gonic/gin"
)

// Register 返回注册接口
func Register(c *gin.Context) {
	err := service.Register(c)
	if err != nil {
		msg := "用户已存在,注册失败"
		resp.ErrorResp(c, http.StatusBadRequest, msg)
	} else {
		resp.SuccessResp(c, "注册成功")
	}
}

//Login 返回登录接口,成功登录返回token
func Login(c *gin.Context) {
	id, err := service.Login(c)
	username := c.PostForm("username")
	if err != nil {
		msg := "用户名或密码错误"
		resp.ErrorResp(c, http.StatusBadRequest, msg)

	} else {
		jwt := util.NewJWT(id, username)
		msg := "欢迎回来" + username

		c.JSON(http.StatusOK,gin.H{
			"code":http.StatusOK,
			"message":msg,
			"token":jwt.Token,
		})
		//data := struct {
		//	token string
		//}{token: jwt.Token}
		//resp.SuccessResp(c, msg, data)
	}
}

// ShowUser
func ShowUser(c *gin.Context)  {
	userResp,err := service.ShowUser(c)
	if err != nil {
		resp.ErrorResp(c, http.StatusBadRequest, "操作失败")
	} else {
		resp.SuccessResp(c, "操作成功",userResp)
	}
}

// AlterUser
func AlterUser(c *gin.Context)  {
	err := service.AlterUser(c)
	if err != nil {
		resp.ErrorResp(c, http.StatusBadRequest, "修改失败")
	} else {
		resp.SuccessResp(c, "修改成功")
	}
}