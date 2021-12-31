// @Title : file
// @Description ://TODO: Add Description
// @Author : MX
// @Update : 2021/12/30 17:21 

package service

import (
	"errors"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"

	"classPai/dao"
	"classPai/model"
	"github.com/gin-gonic/gin"
)

// UploadFile 文件上传
// 实现断点续传
func UploadFile(c *gin.Context) (uploaded int64,err error) {
	var file model.File
	uid, exists := c.Get("uid")
	if !exists {
		return 0,errors.New("uid is not exist")
	}
	file.Uid = uid.(uint)

	srcFile, header, err := c.Request.FormFile("file")
	if err != nil {
		return
	}

	// 拼接路径
	uidStr := strconv.FormatUint(uint64(file.Uid),10)
	path := "./save/" + uidStr + "/"

	//递归创建文件夹
	err = os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return
	}

	fileUploadedStr := c.PostForm("fileUploaded")
	fileUploaded,err := strconv.ParseInt(fileUploadedStr,10,64)
 	if err != nil {
 		return
	}

	file.Name = filepath.Base(header.Filename)
	file.Path = path + file.Name
	err = c.SaveUploadedFile(header, file.Path)

	uploaded, err = FileSave(srcFile, file.Path, fileUploaded)
	err = dao.UploadFile(file)
	return
}

// FileSave 保存文件
// file 文件 dstPath 目标路径 uploaded 已上传的字节数
func FileSave(file multipart.File, dstPath string, uploaded int64) (int64, error) {

	// 本地创建文件,如果存在直接打开,不存在则创建
	saveFile, err := os.OpenFile(dstPath, os.O_RDWR, os.ModePerm)

	if err != nil {
		saveFile, _ = os.Create(dstPath)
	}
	defer saveFile.Close()

	// 创建缓冲区
	buf := make([]byte, 2048)

	// 读取之前的上传进度
	_, err = file.Seek(uploaded, io.SeekStart)
	if err != nil {
		return 0, err
	}

	for {
		read, err := file.Read(buf)
		if err == io.EOF {
			break
		}
		// 跳过uploaded字节追加写
		_, err = saveFile.WriteAt(buf, uploaded)
		if err != nil {
			return 0, err
		}
		// 更新已上传的字节数
		uploaded += int64(read)
	}

	return uploaded, err
}

// DownloadFile
func DownloadFile(c *gin.Context) (file model.File, err error) {
	path := c.PostForm("path")
	file, err = dao.DownloadFile(path)
	if err != nil {
		return
	}
	return
}

// FileExist 判断文件是否已经存在
func FileExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
