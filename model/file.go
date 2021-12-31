// @Title : file
// @Description ://TODO: Add Description
// @Author : MX
// @Update : 2021/12/30 17:21 

package model

import "github.com/jinzhu/gorm"

type File struct {
	*gorm.Model
	Uid  uint   // 上传者id
	Name string // 文件名
	Path string // 文件路径
}
