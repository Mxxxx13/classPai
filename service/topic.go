// @Title : topic
// @Description ://TODO: Add Description
// @Author : MX
// @Update : 2021/12/27 15:48 

package service

import (
	"errors"
	"log"
	"strconv"

	"classPai/dao"
	"classPai/model"
	"github.com/gin-gonic/gin"
)

func UploadTopic(c *gin.Context) (err error) {
	var topic model.Topic
	topic.Title = c.PostForm("title")
	topic.Content = c.PostForm("content")

	uid, exists := c.Get("uid")
	if !exists {
		return errors.New("uid is not exist")
	}

	user ,err := dao.GetUser(uid.(uint))
	if err != nil {
		return err
	}

	topic.Author = user.Username
	err = dao.UploadTopic(topic)
	return err
}

func LikeTopic(c *gin.Context) (err error) {
	strId := c.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		log.Printf("atoi err:%v\n", err)
		return
	}
	err = dao.LikeTopic(id)
	return err
}

func ShowTopic(c *gin.Context) (TopicResp model.TopicResp, err error) {
	strId := c.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		log.Printf("atoi err:%v\n", err)
		return
	}
	topic, err := dao.ShowTopic(id)
	if err != nil {
		return
	}
	TopicResp = model.TopicResp{
		Title:       topic.Title,
		Content:     topic.Content,
		Author:      topic.Author,
		Likes:       topic.Likes,
		CommentResp: SplicingComment(int(topic.ID)),
	}
	return
}

func AlterTopic(c *gin.Context) (err error) {
	strId := c.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		log.Printf("atoi err:%v\n", err)
		return
	}
	var topic model.Topic
	topic.Title = c.PostForm("title")
	topic.Content = c.PostForm("content")
	err = dao.AlterTopic(id, topic)
	return err
}

func DeleteTopic(c *gin.Context) (err error) {
	strId := c.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		log.Printf("atoi err:%v\n", err)
		return
	}
	err = dao.DeleteTopic(id)
	return err
}
