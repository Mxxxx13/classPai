// @Title : user
// @Description :user model
// @Author : MX
// @Update : 2021/12/22 15:06 

package model

import (
	"github.com/jinzhu/gorm"
)

// User 用户模型
type User struct {
	*gorm.Model
	Username string
	Password string
	Role     int // 0表示学生, 1代表老师
}

// UserResp 用户相关响应
type UserResp struct {
	Uid      uint   `json:"uid"`      // 用户id
	Username string `json:"username"` // 用户名
	Role     string	`json:"role"`
}
