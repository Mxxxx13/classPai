// @Title : model
// @Description ://TODO: Add Description
// @Author : MX
// @Update : 2021/12/30 16:18 

package model

import (
	"github.com/jinzhu/gorm"
)

type Score struct {
	*gorm.Model
	Uid      uint
	Username string
	Score    float64
	Subject  string
}

type ScoreResp struct {
	Username string  `json:"username"`
	Score    float64 `json:"score"`
	Subject  string  `json:"subject"`
}
