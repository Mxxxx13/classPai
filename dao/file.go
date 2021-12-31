// @Title : file
// @Description ://TODO: Add Description
// @Author : MX
// @Update : 2021/12/30 17:21 

package dao

import "classPai/model"

func UploadFile(file model.File) (err error) {
	if err = DB.Create(&file).Error; err != nil {
		return err
	}
	return
}

func DownloadFile(path string) (file model.File,err error) {
	if err = DB.Where("path = ?", path).Last(&file).Error; err != nil {
		return
	}
	return
}
