// @Title : classroom
// @Description ://TODO: Add Description
// @Author : MX
// @Update : 2021/12/28 16:36 

package model

import "github.com/jinzhu/gorm"

// Classroom 课堂
type Classroom struct {
	*gorm.Model
	Name    string // 名字
	Subject string // 科目
	Tid     uint // 教师id
}
