// @Title : homework
// @Description ://TODO: Add Description
// @Author : MX
// @Update : 2021/12/29 15:52 

package service

import (
	"errors"
	"log"
	"strconv"

	"classPai/dao"
	"classPai/model"
	"github.com/gin-gonic/gin"
)

func UploadHomework(c *gin.Context) (err error) {
	var homework model.Homework
	homework.Name = c.PostForm("name")
	homework.Content = c.PostForm("content")
	homework.Deadline = c.PostForm("deadline")

	CrIdStr := c.PostForm("crid")
	CrId, err := strconv.Atoi(CrIdStr)
	homework.CrId = uint(CrId)
	if err != nil {
		return
	}

	uid, exists := c.Get("uid")
	if !exists {
		return errors.New("uid is not exist")
	}
	user, err := dao.GetUser(uid.(uint))
	if err != nil {
		return
	}
	// 判断用户身份是否为教师
	if user.Role != 1 {
		return errors.New("权限不够")
	}
	homework.Tid = user.ID

	err = dao.UploadHomework(homework)
	return err
}

func ShowHomework(c *gin.Context) (resp model.HomeworkResp, err error) {
	strId := c.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		log.Printf("atoi err:%v\n", err)
		return
	}

	homework, err := dao.GetHomework(id)
	if err != nil {
		return
	}

	resp = model.HomeworkResp{
		Name:     homework.Name,
		Deadline: homework.Deadline,
		Content:  homework.Content,
	}
	return
}

func AlterHomework(c *gin.Context) (err error) {
	strId := c.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		log.Printf("atoi err:%v\n", err)
		return
	}
	var homework model.Homework
	homework.Name = c.PostForm("Name")
	homework.Content = c.PostForm("content")
	homework.Deadline = c.PostForm("deadline")
	err = dao.AlterHomework(id, homework)
	return err
}

func DeleteHomework(c *gin.Context) (err error) {
	strId := c.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		log.Printf("atoi err:%v\n", err)
		return
	}

	uid, exists := c.Get("uid")
	if !exists {
		return errors.New("uid is not exist")
	}

	homework,err := dao.GetHomework(uid.(int))
	if err != nil {
		return
	}
	// 判断是否为本人操作
	// 只能删除由自己布置的作业
	if uid != homework.Tid{
		return errors.New("你不能删除其他老师布置的作业")
	}
	err = dao.DeleteHomework(id)
	return err
}
