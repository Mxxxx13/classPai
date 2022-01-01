// @Title : topic
// @Description ://TODO: Add Description
// @Author : MX
// @Update : 2021/12/27 15:48 

package controller

import (
	"net/http"

	"classPai/resp"
	"classPai/service"
	"github.com/gin-gonic/gin"
)

func UploadTopic(c *gin.Context) {
	err := service.UploadTopic(c)
	if err != nil {
		resp.ErrorResp(c, http.StatusBadRequest, "上传失败", err)
	} else {
		resp.SuccessResp(c, "上传成功")
	}
}

func LikeTopic(c *gin.Context) {
	err := service.LikeTopic(c)
	if err != nil {
		resp.ErrorResp(c, http.StatusBadRequest, "点赞失败", err)
	} else {
		resp.SuccessResp(c, "点赞成功")
	}
}

func ShowTopic(c *gin.Context) {
	topic, err := service.ShowTopic(c)
	if err != nil {
		resp.ErrorResp(c, http.StatusBadRequest, "操作失败", err)
	} else {
		resp.SuccessResp(c, "操作成功", topic)
	}
}

func AlterTopic(c *gin.Context) {
	err := service.AlterTopic(c)
	if err != nil {
		resp.ErrorResp(c, http.StatusBadRequest, "修改失败", err)
	} else {
		resp.SuccessResp(c, "修改成功")
	}
}

func DeleteTopic(c *gin.Context) {
	err := service.DeleteTopic(c)
	if err != nil {
		resp.ErrorResp(c, http.StatusBadRequest, "删除失败", err)
	} else {
		resp.SuccessResp(c, "删除成功")
	}
}
