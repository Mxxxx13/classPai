// @Title : auth
// @Description ://TODO: Add Description
// @Author : MX
// @Update : 2021/12/22 16:01 

package middleware

import (
	"net/http"

	"classPai/resp"
	"classPai/util"
	"github.com/gin-gonic/gin"
)

func LoginRequired(c *gin.Context) {
	token := c.PostForm("token")
	//根据token解析出jwt
	jwt, err := util.CheckJWT(token)
	if err != nil {
		resp.ErrorResp(c,http.StatusUnauthorized,"需要登录")
		c.Abort()
		return
	}

	// 设置id和username方便后续操作
	c.Set("uid", jwt.Payload.Sub.Uid)
	c.Set("username", jwt.Payload.Sub.Username)

	c.Next()
}