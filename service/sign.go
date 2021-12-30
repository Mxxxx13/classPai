// @Title : sign
// @Description ://TODO: Add Description
// @Author : MX
// @Update : 2021/12/30 12:19 

package service

import (
	"errors"
	"strconv"

	"classPai/dao"
	"classPai/model"
	"github.com/gin-gonic/gin"
)

func UploadSign(c *gin.Context) (err error) {
	var sign model.Sign
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
	sign.Tid = user.ID

	sign.Name = c.PostForm("name")
	cridStr := c.PostForm("crid")
	crid, err := strconv.Atoi(cridStr)
	if err != nil {
		return err
	}
	sign.Crid = uint(crid)

	err = dao.UploadSign(sign)
	return err
}

func Sign(c *gin.Context) (err error) {
	var sign model.SignTable

	uid, exists := c.Get("uid")
	if !exists {
		return errors.New("uid is not exist")
	}
	sign.Uid = uid.(uint)

	sidStr := c.PostForm("sid")
	sid,err := strconv.Atoi(sidStr)
	if err != nil {
		return
	}
	sign.Sid = uint(sid)

	err = dao.Sign(sign)
	return err
}

func ShowSign(c *gin.Context) (usernames []string, err error) {
	strId := c.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		return
	}

	signs ,err := dao.ShowSign(uint(id))
	if err != nil {
		return
	}

	for _, sign := range signs {
		user, err := dao.GetUser(sign.Uid)
		if err != nil {
			return
		}
		usernames = append(usernames, user.Username)
	}
	return
}
