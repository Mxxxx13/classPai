// @Title : score
// @Description ://TODO: Add Description
// @Author : MX
// @Update : 2021/12/30 17:02 

package controller

import (
	"net/http"

	"classPai/resp"
	"classPai/service"
	"github.com/gin-gonic/gin"
)

func UploadScore(c *gin.Context){
	err := service.UploadScore(c)
	if err != nil {
		resp.ErrorResp(c,http.StatusBadRequest,"上传失败")
	} else {
		resp.SuccessResp(c,"上传成功")
	}
}

func ShowScore(c *gin.Context){
	score, err := service.ShowScore(c)
	if err != nil {
		resp.ErrorResp(c,http.StatusBadRequest,"操作失败")
	} else {
		resp.SuccessResp(c,"操作成功",score)
	}
}

func AlterScore(c *gin.Context){
	err := service.AlterScore(c)
	if err != nil {
		resp.ErrorResp(c,http.StatusBadRequest,"修改失败")
	} else {
		resp.SuccessResp(c,"修改成功")
	}
}

func DeleteScore(c *gin.Context){
	err := service.DeleteScore(c)
	if err != nil {
		resp.ErrorResp(c,http.StatusBadRequest,"删除失败")
	} else {
		resp.SuccessResp(c,"删除成功")
	}
}
