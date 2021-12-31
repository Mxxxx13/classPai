// @Title : file
// @Description ://TODO: Add Description
// @Author : MX
// @Update : 2021/12/30 17:21 

package controller

import (
	"net/http"

	"classPai/resp"
	"classPai/service"
	"github.com/gin-gonic/gin"
)

func UploadFile(c *gin.Context) {
	uploaded, err := service.UploadFile(c)
	if err != nil {
		resp.ErrorResp(c, http.StatusBadRequest, "上传失败", err)
	} else {
		data := struct {
			uploaded int64
		}{uploaded: uploaded}
		resp.SuccessResp(c, "上传成功",data)
	}
}

func DownloadFile(c *gin.Context) {
	file, err := service.DownloadFile(c)
	if err != nil {
		resp.ErrorResp(c, http.StatusBadRequest, "下载失败", err)
	} else {
		c.Header("Content-Type", "application/octet-stream")
		c.Header("Content-Disposition", "attachment; filename="+file.Name)
		c.Header("Content-Transfer-Encoding", "binary")
		c.File(file.Path)
	}
}
