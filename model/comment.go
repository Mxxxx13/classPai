// @Title : comment
// @Description ://TODO: Add Description
// @Author : MX
// @Update : 2021/12/27 16:08 

package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Comment struct {
	*gorm.Model
	Uid     uint   // 评论作者id
	Pid     int    // 父评论id,为0表示为一级评论
	Tid     int    // 话题id
	Content string // 评论内容
}

type CommentResp struct {
	Username      string
	Content       string
	time          time.Time
	ChildComments []CommentResp
}
