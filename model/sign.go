// @Title : sign
// @Description ://TODO: Add Description
// @Author : MX
// @Update : 2021/12/30 11:56 

package model

import (
	"github.com/jinzhu/gorm"
)

// 签到
type Sign struct {
	*gorm.Model
	Tid     uint // 发布签到老师的id
	Crid    uint // 课堂id
	Name    string
}

// 签到表
type SignTable struct {
	*gorm.Model
	Uid  uint // 用户id
	Sid  uint // 签到id
}