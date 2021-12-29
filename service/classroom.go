// @Title : classroom
// @Description ://TODO: Add Description
// @Author : MX
// @Update : 2021/12/29 14:24 

package service

import (
	"errors"
	"log"
	"strconv"

	"classPai/dao"
	"classPai/model"
	"github.com/gin-gonic/gin"
)

func CreateClassroom(c *gin.Context) (err error) {
	var classroom model.Classroom
	classroom.Name = c.PostForm("title")
	classroom.Subject = c.PostForm("subject")

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
	classroom.Tid = user.ID

	err = dao.CreateClassroom(classroom)
	return err
}


func DeleteClassroom(c *gin.Context) (err error) {
	strId := c.Param("id")
	id, err := strconv.Atoi(strId)

	uid,exists := c.Get("uid")
	if !exists{
		return errors.New("uid is not exist")
	}

	if uid.(int) != id {
		return errors.New("你没有权限删除别人的创建的课堂")
	}

	if err != nil {
		log.Printf("atoi err:%v\n", err)
		return
	}
	err = dao.DeleteClassroom(uint(id))
	return err
}