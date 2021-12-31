// @Title : homework
// @Description ://TODO: Add Description
// @Author : MX
// @Update : 2021/12/29 15:52 

package controller

import (
	"net/http"

	"classPai/resp"
	"classPai/service"
	"github.com/gin-gonic/gin"
)

func UploadHomework(c *gin.Context) {
	err := service.UploadHomework(c)
	if err != nil {
		resp.ErrorResp(c, http.StatusBadRequest, "上传失败", err)
	} else {
		resp.SuccessResp(c, "上传成功")
	}
}

func ShowHomework(c *gin.Context) {
	blog, err := service.ShowHomework(c)
	if err != nil {
		resp.ErrorResp(c, http.StatusBadRequest, "操作失败", err)
	} else {
		resp.SuccessResp(c, "操作成功", blog)
	}
}

func AlterHomework(c *gin.Context) {
	err := service.AlterHomework(c)
	if err != nil {
		resp.ErrorResp(c, http.StatusBadRequest, "修改失败", err)
	} else {
		resp.SuccessResp(c, "修改成功")
	}
}

func DeleteHomework(c *gin.Context) {
	err := service.DeleteHomework(c)
	if err != nil {
		resp.ErrorResp(c, http.StatusBadRequest, "删除失败", err)
	} else {
		resp.SuccessResp(c, "删除成功")
	}
}
