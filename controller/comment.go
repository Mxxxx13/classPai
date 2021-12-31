// @Title : comment
// @Description ://TODO: Add Description
// @Author : MX
// @Update : 2021/12/27 16:08 

package controller

import (
	"net/http"

	"classPai/resp"
	"classPai/service"
	"github.com/gin-gonic/gin"
)

// UploadComment
func UploadComment(c *gin.Context) {
	err := service.UploadComment(c)
	if err != nil {
		resp.ErrorResp(c, http.StatusBadRequest, "评论失败", err)
	} else {
		resp.SuccessResp(c, "评论成功")
	}
}

// DeleteComment
func DeleteComment(c *gin.Context) {
	err := service.DeleteComment(c)
	if err != nil {
		resp.ErrorResp(c, http.StatusBadRequest, "删除失败", err)
	} else {
		resp.SuccessResp(c, "删除成功")
	}
}
