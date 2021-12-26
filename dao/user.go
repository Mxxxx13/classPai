// @Title : user
// @Description ://TODO: Add Description
// @Author : MX
// @Update : 2021/12/22 15:53 

package dao

import "classPai/model"

//Register 将username,password插入数据库
func Register(user model.User) (err error) {
	if err = DB.Create(&user).Error; err != nil {
		return
	}
	return
}

// Login 根据username查询password并返回
func Login(username string) (password string, err error) {
	var user model.User
	if err = DB.Where("username = ?", username).First(&user).Error; err != nil {
		return
	}
	return user.Password, nil
}

// GetUid 根据username查询uid
func GetUid(username string) (uid uint, err error) {
	var user model.User
	if err = DB.Where("username = ?", username).First(&user).Error; err != nil {
		return
	}
	return user.ID, nil
}

// GetUsername 根据uid查询username
func GetUsername(uid uint) (username string, err error) {
	var user model.User
	if err = DB.Select("username").Where("id = ?", uid).First(&user).Error; err != nil {
		return
	}
	return user.Username, err
}

// GetUser
func GetUser(uid uint) (user model.User,err error) {
	if err = DB.Where("id = ?", uid).First(&user).Error; err != nil {
		return
	}
	return
}

// AlterUser
func AlterUser(username string,uid uint) (err error) {
	var user model.User
	if err = DB.Model(&user).Where("id = ?", uid).Updates(model.User{
		Username: username,
	}).Error; err != nil {
		return
	}
	return
}