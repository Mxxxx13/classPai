// @Title : user
// @Description ://TODO: Add Description
// @Author : MX
// @Update : 2021/12/22 15:52 

package service

import (
	"errors"
	"log"
	"strconv"

	"classPai/dao"
	"classPai/model"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) (err error){
	var user model.User
	user.Username = c.PostForm("username")

	password := c.PostForm("password")
	// 将密码加密
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	if err != nil{
		return errors.New("加密失败")
	}
	user.Password = string(hashPassword)

	roleStr := c.PostForm("role")
	user.Role,err = strconv.Atoi(roleStr)
	err = dao.Register(user)
	return
}

// Login 将输入的password加密后和数据库中的password进行比较
// 返回uid和error, uid用于生成token
func Login(c *gin.Context) (uid uint,err error) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	pass,err := dao.Login(username)	// 获取password
	uid,err = dao.GetUid(username)	// 获取id
	if err != nil {
		return
	}
	// 对password进行验证
	err = bcrypt.CompareHashAndPassword([]byte(pass), []byte(password))
	if err != nil {
		return
	}
	return
}

// ShowUser
func ShowUser(c *gin.Context) (UserResp model.UserResp,err error){
	strId := c.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		log.Printf("atoi err:%v\n", err)
		return
	}
	// 查询用户
	user, err := dao.GetUser(uint(id))
	if err != nil {
		return
	}

	// 判断用户的身份
	var role string
	if user.Role == 0 {
		role = "student"
	} else if user.Role == 1{
		role = "teacher"
	}
	UserResp = model.UserResp{
		Uid:      uint(id),
		Username: user.Username,
		Role: role,
	}
	return
}

// AlterUser
func AlterUser(c *gin.Context) (err error){
	username := c.PostForm("username")
	id, exists := c.Get("uid")
	if !exists  {
		return errors.New("id not exist")
	}
	err = dao.AlterUser(username,id.(uint))
	return
}
