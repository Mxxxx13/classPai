// @Title : classroom
// @Description ://TODO: Add Description
// @Author : MX
// @Update : 2021/12/29 14:23 

package controller

import (
	"net/http"

	"classPai/resp"
	"classPai/service"
	"github.com/gin-gonic/gin"
)

// CreateClassroom
func CreateClassroom(c *gin.Context) {
	err := service.CreateClassroom(c)
	if err != nil {
		resp.ErrorResp(c, http.StatusBadRequest, "创建失败", err)
	} else {
		resp.SuccessResp(c, "创建成功")
	}
}

// DeleteClassroom
func DeleteClassroom(c *gin.Context) {
	err := service.DeleteClassroom(c)
	if err != nil {
		resp.ErrorResp(c, http.StatusBadRequest, "删除失败", err)
	} else {
		resp.SuccessResp(c, "删除成功")
	}
}
