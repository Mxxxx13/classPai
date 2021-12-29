// @Title : homework
// @Description ://TODO: Add Description
// @Author : MX
// @Update : 2021/12/29 15:41 

package model

import (
	"github.com/jinzhu/gorm"
)

// Homework
type Homework struct {
	*gorm.Model
	Tid      uint   // 教师id
	CrId     uint   // 课堂id
	Name     string // 名字
	Deadline string // 截止时间
	Content  string // 作业内容
}

type HomeworkResp struct {
	Name     string // 名字
	Deadline string // 截止时间
	Content  string // 作业内容
}