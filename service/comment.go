// @Title : comment
// @Description ://TODO: Add Description
// @Author : MX
// @Update : 2021/12/27 16:08 

package service

import (
	"errors"
	"log"
	"strconv"

	"classPai/dao"
	"classPai/model"
	"github.com/gin-gonic/gin"
)

// UploadComment
func UploadComment(c *gin.Context) (err error) {
	content := c.PostForm("content")
	bid := c.PostForm("bid")
	uid, _ := c.Get("uid")
	pid := c.PostForm("pid")

	bId, err := strconv.Atoi(bid)
	if err != nil {
		return
	}
	var pId int
	if pid != "" {
		pId, err = strconv.Atoi(pid)
		if err != nil {
			return
		}
	}
	comment := model.Comment{
		Uid:     uid.(uint),
		Pid:     pId,
		Bid:     bId,
		Content: content,
	}
	err = dao.UploadComment(comment)
	return
}

// SplicingComment 拼接评论
func SplicingComment(bid int) (CommentResps []model.CommentResp) {
	// 查询id为bid的blog下的所有评论
	comments, err := dao.GetComments(bid)
	if err != nil {
		return
	}

	CommentResps = GetChildComments(comments,0)
	return
}

// GetChildComments 获取子评论
func GetChildComments(comments []model.Comment,id int) (CommentResps []model.CommentResp){
	if len(comments) == 0 {
		return
	}
	for i, comment := range comments {
		if comment.Pid == id {
			// 复制评论切片防止删除评论后对原切片产生影响
			// 费空间省时间
			commentsCopy := make([]model.Comment,len(comments))
			copy(commentsCopy,comments)
			// 从comments中拿取评论后删除,降低for循环次数
			commentsCopy = append(commentsCopy[0:i],commentsCopy[i + 1:]...)
			// 递归得到子评论
			ChildComments := GetChildComments(commentsCopy, int(comment.ID))
			// 构造当前评论
			username, err := dao.GetUsername(comment.Uid)
			if err != nil {
				return
			}
			commentResp := model.CommentResp{
				Username:      username,
				Content:       comment.Content,
				ChildComments: ChildComments,
			}
			// 将当前的评论加入切片
			CommentResps = append(CommentResps,commentResp)
		}
	}
	return
}

func DeleteComment(c *gin.Context) (err error) {
	strId := c.Param("id")
	id, err := strconv.Atoi(strId)

	uid,exists := c.Get("uid")
	if !exists{
		return errors.New("uid is not exist")
	}

	if uid.(int) != id {
		return errors.New("你没有权限删除别人的评论")
	}

	if err != nil {
		log.Printf("atoi err:%v\n", err)
		return
	}
	err = dao.DeleteComment(uint(id))
	return err
}