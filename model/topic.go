// @Title : topic
// @Description ://TODO: Add Description
// @Author : MX
// @Update : 2021/12/27 15:48 

package model

import "github.com/jinzhu/gorm"

// topic 话题模型
type Topic struct {
	*gorm.Model
	Title   string // 标题
	Content string // 内容
	Uid     uint   // 发布者id
	Author  string // 发布者
	Likes   int    // 获赞
}

type TopicResp struct {
	Title       string        `json:"title"`   // 标题
	Content     string        `json:"content"` // 内容
	Author      string        `json:"author"`  // 发布者
	Likes       int           `json:"likes"`   // 获赞
	CommentResp []CommentResp `json:"comments"`
}

type TopicList struct {
	Tid   uint   `json:"tid"`   // 话题id
	Title string `json:"title"` // 标题
}
