// @Title : sign
// @Description ://TODO: Add Description
// @Author : MX
// @Update : 2021/12/30 12:19 

package controller

import (
	"net/http"

	"classPai/resp"
	"classPai/service"
	"github.com/gin-gonic/gin"
)

func UploadSign(c *gin.Context) {
	err := service.UploadSign(c)
	if err != nil {
		resp.ErrorResp(c, http.StatusBadRequest, "上传失败", err)
	} else {
		resp.SuccessResp(c, "上传成功")
	}
}

func Sign(c *gin.Context) {
	err := service.Sign(c)
	if err != nil {
		resp.ErrorResp(c, http.StatusBadRequest, "签到失败", err)
	} else {
		resp.SuccessResp(c, "签到成功")
	}
}

func ShowSign(c *gin.Context) {
	usernames, err := service.ShowSign(c)
	if err != nil {
		resp.ErrorResp(c, http.StatusBadRequest, "操作失败", err)
	} else {
		resp.SuccessResp(c, "操作成功", "签到人员名单如下:", usernames)
	}
}
